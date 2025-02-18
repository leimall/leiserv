package website

import "leiserv/global"

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

type ProductListItme struct {
	Produce
	Review ProductReviewsItem `gorm:"-"`
}

type AllProduct struct {
	Produce
	Brand     ProductBrandInfo   `gorm:"-"`
	Category  []CategoryInfo     `gorm:"-"`
	ImageList []ProductImg       `gorm:"-"`
	Tags      []TagInfo          `gorm:"-"`
	Sku       SkuInfo            `gorm:"-"`
	Detail    []ProductDetail    `gorm:"-"`
	Review    ProductReviewsItem `gorm:"-"`
}

func (AllProduct) TableName() string {
	return "product"
}
