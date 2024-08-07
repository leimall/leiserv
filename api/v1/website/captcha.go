package website

import (
	"image/color"
	"leiserv/global"
	"leiserv/models/common/response"
	websiteRes "leiserv/models/website/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type CaptchaAPI struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (a *CaptchaAPI) GetCaptcha(c *gin.Context) {
	// TODO: Implement captcha generation
	openCaptcha := global.MALL_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.MALL_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	// driver := base64Captcha.NewDriverDigit(
	// 	global.MALL_CONFIG.Captcha.ImgHeight,
	// 	global.MALL_CONFIG.Captcha.ImgWidth,
	// 	global.MALL_CONFIG.Captcha.KeyLong,
	// 	0.7,
	// 	80,
	// )
	driverString := base64Captcha.NewDriverString(
		global.MALL_CONFIG.Captcha.ImgHeight,
		global.MALL_CONFIG.Captcha.ImgWidth,
		1,
		3,
		global.MALL_CONFIG.Captcha.KeyLong,
		"abcdefghijklmnopqrstuvwxyz",
		&color.RGBA{R: 240, G: 240, B: 240, A: 240},
		// &base64Captcha.FontsStorage([]string{"wqy-microhei.ttc"}),
		nil,
		[]string{"wqy-microhei.ttc"},
	)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driverString, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.MALL_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(websiteRes.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.MALL_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
