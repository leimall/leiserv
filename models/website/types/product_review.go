package website

import "leiserv/global"

type ProductReviewsItem struct {
	ProductID string  `json:"productId" gorm:"index;comment:'商品ID'"`
	Star1     uint64  `json:"star1" gorm:"comment:'星级1'"`
	Star2     uint64  `json:"star2" gorm:"comment:'星级2'"`
	Star3     uint64  `json:"star3" gorm:"comment:'星级3'"`
	Star4     uint64  `json:"star4" gorm:"comment:'星级4'"`
	Star5     uint64  `json:"star5" gorm:"comment:'星级5'"`
	Total     uint64  `json:"total" gorm:"comment:'总评数'"`
	Reviews   uint64  `json:"reviews" gorm:"comment:'总评论数'"`
	Average   float64 `json:"average" gorm:"comment:'平均分'"`
	global.DATE_MODEL
}

func (ProductReviewsItem) TableName() string {
	return "product_reviews"
}
