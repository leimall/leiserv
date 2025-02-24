package website

import (
	"fmt"
	"leiserv/global"
	website "leiserv/models/website/types"
)

type PaymentLLPayService struct{}

// CreatePaymentLLPayRecord creates a new payment record in the database for LLPAY payment gateway.
func (s *PaymentLLPayService) CreatePaymentLLPayRecord(record website.PaymentLlPay) {
	err := global.MALL_DB.Save(&record).Error
	fmt.Println(err)
}

// get payment record by order id and user id
func (s *PaymentLLPayService) GetPaymentLLPayRecord(orderID string, userID string) website.PaymentLlPay {
	var record website.PaymentLlPay
	global.MALL_DB.Where("order_id =? and user_id =?", orderID, userID).First(&record)
	return record
}
