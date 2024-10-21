package middleware

import (
	"errors"
	"leiserv/global"
	"leiserv/models/system"
	"leiserv/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"leiserv/models/common/response"
	"leiserv/service"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.WebsiteServiceGroup.JWTWebService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetWebToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.NoAuth("您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewWebJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseWebToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth("授权已过期", c)
				utils.ClearWebToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearWebToken(c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.WebParseDuration(global.MALL_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateWebTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseWebToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.MALL_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.MALL_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
