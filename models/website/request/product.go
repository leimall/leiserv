package request

// "title": "aabb",
// "desction": "aaabbbsdfsaf",
// "seoKeywords": "ssdf",
// "seoDescription": "fasfsdf",
// "price": 100,
// "priceOff": 1,
// "mainImg": "",
// "stock": "2000"
type CreateProductRequest struct {
	Title          string  `json:"title" form:"title" gorm:"column:title;"`
	Description    string  `json:"desction" form:"desction" gorm:"column:description;"`
	SeoKeywords    string  `json:"seoKeywords" form:"seoKeywords" gorm:"column:seo_keywords;"`
	SeoDescription string  `json:"seoDescription" form:"seoDescription" gorm:"column:seo_description;"`
	Price          float64 `json:"price" form:"price" gorm:"column:price;"`
	PriceOff       float64 `json:"priceOff" form:"priceOff" gorm:"column:price_off;"`
	MainImg        string  `json:"mainImg" form:"mainImg" gorm:"column:main_img;"`
	Stock          uint64  `json:"stock" form:"stock" gorm:"column:stock;"`
}
