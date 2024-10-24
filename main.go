package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	"leiserv/core"
	"leiserv/global"
	"leiserv/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       MALL Swagger API接口文档
// @version                     v2.6.7
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.MALL_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.MALL_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.MALL_LOG)
	global.MALL_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.MALL_DB != nil {
		// initialize.RegisterTables() // 初始化表

		// initPayment() // 初始化支付接口

		// initShipping() // 初始化物流接口

		// 程序结束前关闭数据库链接
		db, _ := global.MALL_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
