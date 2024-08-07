package request

type Address struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	City        string `json:"city"`
	Region      string `json:"region"`
	CountryCode string `json:"countryCode"`
	Country     string `json:"country"`
	ZipCode     string `json:"zipCode"`
	Phone       string `json:"phone"`
	IsDefault   int8   `json:"isDefault"`
}
