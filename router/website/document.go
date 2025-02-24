package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type DocumentRouter struct{}

func (a *DocumentRouter) InitDocumentRouter(Router *gin.RouterGroup) {

	documentRouter := Router.Group("document")

	documnetAPI := v1.ApiGroupApp.WebSiteAPIPack.DocumentAPI

	{
		documentRouter.GET("title", documnetAPI.GetDocumentList)
	}

}
