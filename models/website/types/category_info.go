package website

import "leiserv/global"

type CategoryInfo struct {
	global.DATE_MODEL
	ProductID string `json:"product_id" gorm:"index;comment:'商品ID'"`
	Title     string `json:"title" gorm:"comment:'分类名称'"`
	Value     string `json:"value" gorm:"comment:'分类值'"`
	Level     int    `json:"level" gorm:"comment:'分类级别'"`
}

func (CategoryInfo) TableName() string {
	return "category_info"
}

type Category struct {
	ID    uint   `gorm:"primary_key;auto_increment;comment:'分类ID'"`
	Title string `gorm:"comment:'分类名称'"`
}

func (Category) TableName() string {
	return "category"
}
