package website

import "leiserv/global"

type Tag struct {
	global.DATE_MODEL
	Type     uint   `json:"type" gorm:"type:tinyint UNSIGNED NOT NULL;not null;comment:'商品标签类型, 0:分类标签, 1:表示商品标签, 2:表示商品sku类型'"`
	ParentID uint   `json:"parent_id" gorm:"type:bigint UNSIGNED default:0;comment:'父级标签ID'"`
	Title    string `json:"title" gorm:"type:VARCHAR(255);not null;comment:'商品标签名称'"`
	TitleEn  string `json:"title_en" gorm:"type:VARCHAR(255);comment:'商品标签名称EN'"`
	Value    string `json:"value" gorm:"type:VARCHAR(128);comment:'商品标签值'"`
	ValueCm  string `json:"value_cm" gorm:"type:VARCHAR(128);comment:'商品标签值EN'"`
}

func (Tag) TableName() string {
	return "tags"
}

type TagInfo struct {
	global.DATE_MODEL
	TagID     uint   `json:"tag_id" gorm:"type:bigint UNSIGNED NOT NULL;not null;comment:'商品标签ID'"`
	ProductID string `json:"product_id" gorm:"type:VARCHAR(64);not null;comment:'商品ID'"`
	Title     string `json:"title" gorm:"type:VARCHAR(255);not null;comment:'商品标签名称'"`
	Value     string `json:"value" gorm:"type:VARCHAR(128);comment:'商品标签值'"`
	Level     uint   `json:"level" gorm:"type:tinyint UNSIGNED NOT NULL DEFAULT 1;not null;default:1;comment:'商品标签级别, 级别越高, 优先级越高, 1-9'"`
}

type TagList struct {
	Tag
	Children []Tag `gorm:"-"`
}

func (TagInfo) TableName() string {
	return "tag_info"
}
