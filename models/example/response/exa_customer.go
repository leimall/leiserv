package response

import "leiserv/models/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
