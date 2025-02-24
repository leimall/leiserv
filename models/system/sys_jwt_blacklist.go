package system

import (
	"leiserv/global"
)

type JwtBlacklist struct {
	global.DATE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
