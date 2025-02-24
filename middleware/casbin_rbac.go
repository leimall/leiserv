package middleware

import (
	"strings"

	"leiserv/global"
	"leiserv/models/common/response"
	"leiserv/service"
	"leiserv/utils"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.WebsiteServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetWebClaims(c)
		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.MALL_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.Permission
		e := casbinService.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
