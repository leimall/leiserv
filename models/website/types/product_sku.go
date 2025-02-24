package website

import "leiserv/global"

type SkuInfo struct {
	SkuItem
	List []SkuItem `gorm:"-"`
	global.DATE_MODEL
}

func (SkuItem) TableName() string {
	return "product_sku"
}

type SkuItem struct {
	TagID     uint    `json:"tagId" gorm:"index;comment:'SKIU标签ID'"`
	ProductID string  `json:"productId" gorm:"index;comment:'商品ID'"`
	ParentID  uint    `json:"parentId" gorm:"index;comment:'父ID'"`
	Title     string  `json:"title" gorm:"comment:'分类名称'"`
	Stock     uint    `json:"stock" gorm:"comment:'库存'"`
	Price     float64 `json:"price" gorm:"comment:'价格'"`
	PriceOff  float64 `json:"priceOff" gorm:"comment:'折扣'"`
	MainImg   string  `json:"mainImg" gorm:"comment:'主图'"`
	global.DATE_MODEL
}
