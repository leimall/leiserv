package request

type ClientAddress struct {
	ID          uint   `json:"ID"`
	UserId      string `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street1     string `json:"street1"`
	Email       string `json:"email"`
	City        string `json:"city"`
	State       string `json:"state"`
	CountryCode string `json:"countryCode"`
	Country     string `json:"country"`
	ZipCode     string `json:"zipCode"`
	Phone       string `json:"phone"`
	IsDefault   int8   `json:"isDefault"`
}

type AddressID struct {
	ID int64 `json:"id"`
}

// billing address request
// UserId       string `json:"userId" gorm:"index;comment:用户ID"`
// FirstName    string `json:"firstName" gorm:"comment:持卡人名字"`
// LastName     string `json:"lastName" gorm:"comment:持卡人姓氏"`
// HolderName   string `json:"holderName" gorm:"comment:持卡人全名"`
// Line1        string `json:"line1" gorm:"comment:地址行1"`
// Line2        string `json:"line2" gorm:"comment:地址行2"`
// City         string `json:"city" gorm:"comment:城市"`
// State        string `json:"state" gorm:"comment:州/省份"`
// PostalCode   string `json:"postalCode" gorm:"comment:邮政编码"`
// District     string `json:"district" gorm:"comment:行政区"`
// CardNumber   string `json:"cardNumber" gorm:"comment:卡号"`
// CardType     string `json:"cardType" gorm:"comment:卡类型:D=借记,C=信用卡"`
// BankCode     string `json:"bankCode" gorm:"comment:银行编码"`
// CardBrand    string `json:"cardBrand" gorm:"comment:信用卡品牌,例如Visa、MasterCard"`
// CardExpYear  string `json:"cardExpYear" gorm:"comment:卡年有效期"`
// CardExpMonth string `json:"cardExpMonth" gorm:"comment:卡月有效期"`
// Cvv          uint   `json:"cvv" gorm:"comment:卡安全码"`
// PhoneNumber  string `json:"phoneNumber" gorm:"comment:手机号，格式“+区号-手机号”"`
// Email        string `json:"email" gorm:"comment:邮箱地址"`
type BillingAddress struct {
	ID           uint   `json:"ID"`
	UserId       string `json:"userId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Line1        string `json:"line1"`
	Line2        string `json:"line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postalCode"`
	District     string `json:"district"`
	CardNumber   string `json:"cardNumber"`
	CardType     string `json:"cardType"`
	BankCode     string `json:"bankCode"`
	CardBrand    string `json:"cardBrand"`
	CardExpYear  string `json:"cardExpYear"`
	CardExpMonth string `json:"cardExpMonth"`
	Cvv          uint   `json:"cvv"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
}
