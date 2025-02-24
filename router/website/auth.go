package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (a *AuthRouter) InitAuthRouter(PublicRouter *gin.RouterGroup, PrivatRouter *gin.RouterGroup) {
	authRouter := PublicRouter.Group("auth")

	privatRouter := PrivatRouter.Group("myself")

	captchaAPI := v1.ApiGroupApp.WebSiteAPIPack.CaptchaAPI
	authAPI := v1.ApiGroupApp.WebSiteAPIPack.AuthAPI
	{
		authRouter.POST("captcha", captchaAPI.GetCaptcha)
		authRouter.POST("signup", authAPI.PostSignup)
		authRouter.POST("signin", authAPI.PostLogin)
	}

	{
		privatRouter.POST("info", authAPI.GetUserInfo)
		privatRouter.POST("update", authAPI.SetUserInfo)
		privatRouter.POST("order", authAPI.GetUserInfo)
	}
}
