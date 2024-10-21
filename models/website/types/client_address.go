package website

import "leiserv/global"

type ClientAddress struct {
	global.DATE_MODEL
	UserId      string `json:"userId" gorm:"index;comment:userID"`
	FirstName   string `json:"firstName" gorm:"comment:名子"`
	LastName    string `json:"lastName" gorm:"comment:姓氏"`
	Street1     string `json:"street1" gorm:"comment:街道 1"`
	Email       string `json:"email" gorm:"comment:Email"`
	City        string `json:"city" gorm:"comment:city"`
	State       string `json:"state" gorm:"comment:省/区"`
	CountryCode string `json:"countryCode" gorm:"comment:国家代码"`
	Country     string `json:"country" gorm:"comment:国家"`
	ZipCode     string `json:"zipCode" gorm:"comment: 邮编"`
	Phone       string `json:"phone" gorm:"comment:收货人电话"`
	IsDefault   int8   `json:"isDefault" gorm:"comment:是否默认地址"`
}

func (ClientAddress) TableName() string {
	return "client_address"
}
