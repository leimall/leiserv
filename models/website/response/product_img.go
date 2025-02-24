package response

import website "leiserv/models/website/types"

type ProductImgResponse struct {
	File website.ProductImg `json:"file"`
}
