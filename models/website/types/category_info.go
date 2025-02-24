package website

import "leiserv/global"

// product_id varchar(16)  NOT NULL COMMENT '商品ID',
// title VARCHAR(255) NOT NULL COMMENT '分类名称',
// parent_id bigint UNSIGNED NOT NULL default 0 COMMENT '父级分类ID',
// level tinyint UNSIGNED NOT NULL COMMENT '分类级别',
type CategoryInfo struct {
	global.DATE_MODEL
	ProductID string `gorm:"index;comment:'商品ID'"`
	Title     string `gorm:"comment:'分类名称'"`
	ParentID  int    `gorm:"default:0;comment:'父级分类ID'"`
	Level     int    `gorm:"comment:'分类级别'"`
}

func (CategoryInfo) TableName() string {
	return "category_info"
}

// id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '分类ID',
// title VARCHAR(255) NOT NULL COMMENT '分类名称',
type Category struct {
	ID    uint   `gorm:"primary_key;auto_increment;comment:'分类ID'"`
	Title string `gorm:"comment:'分类名称'"`
}

func (Category) TableName() string {
	return "category"
}
