package website

import (
	"fmt"
	"leiserv/global"
	"leiserv/models/common/response"
	"leiserv/models/system"
	webauthReq "leiserv/models/website/request"
	webauthRes "leiserv/models/website/response"
	websiteType "leiserv/models/website/types"
	"leiserv/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AuthAPI struct{}

func (a *AuthAPI) PostSignup(c *gin.Context) {
	var r webauthReq.Signup
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.WebVerify(r, utils.WebRegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &websiteType.ClientUser{
		Email:    r.Email,
		Password: r.Password,
		Username: r.Username,
	}
	userReturn, err := userService.SignUp(*user)
	if err != nil {
		fmt.Println("注册失败:", err)
		response.FailWithDetailed(webauthRes.ClientResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(webauthRes.ClientResponse{User: userReturn}, "注册成功", c)
}

func (a *AuthAPI) PostLogin(c *gin.Context) {

	var l webauthReq.Signin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.MALL_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.MALL_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &websiteType.ClientUser{Email: l.Email, Password: l.Password}
		user, err := userService.SignIn(u)
		if err != nil {
			global.MALL_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.MALL_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		a.TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

func (b *AuthAPI) TokenNext(c *gin.Context, user websiteType.ClientUser) {
	j := &utils.WebJWT{SigningKey: []byte(global.MALL_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateWebClaims(websiteType.BaseClient{
		UUID:       user.UUID,
		ID:         user.ID,
		UserId:     user.UserId,
		Email:      user.Email,
		Username:   user.Username,
		Permission: user.Permission,
	})

	token, err := j.CreateWebToken(claims)
	if err != nil {
		global.MALL_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.MALL_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(webauthRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtWebService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtWebService.SetRedisJWT(token, user.Username); err != nil {
			global.MALL_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(webauthRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.MALL_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtWebService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtWebService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(webauthRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

func (a *AuthAPI) GetUserInfo(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	fmt.Println("userId:", userId)
	ReqUser, err := userService.GetUserInfo(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取用户信息成功", c)
}

func (a *AuthAPI) SetUserInfo(c *gin.Context) {
	var user webauthReq.UserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetWebUserID(c)
	fmt.Println("userId:", userId)
	err = userService.UpdateUserInfo(websiteType.ClientUser{
		UserId:    userId,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Username:  user.Username,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("修改用户信息成功", c)
}
