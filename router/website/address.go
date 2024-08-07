package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type AddressRouter struct{}

func (a *AddressRouter) InitAddressRouter(RouterPrivate *gin.RouterGroup) {
	addressPrivateRouter := RouterPrivate.Group("myself")

	addressAPI := v1.ApiGroupApp.WebSiteAPIPack.AddressAPI
	{
		addressPrivateRouter.GET("/address", addressAPI.GetAddress)
		addressPrivateRouter.POST("/address", addressAPI.CreateAddress)
		addressPrivateRouter.PUT("/address", addressAPI.CreateAddress)
		addressPrivateRouter.DELETE("/address", addressAPI.CreateAddress)
	}

}
