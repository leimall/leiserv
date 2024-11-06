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
