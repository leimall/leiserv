package website

import (
	"leiserv/global"
)

type JwtBlacklist struct {
	global.DATE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (JwtBlacklist) TableName() string {
	return "web_jwtblacklists"
}
