package website

import "leiserv/global"

type ProductDetail struct {
	global.DATE_MODEL
	ProductID string `json:"productId" gorm:"index;comment:商品ID"`
	Lang      string `json:"lang" gorm:"comment:语言"`
	Content   string `json:"content" gorm:"comment:商品详情内容"`
}

func (ProductDetail) TableName() string {
	return "product_detail"
}
