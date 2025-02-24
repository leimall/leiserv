package website

// LianLianPay defines the main payment structure, including transaction and merchant details,
// optional URLs, and additional payment-related information.
type LianLianPay struct {
	MerchantTransactionId string           `json:"merchant_transaction_id"` // 支付交易ID
	MerchantId            string           `json:"merchant_id"`             // 商户ID
	SubMerchantId         string           `json:"sub_merchant_id"`         // 站点号, 本地钱包支付方式必填
	NotificationUrl       string           `json:"notification_url"`        // 支付结果通知URL, API国际信用卡必填
	RedirectUrl           string           `json:"redirect_url"`            // 支付成功后的跳转URL, API国际信用卡必填
	Country               string           `json:"country"`                 // 国家代码
	PaymentMethod         string           `json:"payment_method"`          // 支付方式, Direct API模式/IFrame模式必填
	MerchantOrder         MerchantOrder    `json:"merchant_order"`          // 商户订单信息
	Customer              LLPayCustomer    `json:"customer"`                // 客户详情
	PaymentData           LLPayPaymentData `json:"payment_data"`            // 支付数据
}

// MerchantOrder defines the merchant's order structure, containing order metadata and currency information.
type MerchantOrder struct {
	MerchantOrderId   string         `json:"merchant_order_id"`   // 商户订单ID
	MerchantOrderTime string         `json:"merchant_order_time"` // 商户订单时间 (yyyyMMddHHmmss)
	OrderAmount       float64        `json:"order_amount"`        // 订单金额
	OrderCurrencyCode string         `json:"order_currency_code"` // 订单币种代码
	Products          []LLPayProduct `json:"products"`            // 订单中的商品列表
	Shipping          LLPayShipping  `json:"shipping"`            // 订单中的物流信息
}

// LLPayProduct defines the structure for each product in the merchant order.
type LLPayProduct struct {
	ProductId        string  `json:"product_id"`        // 商品ID
	Name             string  `json:"name"`              // 商品名称
	Description      string  `json:"description"`       // 商品描述
	Price            float64 `json:"price"`             // 商品单价
	Quantity         uint64  `json:"quantity"`          // 商品数量
	Category         string  `json:"category"`          // 商品分类
	Sku              string  `json:"sku"`               // 商品SKU
	Url              string  `json:"url"`               // 商品网址
	ShippingProvider string  `json:"shipping_provider"` // 物流供应商
}

type LLPayShipping struct {
	Name    string              `json:"name"`    // 物流名称
	Cycle   string              `json:"cycle"`   // 物流周期
	Address LLPayBillingAddress `json:"address"` // 账单地址
}

// LLPayPaymentData includes installment and card information for payment data.
type LLPayPaymentData struct {
	Installments uint8     `json:"installments"` // 分期付款期数
	Card         LLPayCard `json:"card"`         // 信用卡信息
}

// LLPayCard defines the cardholder details and billing address for the transaction.
type LLPayCard struct {
	HolderName     string              `json:"holder_name"` // 持卡人姓名
	CardToken      string              `json:"card_token"`  // 卡的Token
	BillingAddress LLPayBillingAddress `gorm:"-"`           // 账单地址
}

// LLPayBillingAddress defines the billing address details.
type LLPayBillingAddress struct {
	Line1      string `json:"line1"`       // 地址行1
	Line2      string `json:"line2"`       // 地址行2
	City       string `json:"city"`        // 城市
	State      string `json:"state"`       // 省/自治区/直辖市
	Country    string `json:"country"`     // 国家
	PostalCode string `json:"postal_code"` // 邮政编码
	District   string `json:"district"`    // 区/县
}

// LLPayCustomer defines the structure for customer information.
type LLPayCustomer struct {
	CustomerType string              `json:"customer_type"` // 客户类型 (如公司C)
	FullName     string              `json:"full_name"`     // 客户全名或公司名称
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	Address      LLPayBillingAddress `json:"address"` // 账单地址
}

type LLPayConfigType struct {
	BaseUrl    string `json:"base_url"`
	LLPubKey   string `json:"ll_pub_key"`
	PublickKey string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	MerchantID string `json:"meerchant_id"`
}
