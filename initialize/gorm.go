package initialize

import (
	"os"

	"leiserv/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.MALL_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.MALL_DB
	err := db.AutoMigrate()
	if err != nil {
		global.MALL_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel(db)

	if err != nil {
		global.MALL_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.MALL_LOG.Info("register table success")
}
