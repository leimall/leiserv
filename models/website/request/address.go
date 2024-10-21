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
