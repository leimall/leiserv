package website

import "leiserv/global"

type ClientAddress struct {
	global.DATE_MODEL
	UserId      string `json:"userId" gorm:"index;comment:userID"`
	FirstName   string `json:"firstName" gorm:"comment:名子"`
	LastName    string `json:"lastName" gorm:"comment:姓氏"`
	Line1       string `json:"line1" gorm:"comment:地址行1"`
	Line2       string `json:"line2" gorm:"comment:地址行2"`
	Email       string `json:"email" gorm:"comment:Email"`
	City        string `json:"city" gorm:"comment:city"`
	State       string `json:"state" gorm:"comment:省/区"`
	CountryName string `json:"countryName" gorm:"comment:国家代码"`
	Country     string `json:"country" gorm:"comment:国家"`
	PostalCode  string `json:"postalCode" gorm:"comment:邮政编码"`
	Phone       string `json:"phone" gorm:"comment:收货人电话"`
	Mark        string `json:"mark" gorm:"comment:备注"`
	IsDefault   int8   `json:"isDefault" gorm:"comment:是否默认地址"`
}

func (ClientAddress) TableName() string {
	return "client_address"
}
