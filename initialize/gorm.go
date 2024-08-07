package initialize

import (
	"os"

	"leiserv/global"
	"leiserv/models/example"
	"leiserv/models/system"

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
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
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
