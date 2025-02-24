package website

import "leiserv/global"

type ProductBrand struct {
	global.DATE_MODEL
	ProductID  string `gorm:"column:product_id;type:varchar(64);not null;default:''" json:"product_id"`
	BrandID    uint   `gorm:"column:brand_id;type:bigint unsigned;not null;default:0" json:"brand_id"`
	BrandTitle string `gorm:"column:brand_title;type:varchar(128);not null;default:''" json:"brand_title"`
	ShapeID    uint   `gorm:"column:shape_id;type:bigint unsigned;not null;default:0" json:"shape_id"`
	ShapeTitle string `gorm:"column:shape_title;type:varchar(128);not null;default:''" json:"shape_title"`
	TagID      uint   `gorm:"column:tag_id;type:bigint unsigned;not null" json:"tag_id"`
}

type ProductBrandInfo struct {
	Info ProductBrand `gorm:"-"`
	Tags []TagList    `gorm:"-"`
}

func (ProductBrand) TableName() string {
	return "product_brand"
}
