package website

import (
	"encoding/json"
	"fmt"
	"leiserv/models/common/response"
	models "leiserv/models/website/lianlianpay"
	websiteReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"leiserv/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type LLPayAPI struct{}

var lianlianpay = models.LLPayConfigType{
	BaseUrl:       "https://celer-api.LianLianpay-inc.com/v3",
	LLPubKey:      "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1LUX1mFiLT7XgcAbwC8RxRl8S4o/ihxsTl8d6X1AYJxjeT9kq2I1gmGSJPVdRqQWZZmZ86e4EwZlwOAmuPaXT7ILSbepkSU2IJmv8+Pmx8lG0KybsEgezQ7la+LAllXvKtUp/AA1c3EEZGwjVoWEWvO9PUnyRVYgvJ1seM1AQDtQYvUHfm8a0CNTCYhqNiIivDtnxG4a8H7tcnnwPOipRCffi+S4CnQxXPaWoRbEISjZnt1KpfbkGAFXnbvJPAzDdD",
	PublickKey:    "",
	PrivateKey:    "",
	MerchantID:    "202410250002787001",
	SubMerchantId: "1020241025696001",
}

func (p *LLPayAPI) GetToken(c *gin.Context) {
	url := lianlianpay.BaseUrl + fmt.Sprintf("/merchants/%s/token", lianlianpay.MerchantID)
	res, err := lianlianpayService.GetLLPayToken(url, lianlianpay.MerchantID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var trackResponse models.TokenResponse
	json.Unmarshal([]byte(res), &trackResponse)
	response.OkWithDetailed(trackResponse, "OK", c)
}

// create payment
func (p *LLPayAPI) CreatePayment(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	var orders websiteReq.CreateLianLianPayType
	if err := c.ShouldBindJSON(&orders); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(orders)
	switch orders.Method {
	case "llpay":
		LLPay(c, orders, userId)
	default:
		response.FailWithMessage("payment method not supported", c)
	}
}

func LLPay(c *gin.Context, orders websiteReq.CreateLianLianPayType, userId string) {

	// 1. orderid find order
	order_desc, err := ordersService.GetOneOrderByIDDB(orders.OrderId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order_products, err := ordersService.GetOneOrderByIDProductsDB(orders.OrderId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	order_address, err := addressService.GetAddressById(userId, order_desc.ShippingAddressID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order_billing_address, err := billingAddressService.GetBillingAddressByUserID(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	url := lianlianpay.BaseUrl + fmt.Sprintf("/merchants/%s/payments", lianlianpay.MerchantID)

	ordersStruct := requestLLPay(orders.CardToken, orders.OrderId, order_desc, order_products, order_address, order_billing_address)

	timestamp := time.Now().Format("20060102150405")
	ordersStruct.MerchantOrder.MerchantOrderTime = timestamp

	fmt.Printf("ordersStruct: %+v\n", ordersStruct)

	res, err := lianlianpayService.CreatePayment(url, ordersStruct, timestamp)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var trackResponse models.PaymentResponseData
	json.Unmarshal([]byte(res), &trackResponse)
	if trackResponse.ReturnCode == "SUCCESS" {
		SaveLLPayPayment(trackResponse, userId, orders.OrderId)
		order_desc.PaymentStatus = "PS"
	} else {
		SaveLLPayPaymentError(trackResponse, userId, orders.OrderId)
		order_desc.PaymentStatus = "WF"
	}
	err = ordersService.UpdateOrdersProductDB(order_desc)
	if err != nil {
		fmt.Println(err)
	}
	response.OkWithDetailed(trackResponse, "OK", c)
}

// save payment information to database
func SaveLLPayPayment(trackResponse models.PaymentResponseData, userId string, orderId string) {
	payment := website.PaymentLlPay{
		OrderId:                orderId,
		UserId:                 userId,
		TraceId:                trackResponse.TraceId,
		LlTransactionId:        trackResponse.Order.LlTransactionId,
		MerchantTransactionId:  trackResponse.Order.MerchantTransactionId,
		PaymentCurrencyCode:    trackResponse.Order.PaymentData.PaymentCurrencyCode,
		PaymentAmount:          trackResponse.Order.PaymentData.PaymentAmount,
		SettlementCurrencyCode: trackResponse.Order.PaymentData.SettlementCurrencyCode,
		SettlementAmount:       trackResponse.Order.PaymentData.SettlementAmount,
		ExchangeRate:           trackResponse.Order.PaymentData.ExchangeRate,
		PaymentTime:            trackResponse.Order.PaymentData.PaymentTime,
		PaymentStatus:          trackResponse.Order.PaymentData.PaymentStatus,
		PaymentUrl:             trackResponse.Order.PaymentUrl,
		KeyValue:               trackResponse.Order.Key,
	}
	PaymentLlPayService.CreatePaymentLLPayRecord(payment)
}

// record payment error
func SaveLLPayPaymentError(trackResponse models.PaymentResponseData, userId string, orderId string) {
}

// create payment order string
func requestLLPay(cardToken string, orderID string, order_desc website.OrdersType, order_products []website.OrdersProduct, order_address website.ClientAddress, order_billing_address website.BillingAddress) website.LianLianPay {
	return website.LianLianPay{
		MerchantTransactionId: orderID,
		MerchantId:            lianlianpay.MerchantID,
		SubMerchantId:         lianlianpay.SubMerchantId,
		NotificationUrl:       "https://baidu.com",
		RedirectUrl:           "http://localhost:3008/payment",
		Country:               "US",
		PaymentMethod:         "inter_credit_card",
		MerchantOrder:         setMerchantOrder(orderID, order_desc, order_products, order_address),
		Customer:              setCustomer(order_address),
		PaymentData:           setPaymentData(cardToken, order_billing_address),
	}

}
func setMerchantOrder(orderID string, order_desc website.OrdersType, order_products []website.OrdersProduct, order_address website.ClientAddress) website.MerchantOrder {
	return website.MerchantOrder{
		MerchantOrderId:   orderID,
		MerchantOrderTime: "20060102150405",
		OrderAmount:       order_desc.TotalPrice,
		OrderCurrencyCode: "USD",
		Products:          setProducts(order_products),
		Shipping:          setShippping(order_address),
	}
}
func setProducts(order_products []website.OrdersProduct) []website.LLPayProduct {
	var products []website.LLPayProduct
	for _, product := range order_products {
		products = append(products, website.LLPayProduct{
			ProductId:        product.ProductID,
			Name:             product.ProductID + "-" + product.Size,
			Description:      product.Title,
			Price:            product.Price,
			Quantity:         product.Quantity,
			Category:         "Press on nails",
			Sku:              "SN" + product.ProductID + "-" + product.Size,
			Url:              "https://ftanails.com",
			ShippingProvider: "YANWEN",
		})
	}
	return products
}

func setShippping(order_address website.ClientAddress) website.LLPayShipping {
	return website.LLPayShipping{
		Name:    order_address.FirstName + " " + order_address.LastName,
		Cycle:   "other",
		Address: setCustomerAddress(order_address),
	}
}

func setCustomer(order_address website.ClientAddress) website.LLPayCustomer {
	return website.LLPayCustomer{
		CustomerType: "I",
		FullName:     order_address.FirstName + " " + order_address.LastName,
		FirstName:    order_address.FirstName,
		LastName:     order_address.LastName,
		Address:      setCustomerAddress(order_address),
	}
}

func setCustomerAddress(order_address website.ClientAddress) website.LLPayBillingAddress {
	return website.LLPayBillingAddress{
		Line1:      order_address.Line1,
		Line2:      order_address.Line2,
		City:       order_address.City,
		State:      order_address.State,
		PostalCode: order_address.PostalCode,
		Country:    order_address.Country,
		District:   "",
	}
}
func setPaymentData(cardToken string, order_billing_address website.BillingAddress) website.LLPayPaymentData {
	return website.LLPayPaymentData{
		Installments: 1,
		Card:         setCard(cardToken, order_billing_address),
	}
}

func setCard(cardToken string, billing_address website.BillingAddress) website.LLPayCard {
	return website.LLPayCard{
		CardToken:      cardToken,
		HolderName:     "John Doe",
		BillingAddress: setBillngAddress(billing_address),
	}
}
func setBillngAddress(billing_address website.BillingAddress) website.LLPayBillingAddress {

	fmt.Println(billing_address)
	return website.LLPayBillingAddress{
		Line1:      "123 Main St",
		Line2:      "",
		City:       "New York",
		State:      "NY",
		PostalCode: "10001",
		Country:    "US",
		District:   "",
	}
}
