package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type CartRouter struct{}

func (c *CartRouter) InitCartRouter(PublicRouter *gin.RouterGroup, PrivatRouter *gin.RouterGroup) {
	cartGroup := PublicRouter.Group("cart")
	cartAPI := v1.ApiGroupApp.WebSiteAPIPack.CartAPI
	{
		cartGroup.POST("add", cartAPI.AddCart)
		cartGroup.POST("update", cartAPI.UpdateCart)
		cartGroup.POST("delete", cartAPI.DeleteCart)
		cartGroup.POST("deleteone", cartAPI.DeleteCartOne)
		cartGroup.POST("list", cartAPI.GetCartList)
	}
}
