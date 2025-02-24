package core

import (
	"fmt"
	"leiserv/core/internal"
	"leiserv/global"
	"leiserv/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.MALL_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.MALL_CONFIG.Zap.Director)
		_ = os.Mkdir(global.MALL_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.MALL_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.MALL_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
