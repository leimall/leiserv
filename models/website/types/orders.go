package website

import "leiserv/global"

const (
	// 支付状态常量
	PaymentStatusUnpaid   = "unpaid"
	PaymentStatusPaid     = "paid"
	PaymentStatusRefunded = "refunded"

	// 订单状态常量
	OrderStatusPending                    = "pending"
	OrderStatusProcessing                 = "processing"
	OrderStatusInProduction               = "in_production"
	OrderStatusProductionCompleted        = "production_completed"
	OrderStatusPendingQualityInspection   = "pending_quality_inspection"
	OrderStatusQualityInspectionCompleted = "quality_inspection_completed"
	OrderStatusPackaging                  = "packaging"
	OrderStatusPacked                     = "packed"
	OrderStatusPendingShipment            = "pending_shipment"
	OrderStatusShipping                   = "shipping"
	OrderStatusCompleted                  = "completed"
	OrderStatusCancelled                  = "cancelled"
	OrderStatusRefunded                   = "refunded"
)

type OrdersType struct {
	global.DATE_MODEL
	UserID            string          `json:"user_id" gorm:"comment:'用户ID'"`
	OrderID           string          `json:"order_id" gorm:"comment:'订单号'"`
	TotalPrice        float64         `json:"total_price" gorm:"comment:'订单总价格'"`
	Discount          float64         `json:"discount" gorm:"comment:'折扣金额'"`
	PaymentMethod     string          `json:"payment_method" gorm:"comment:'支付方式'"`
	PaymentStatus     string          `json:"payment_status" gorm:"comment:'支付状态'"`
	OrderStatus       string          `json:"order_status" gorm:"comment:'订单状态'"`
	ShippingCompanyID uint64          `json:"shipping_company_id" gorm:"comment:'快递公司ID'"`
	ShippingMethod    string          `json:"shipping_method" gorm:"comment:'配送方式'"`
	ShippingPrice     float64         `json:"shipping_price" gorm:"comment:'配送费用'"`
	ShippingAddressID uint64          `json:"shipping_address_id" gorm:"comment:'配送地址ID'"`
	TrackingNumber    string          `json:"tracking_number" gorm:"comment:'快递单号'"`
	Mark              string          `json:"mark" gorm:"comment:'客户备注'"`
	Note              string          `json:"note" gorm:"comment:'商家内部备注'"`
	Products          []OrdersProduct `gorm:"-"`
	Address           ClientAddress   `gorm:"-"`
}

func (OrdersType) TableName() string {
	return "orders"
}

type OrdersProduct struct {
	global.DATE_MODEL
	OrderID   string  `json:"order_id" gorm:"index;comment:订单号，关联到orders表"`
	UserID    string  `json:"user_id" gorm:"index;comment:用户ID"`
	ProductID string  `json:"product_id" gorm:"index;comment:商品ID，关联到product表"`
	Quantity  uint64  `json:"quantity" gorm:"comment:商品数量"`
	Price     float64 `json:"price" gorm:"comment:商品价格"`
	OldPrice  float64 `json:"old_price" gorm:"comment:商品价格"`
	PriceOff  uint8   `json:"price_off" gorm:"comment:商品折扣"`
	Stock     uint64  `json:"stock" gorm:"comment:商品库存"`
	Title     string  `json:"title" gorm:"comment:商品名称"`
	MainImg   string  `json:"main_img" gorm:"comment:商品图片"`
	SizeTitle string  `json:"size_title" gorm:"comment:商品尺寸标题"`
	Color     string  `json:"color" gorm:"comment:商品颜色"`
	Size      string  `json:"size" gorm:"comment:商品尺寸"`
	Shape     string  `json:"shape" gorm:"comment:商品甲型形状"`
	Mark      string  `json:"mark" gorm:"comment:商品备注"`
}

func (OrdersProduct) TableName() string {
	return "orders_product"
}
