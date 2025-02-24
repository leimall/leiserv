package lianlianpay

type TokenRequest struct {
	MerchantID string `json:"merchant_id"`
	Timestamp  string `json:"timestamp"`
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

// "order": {
//       "ll_transaction_id": "2024110703169182",
//       "merchant_transaction_id": "667356745156268032",
//       "payment_data": {
//          "payment_currency_code": "USD",
//          "payment_amount": "200.00",
//          "exchange_rate": "1.00000000",
//          "payment_status": "WP",
//          "settlement_currency_code": "USD",
//          "settlement_amount": "200.00",
//          "installments": "1"
//       },
//       "3ds_status": "CHALLENGE",
//       "payment_url": "https://gacashier.lianlianpay-inc.com/3ds?tdsKey=583c68e9730344e78614ec5eeb6dd43c",
//       "key": "1pgtbvn5jku"
//    }

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
