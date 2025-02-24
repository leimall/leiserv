package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type LiangLiangPayRouter struct{}

func (e *OrdersRouter) InitLianLiangPayRouter(RouterPrivate *gin.RouterGroup) {
	llpayRouter := RouterPrivate.Group("llpay")
	llpayAPI := v1.ApiGroupApp.WebSiteAPIPack.LLPayAPI
	{
		llpayRouter.GET("token", llpayAPI.GetToken)

		llpayRouter.POST("payment", llpayAPI.CreatePayment)
	}
}
