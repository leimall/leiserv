package website

import (
	"encoding/json"
	"fmt"
	"leiserv/models/common/response"
	models "leiserv/models/website/lianlianpay"
	websiteReq "leiserv/models/website/request"

	"github.com/gin-gonic/gin"
)

type LLPayAPI struct{}

var lianlianpay = models.LLPayConfigType{
	BaseUrl:    "https://celer-api.LianLianpay-inc.com/v3",
	LLPubKey:   "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1LUX1mFiLT7XgcAbwC8RxRl8S4o/ihxsTl8d6X1AYJxjeT9kq2I1gmGSJPVdRqQWZZmZ86e4EwZlwOAmuPaXT7ILSbepkSU2IJmv8+Pmx8lG0KybsEgezQ7la+LAllXvKtUp/AA1c3EEZGwjVoWEWvO9PUnyRVYgvJ1seM1AQDtQYvUHfm8a0CNTCYhqNiIivDtnxG4a8H7tcnnwPOipRCffi+S4CnQxXPaWoRbEISjZnt1KpfbkGAFXnbvJPAzDdD",
	PublickKey: "",
	PrivateKey: "",
	MerchantID: "202410250002787001",
}

func (p *LLPayAPI) GetToken(c *gin.Context) {
	url := lianlianpay.BaseUrl + fmt.Sprintf("/merchants/%s/token", lianlianpay.MerchantID)
	res, err := lianlianpayService.GetLLPayToken(url, lianlianpay.MerchantID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var trackResponse models.TokenResponse
	json.Unmarshal([]byte(res), &trackResponse)
	response.OkWithDetailed(trackResponse, "OK", c)
}

// create payment
func (p *LLPayAPI) CreatePayment(c *gin.Context) {
	var orders websiteReq.CreateLianLianPayType
	if err := c.ShouldBindJSON(&orders); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(orders)
	response.OkWithMessage("success", c)
}
