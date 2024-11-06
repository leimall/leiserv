package website

import (
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"leiserv/utils"

	"github.com/gin-gonic/gin"
)

type OrdersApi struct{}

func (p *OrdersApi) GetOrdersId(c *gin.Context) {
	orderid := utils.SnowflakeID()
	response.OkWithDetailed(orderid, "OK", c)
}

func (p *OrdersApi) GetOrdersList(c *gin.Context) {
	var pageinfo webauthReq.PageInfo
	err := c.ShouldBindQuery(&pageinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ordersService.GetOrdersListDB(pageinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (p *OrdersApi) CreateOrders(c *gin.Context) {
	var orders webauthReq.OrdersRequest
	if err := c.ShouldBindJSON(&orders); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	orderRecord := website.OrdersType{
		OrderID:           orders.OrderID,
		UserID:            orders.UserID,
		TotalPrice:        orders.TotalPrice,
		PaymentStatus:     orders.PaymentStatus,
		OrderStatus:       orders.OrderStatus,
		ShippingMethod:    orders.ShippingMethod,
		ShippingPrice:     orders.ShippingPrice,
		ShippingAddressID: orders.ShippingAddressID,
	}
	err := ordersService.CreateOrdersDB(orderRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ordersService.CreateOrdersProductDB(orders.Products)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("OK", c)
}

func (p *OrdersApi) UpdateOrder(c *gin.Context) {
	var orders website.OrdersType
	if err := c.ShouldBindJSON(&orders); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := ordersService.UpdateOrdersProductDB(orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}

// get myself orders list
func (p *OrdersApi) GetMyselfOrders(c *gin.Context) {
	var pageinfo webauthReq.PageInfo
	err := c.ShouldBindQuery(&pageinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ordersService.GetMyOrdersListDB(pageinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

// get one order by id
func (p *OrdersApi) GetOneOrderById(c *gin.Context) {
	orderid := c.Param("id")
	order, err := ordersService.GetOneOrderByIDDB(orderid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.Products, err = ordersService.GetOneOrderByIDProductsDB(orderid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(order, "OK", c)
}
