package utils

import (
	"leiserv/global"
	systemReq "leiserv/models/system/request"
	website "leiserv/models/website/types"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
)

func ClearWebToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func SetWebToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetWebToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

func GetHeaderUserId(c *gin.Context) string {
	userId, _ := c.Cookie("x-user-id")
	if userId == "" {
		userId = c.Request.Header.Get("x-user-id")
	}
	return userId
}

func GetWebClaims(c *gin.Context) (*website.JWTClaims, error) {
	token := GetWebToken(c)
	j := NewWebJWT()
	claims, err := j.ParseWebToken(token)
	if err != nil {
		global.MALL_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetWebUserID(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetWebClaims(c); err != nil {
			return ""
		} else {
			return cl.UserId
		}
	} else {
		waitUse := claims.(*website.JWTClaims)
		return waitUse.UserId
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetWebUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetWebClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}
