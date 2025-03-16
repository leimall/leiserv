package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type OrdersRouter struct{}

func (e *OrdersRouter) InitOrdersRouter(RouterPublic *gin.RouterGroup, RouterPrivate *gin.RouterGroup) {
	ordersRouter := RouterPrivate.Group("orders")
	ordersPublicRouter := RouterPublic.Group("orders")
	baseApi := v1.ApiGroupApp.WebSiteAPIPack.OrdersApi
	{
		ordersRouter.GET("list", baseApi.GetOrdersList)
		ordersRouter.POST("create", baseApi.CreateOrders)
		ordersRouter.POST("update", baseApi.UpdateOrder)
		ordersRouter.POST("status", baseApi.UpdateOrderStatus)
		ordersRouter.GET("myself", baseApi.GetMyselfOrders)
		ordersRouter.GET("/:id", baseApi.GetOneOrderById)
	}

	// public router
	{
		ordersPublicRouter.GET("orderid", baseApi.GetOrdersId)
	}
}
