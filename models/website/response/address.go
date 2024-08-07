package response

import website "leiserv/models/website/types"

type Address struct {
	Address website.ClientAddress `json:"address"`
}
