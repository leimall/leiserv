package response

import websiteTypes "leiserv/models/website/types"

type ClientResponse struct {
	User websiteTypes.ClientUser `json:"user"`
}

type LoginResponse struct {
	User      websiteTypes.ClientUser `json:"user"`
	Token     string                  `json:"token"`
	ExpiresAt int64                   `json:"expiresAt"`
}
