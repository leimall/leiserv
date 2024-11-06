package router

import (
	"leiserv/middleware"
	"leiserv/plugin/lianlianpay/api"

	"github.com/gin-gonic/gin"
)

type LLPayRouter struct{}

func (s *LLPayRouter) InitLLPayRouter(Router *gin.RouterGroup) {
	payRouter := Router.Use(middleware.OperationRecord())
	EmailApi := api.ApiGroupApp.LLPayAPI.EmailTest
	SendEmail := api.ApiGroupApp.LLPayAPI.SendEmail
	{
		payRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		payRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
