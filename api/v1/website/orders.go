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
		PaymentMethod:     orders.PaymentMethod,
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

// func (p *OrdersApi) GetOrderDetail(c *gin.Context) {
// 	var orderid string = c.Param("id")
// 	order, err := ordersService.GetOrderDetailDB(orderid)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	response.OkWithDetailed(order, "OK", c)
// }

// func (p *OrdersApi) UpdateOrder(c *gin.Context) {
// 	var order webauthReq.OrdersRequest
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err := ordersService.UpdateOrderDB(order)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	response.OkWithMessage("OK", c)
// }
// func (p *OrdersApi) DeleteOrder(c *gin.Context) {
// 	var orderid string = c.Param("id")
// 	err := ordersService.DeleteOrderDB(orderid)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	response.OkWithMessage("OK", c)
// }
