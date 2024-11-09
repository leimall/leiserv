package website

import (
	"encoding/json"
	models "leiserv/models/website/lianlianpay"
	website "leiserv/models/website/types"
	"log"
	"time"

	tools "leiserv/utils/lianlianpay"
)

type LLPayService struct{}

func (l *LLPayService) GetLLPayToken(url string, merchant_id string) (res []byte, err error) {
	timestamp := time.Now().Format("20060102150405")
	data := &models.TokenRequest{
		MerchantID: merchant_id,
		Timestamp:  timestamp,
	}
	signatureString := tools.ConvertStructToSignatureString(data)
	sign, err := tools.Sign(signatureString)
	if err != nil {
		log.Println("Error signing data:", err)
		return nil, err
	}
	body, _ := json.Marshal(data)
	res, err = tools.GetRequest(url, sign, timestamp, body)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (l *LLPayService) CreatePayment(url string, data website.LianLianPay, timestamp string) (res string, err error) {
	signatureString := tools.ConvertStructToSignatureString(data)
	sign, err := tools.Sign(signatureString)
	if err != nil {
		log.Println("Error signing data:", err)
		return "", err
	}
	body, _ := json.Marshal(data)
	res, err = tools.PostRequest(url, sign, timestamp, body)
	if err != nil {
		return "", err
	}
	return res, err
}
