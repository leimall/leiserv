package website

import (
	"leiserv/global"

	"github.com/gofrs/uuid/v5"
)

type ClientUser struct {
	global.DATE_MODEL
	UUID       uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`   // 用户UUID
	UserId     string    `json:"userId" gorm:"index;comment:userID"` // 用户UUID
	Introducer string    `json:"introducer" gorm:"index;comment:介绍人 id"`
	Level      uint      `json:"level" gorm:"comment:用户等级"`
	Permission string    `json:"permission" gorm:"comment:用户权限"`
	Username   string    `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password   string    `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName   string    `json:"nickName" gorm:"comment:用户昵称"`                    // 用户昵称
	HeaderImg  string    `json:"headerImg" gorm:"comment:用户头像"`                   // 用户头像
	Phone      string    `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email      string    `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable     int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (ClientUser) TableName() string {
	return "client_info"
}
