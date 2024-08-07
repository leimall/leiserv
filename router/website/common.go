package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type CommonRouter struct{}

func (c *CommonRouter) InitCommonRouter(publicRouter *gin.RouterGroup) {

	commonRouter := publicRouter.Group("common")
	commonAPI := v1.ApiGroupApp.WebSiteAPIPack.CommonAPI

	{
		commonRouter.GET("constry", commonAPI.GetCountries)
		commonRouter.GET("city", commonAPI.GetCity)
	}
}
