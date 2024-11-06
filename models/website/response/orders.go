package response

import website "leiserv/models/website/types"

type OrderOnePesponse struct {
	Order website.OrdersType `json:"order"`
}
