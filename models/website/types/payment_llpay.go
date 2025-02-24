package website

import "leiserv/global"

// id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键，自增',
// order_id VARCHAR(64) NOT NULL COMMENT '订单号，关联到orders表',
// user_id VARCHAR(255) NOT NULL COMMENT '用户 ID',
// ll_transaction_id VARCHAR(255) NOT NULL COMMENT '第三方支付系统生成的交易 ID',
// merchant_transaction_id VARCHAR(255) NOT NULL COMMENT '商户生成的交易 ID',
// payment_currency_code VARCHAR(10) NOT NULL COMMENT '支付币种，例如 "USD"',
// payment_amount DECIMAL(10, 2) NOT NULL COMMENT '支付金额，保留两位小数',
// exchange_rate DECIMAL(10, 8) NOT NULL COMMENT '汇率，精度为 8 位',
// payment_status VARCHAR(10) NOT NULL COMMENT '支付状态（例如 "WP" - 待支付，"PS" - 支付成功）',
// settlement_currency_code VARCHAR(10) NOT NULL COMMENT '结算币种，通常与支付币种相同',
// settlement_amount DECIMAL(10, 2) NOT NULL COMMENT '结算金额，保留两位小数',
// installments INT NOT NULL COMMENT '分期付款数量，默认为 1，表示一次性支付',
// three_d_secure_status VARCHAR(50) NOT NULL COMMENT '3D Secure 状态（例如 "CHALLENGE" - 挑战，"PASSED" - 通过验证）',
// payment_url TEXT NOT NULL COMMENT '支付 URL，用于支付过程中的重定向链接',
// trace_id VARCHAR(255) NOT NULL COMMENT '追踪 ID，用于唯一标识支付流程，便于追踪',
// key_value VARCHAR(255) NOT NULL COMMENT '支付相关的 key，通常用于防止重复支付等用途',
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认为当前时间',
// payment_time TIMESTAMP NOT NULL COMMENT '支付时间，记录支付成功的时间',
// updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',

type PaymentLlPay struct {
	global.MODEL
	OrderId                string `json:"order_id" gorm:"column:order_id;type:varchar(64);not null;comment:订单号，关联到orders表"`
	UserId                 string `json:"user_id" gorm:"column:user_id;type:varchar(255);not null;comment:用户 ID"`
	LlTransactionId        string `json:"ll_transaction_id" gorm:"column:ll_transaction_id;type:varchar(255);not null;comment:第三方支付系统生成的交易 ID"`
	MerchantTransactionId  string `json:"merchant_transaction_id" gorm:"column:merchant_transaction_id;type:varchar(255);not null;comment:商户生成的交易 ID"`
	PaymentCurrencyCode    string `json:"payment_currency_code" gorm:"column:payment_currency_code;type:varchar(10);not null;comment:支付币种，例如 \"USD\""`
	PaymentAmount          string `json:"payment_amount" gorm:"column:payment_amount;type:decimal(10,2);not null;comment:支付金额，保留两位小数"`
	ExchangeRate           string `json:"exchange_rate" gorm:"column:exchange_rate;type:decimal(10,8);not null;comment:汇率，精度为 8 位"`
	PaymentTime            string `json:"payment_time" gorm:"column:payment_time;type:varchar(255);not null;comment:支付时间，记录支付成功的时间"`
	PaymentStatus          string `json:"payment_status" gorm:"column:payment_status;type:varchar(10);not null;comment:支付状态（例如 \"WP\" - 待支付，\"PS\" - 支付成功）"`
	SettlementCurrencyCode string `json:"settlement_currency_code" gorm:"column:settlement_currency_code;type:varchar(10);not null;comment:结算币种，通常与支付币种相同"`
	SettlementAmount       string `json:"settlement_amount" gorm:"column:settlement_amount;type:decimal(10,2);not null;comment:结算金额，保留两位小数"`
	Installments           string `json:"installments" gorm:"column:installments;type:int;not null;comment:分期付款数量，默认为 1，表示一次性支付"`
	PaymentUrl             string `json:"payment_url" gorm:"column:payment_url;type:text;not null;comment:支付 URL，用于支付过程中的重定向链接"`
	TraceId                string `json:"trace_id" gorm:"column:trace_id;type:varchar(255);not null;comment:追踪 ID，用于唯一标识支付流程，便于追踪"`
	KeyValue               string `json:"key_value" gorm:"column:key_value;type:varchar(255);not null;comment:支付相关的 key，通常用于防止重复支付等用途"`
}

func (PaymentLlPay) TableName() string {
	return "payment_llpay"
}
