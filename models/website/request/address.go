package request

type ClientAddress struct {
	ID          uint   `json:"ID"`
	UserId      string `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Email       string `json:"email"`
	City        string `json:"city"`
	State       string `json:"state"`
	StateName   string `json:"stateName"`
	CountryName string `json:"countryName"`
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
	Phone       string `json:"phone"`
	Mark        string `json:"mark"`
	IsDefault   int8   `json:"isDefault"`
}

type AddressID struct {
	ID int64 `json:"id"`
}

type BillingAddress struct {
	ID           uint   `json:"ID"`
	UserId       string `json:"userId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Line1        string `json:"line1"`
	Line2        string `json:"line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
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
