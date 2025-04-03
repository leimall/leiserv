package lianlianpay

type TokenRequest struct {
	MerchantID string `json:"merchant_id"`
	Timestamp  string `json:"timestamp"`
}

type PaymentInquiryRequest struct {
	MerchantID            string `json:"merchant_id"`
	MerchantTransactionID string `json:"merchant_transaction_id"`
}

type PaymentInquiryResponse struct {
	Response
	Order PaymentOrderOrder `grom:"-"`
}

type TokenResponse struct {
	ReturnCode    string `json:"return_code"`
	ReturnMessage string `json:"return_message"`
	TraceId       string `json:"trace_id"`
	Order         string `json:"order"`
}
type Response struct {
	ReturnCode    string `json:"return_code"`
	ReturnMessage string `json:"return_message"`
	TraceId       string `json:"trace_id"`
}

type PaymentResponseData struct {
	Response
	Order PaymentOrderOrder `grom:"-"`
}

type PaymentOrderOrder struct {
	LlTransactionId       string                       `json:"ll_transaction_id"`
	MerchantTransactionId string                       `json:"merchant_transaction_id"`
	PaymentData           PaymentOrderOrderPaymentData `json:"payment_data"`
	ThreeDsStatus         string                       `json:"3ds_status"`
	PaymentUrl            string                       `json:"payment_url"`
	Key                   string                       `json:"key"`
}

type PaymentOrderOrderPaymentData struct {
	PaymentCurrencyCode    string `json:"payment_currency_code"`
	PaymentAmount          string `json:"payment_amount"`
	ExchangeRate           string `json:"exchange_rate"`
	PaymentTime            string `json:"payment_time"`
	PaymentStatus          string `json:"payment_status"`
	SettlementCurrencyCode string `json:"settlement_currency_code"`
	SettlementAmount       string `json:"settlement_amount"`
	Installments           string `json:"installments"`
}
