package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (a *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category")

	categoryAPI := v1.ApiGroupApp.WebSiteAPIPack.CategoryAPI
	{
		categoryRouter.GET("list", categoryAPI.GetCategoryList)
		categoryRouter.GET("style", categoryAPI.GetStyleList)
		categoryRouter.GET("shape", categoryAPI.GetShapeList)
		categoryRouter.GET("menu", categoryAPI.GetMenuList)
	}

}
