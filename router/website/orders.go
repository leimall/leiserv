package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type OrdersRouter struct{}

func (e *OrdersRouter) InitOrdersRouter(RouterPrivate *gin.RouterGroup) {
	ordersRouter := RouterPrivate.Group("orders")
	baseApi := v1.ApiGroupApp.WebSiteAPIPack.OrdersApi
	{
		ordersRouter.GET("list", baseApi.GetOrdersList)
		ordersRouter.GET("orderid", baseApi.GetOrdersId)
		ordersRouter.POST("create", baseApi.CreateOrders)
		ordersRouter.POST("update", baseApi.UpdateOrder)
		ordersRouter.POST("status", baseApi.UpdateOrderStatus)
		ordersRouter.GET("myself", baseApi.GetMyselfOrders)
		ordersRouter.GET("/:id", baseApi.GetOneOrderById)
	}
}
