package lianlianpay

type LLPayConfigType struct {
	BaseUrl       string `json:"base_url"`
	LLPubKey      string `json:"ll_pub_key"`
	PublickKey    string `json:"public_key"`
	PrivateKey    string `json:"private_key"`
	MerchantID    string `json:"meerchant_id"`
	SubMerchantId string `json:"sub_merchant_id"`
}
