package website

import "leiserv/global"

type BillingAddress struct {
	global.DATE_MODEL
	UserId       string `json:"userId" gorm:"index;comment:用户ID"`
	FirstName    string `json:"firstName" gorm:"comment:持卡人名字"`
	LastName     string `json:"lastName" gorm:"comment:持卡人姓氏"`
	HolderName   string `json:"holderName" gorm:"comment:持卡人全名"`
	Line1        string `json:"line1" gorm:"comment:地址行1"`
	Line2        string `json:"line2" gorm:"comment:地址行2"`
	City         string `json:"city" gorm:"comment:城市"`
	State        string `json:"state" gorm:"comment:州/省份"`
	PostalCode   string `json:"postalCode" gorm:"comment:邮政编码"`
	District     string `json:"district" gorm:"comment:行政区"`
	CardNumber   string `json:"cardNumber" gorm:"comment:卡号"`
	CardType     string `json:"cardType" gorm:"comment:卡类型:D=借记,C=信用卡"`
	BankCode     string `json:"bankCode" gorm:"comment:银行编码"`
	CardBrand    string `json:"cardBrand" gorm:"comment:信用卡品牌,例如Visa、MasterCard"`
	CardExpYear  string `json:"cardExpYear" gorm:"comment:卡年有效期"`
	CardExpMonth string `json:"cardExpMonth" gorm:"comment:卡月有效期"`
	Cvv          uint   `json:"cvv" gorm:"comment:卡安全码"`
	PhoneNumber  string `json:"phoneNumber" gorm:"comment:手机号，格式“+区号-手机号”"`
	Email        string `json:"email" gorm:"comment:邮箱地址"`
	LlpayToken   string `json:"llpayToken" gorm:"comment:连连支付token"`
}

func (BillingAddress) TableName() string {
	return "billing_address"
}
