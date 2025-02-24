package website

import "leiserv/global"

type CartItem struct {
	global.DATE_MODEL
	UserId    string  `json:"user_id" gorm:"index;comment:userID"`
	ProductID string  `json:"product_id" gorm:"index;comment:商品ID"`
	Title     string  `json:"title" gorm:"index;comment:商品名称"`
	Price     float64 `json:"price" gorm:"comment:商品价格"`
	MainImg   string  `json:"main_img" gorm:"comment:商品图片"`
	Stock     uint64  `json:"stock" gorm:"comment:商品库存"`
	OldPrice  float64 `json:"old_price" gorm:"comment:商品原价"`
	PriceOff  uint8   `json:"price_off" gorm:"comment:商品折扣"`
	Quantity  uint64  `json:"quantity" gorm:"comment:商品数量"`
	SizeTitle string  `json:"size_title" gorm:"comment:商品尺寸标题"`
	Color     string  `json:"color" gorm:"comment:商品颜色"`
	Size      string  `json:"size" gorm:"comment:商品尺寸"`
	Shape     string  `json:"shape" gorm:"comment:商品甲型形状"`
	Status    uint8   `json:"status" gorm:"comment:商品状态, 04,表示删除. 1.正常"`
	UniqueId  string  `json:"unique_id" gorm:"comment:唯一标识"`
}

func (CartItem) TableName() string {
	return "cart"
}
