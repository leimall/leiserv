package request

import website "leiserv/models/website/types"

type OrdersRequest struct {
	OrderID           string                  `json:"orderId"`
	UserID            string                  `json:"userId"`
	TotalPrice        float64                 `json:"totalPrice"`
	PaymentMethod     string                  `json:"paymentMethod"`
	PaymentStatus     string                  `json:"paymentStatus"`
	OrderStatus       string                  `json:"orderStatus"`
	ShippingMethod    string                  `json:"shippingMethod"`
	ShippingPrice     float64                 `json:"shippingPrice"`
	ShippingAddressID uint64                  `json:"shippingAddressId"`
	Products          []website.OrdersProduct `json:"products"`
}

type OrdersRequestType struct {
	website.OrdersType
}
