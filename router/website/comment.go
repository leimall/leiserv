package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (r *CommentRouter) InitCommentRouter(PublicRouter *gin.RouterGroup, PrivatRouter *gin.RouterGroup) {
	pubRouter := PublicRouter.Group("/product")
	prvRouter := PrivatRouter.Group("/product")

	commentAPI := v1.ApiGroupApp.WebSiteAPIPack.AuthAPI

	{
		pubRouter.GET("comment", commentAPI.GetUserInfo)

	}
	{
		prvRouter.POST("comment", commentAPI.GetUserInfo)
		prvRouter.PUT("comment", commentAPI.GetUserInfo)
		prvRouter.DELETE("comment", commentAPI.GetUserInfo)
	}

}
