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
		categoryRouter.POST("create", categoryAPI.CreateCategory)
		// categoryRouter.POST("update", categoryAPI.UpdateCategory)
		// categoryRouter.POST("delete", categoryAPI.DeleteCategory)
		categoryRouter.POST("list", categoryAPI.GetCategory)
	}

}
