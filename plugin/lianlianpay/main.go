package lianlianpay

import (
	"leiserv/plugin/lianlianpay/global"
	"leiserv/plugin/lianlianpay/router"

	"github.com/gin-gonic/gin"
)

type llPayPlugin struct{}

func CreateLLPayPlug(LLPubKey, PrivateKey, MerchantID string) *llPayPlugin {
	global.GlobalConfig.LLPubKey = LLPubKey
	global.GlobalConfig.PrivateKey = PrivateKey
	global.GlobalConfig.MerchantID = MerchantID
	return &llPayPlugin{}
}

func (*llPayPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitLLPayRouter(group)
}

func (*llPayPlugin) RouterPath() string {
	return "llpay"
}
