package core

import (
	"fmt"
	"leiserv/global"
	"leiserv/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.MALL_CONFIG.System.UseMultipoint || global.MALL_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	if global.MALL_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// // 从db加载jwt数据
	// if global.MALL_DB != nil {
	// 	system.LoadAll()
	// }

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.MALL_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.MALL_LOG.Info("server run success on ", zap.String("address", address))

	global.MALL_LOG.Error(s.ListenAndServe().Error())
}
