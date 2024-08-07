package website

import "leiserv/global"

// product_id varchar(16)  NOT NULL COMMENT '商品ID',
// title VARCHAR(255) NOT NULL COMMENT '商品名称',
// desction text COMMENT '商品描述',
// seo_keywords VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'SEO关键词',
// seo_description VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'SEO描述',
// price DECIMAL(12,2) UNSIGNED NOT NULL default 0.00 COMMENT '商品价格',
// price_off decimal(3,2) UNSIGNED NOT NULL DEFAULT 1.00 COMMENT '商品折扣',
// main_img VARCHAR(255) NOT NULL COMMENT '商品图片',
// stock INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
// sales INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品销量',
// status tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品状态 0:未上架, 1: 上架',
// is_delete tinyint unsigned NOT NULL DEFAULT 0 COMMENT '商品删除状态 0:未删除, 1: 已下架',

type Produce struct {
	global.DATE_MODEL
	ProductID      string  `json:"productId" gorm:"index;comment:商品ID"`
	Title          string  `json:"title" gorm:"index;comment:商品名称"`
	Desction       string  `json:"desction" gorm:"comment:商品描述"`
	SeoKeywords    string  `json:"seoKeywords" gorm:"comment:SEO关键词"`
	SeoDescription string  `json:"seoDescription" gorm:"comment:SEO描述"`
	Price          float64 `json:"price" gorm:"comment:商品价格"`
	PriceOff       float64 `json:"priceOff" gorm:"comment:商品折扣"`
	MainImg        string  `json:"mainImg" gorm:"comment:商品图片"`
	Stock          uint64  `json:"stock" gorm:"comment:商品库存"`
	Sales          uint64  `json:"sales" gorm:"comment:商品销量"`
	Status         int     `json:"status" gorm:"index;comment:商品状态 0:未上架, 1: 上架"`
	IsDelete       int     `json:"isDelete" gorm:"index;comment:商品删除状态 0:未删除, 1: 已下架"`
}

func (Produce) TableName() string {
	return "product"
}

type AllProduct struct {
	global.DATE_MODEL
	ProductID      string         `json:"productId" gorm:"index;comment:商品ID"`
	Title          string         `json:"title" gorm:"index;comment:商品名称"`
	Desction       string         `json:"desction" gorm:"comment:商品描述"`
	SeoKeywords    string         `json:"seoKeywords" gorm:"comment:SEO关键词"`
	SeoDescription string         `json:"seoDescription" gorm:"comment:SEO描述"`
	Price          float64        `json:"price" gorm:"comment:商品价格"`
	PriceOff       float64        `json:"priceOff" gorm:"comment:商品折扣"`
	MainImg        string         `json:"mainImg" gorm:"comment:商品图片"`
	Stock          uint64         `json:"stock" gorm:"comment:商品库存"`
	Sales          uint64         `json:"sales" gorm:"comment:商品销量"`
	Status         int            `json:"status" gorm:"index;comment:商品状态 0:未上架, 1: 上架"`
	IsDelete       int            `json:"isDelete" gorm:"index;comment:商品删除状态 0:未删除, 1: 已下架"`
	Category       []CategoryInfo `gorm:"-"`
	ImageList      []ProductImg   `gorm:"-"`
}

func (AllProduct) TableName() string {
	return "product"
}
