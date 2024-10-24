package website

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	model "leiserv/models/website/yanwen"

	"github.com/go-resty/resty/v2"
	"github.com/golang-module/dongle"
)

type YanWenService struct {
	Url       string      `json:"url"`
	UserId    string      `json:"userId"`
	Format    string      `json:"format"`
	Method    string      `json:"method"`
	Timestamp int64       `json:"timestamp"`
	Version   string      `json:"version"`
	ApiToken  string      `json:"apiToken"`
	Data      interface{} `json:"data"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("Message:")
}

func (yanWen *YanWenService) Channel() model.YanWenChannlResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.channel.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body := []byte{}
	resp := PostReuqest(reqUrl, body)
	var channlResponse model.YanWenChannlResponse
	json.Unmarshal([]byte(resp), &channlResponse)
	return channlResponse
}

func (yanWen *YanWenService) Country() model.YanWenCountryResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.country.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body := []byte{}
	resp := PostReuqest(reqUrl, body)
	var countryResponse model.YanWenCountryResponse
	json.Unmarshal([]byte(resp), &countryResponse)
	return countryResponse
}

func (yanWen *YanWenService) Warehouse() model.YanWenWarehouseResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.warehouse.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var warehouseResponse model.YanWenWarehouseResponse
	json.Unmarshal([]byte(resp), &warehouseResponse)
	return warehouseResponse
}

func (yanWen *YanWenService) Order() model.YanWenOrderResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.order.create"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var orderResponse model.YanWenOrderResponse
	json.Unmarshal([]byte(resp), &orderResponse)
	return orderResponse
}

func (yanWen *YanWenService) Label() model.YanWenLabelResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.order.label.get"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var labelResponse model.YanWenLabelResponse
	json.Unmarshal([]byte(resp), &labelResponse)
	return labelResponse
}

func (yanWen *YanWenService) KrPccc() model.YanWenKrPcccResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.verify.kr.pccc"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var krResponse model.YanWenKrPcccResponse
	json.Unmarshal([]byte(resp), &krResponse)
	return krResponse
}

func (yanWen *YanWenService) USAddress() model.YanWenUSAddressResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.verify.us.address"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	log.Println(resp)
	var usAddressResponse model.YanWenUSAddressResponse
	json.Unmarshal([]byte(resp), &usAddressResponse)
	return usAddressResponse
}

func (yanWen *YanWenService) OrderInfo() model.YanWenOrderInfoResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.order.get"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var orderInfoResponse model.YanWenOrderInfoResponse
	json.Unmarshal([]byte(resp), &orderInfoResponse)
	return orderInfoResponse
}

func (yanWen *YanWenService) Cancel() model.YanWenCancelResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.order.cancel"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	log.Println(resp)
	var cancelResponse model.YanWenCancelResponse
	json.Unmarshal([]byte(resp), &cancelResponse)
	return cancelResponse
}

func (y *YanWenService) YangwenTracking(num string, auth string) model.YanWenTrackResponse {
	resp := GetTrackReuqest(num, auth)
	log.Println(resp)
	var trackResponse model.YanWenTrackResponse
	json.Unmarshal([]byte(resp), &trackResponse)
	return trackResponse
}

func (yanWen *YanWenService) Sign() string {
	var signStr string
	if yanWen.Data == "" {
		signStr = yanWen.ApiToken +
			yanWen.UserId +
			yanWen.Format +
			yanWen.Method +
			strconv.FormatInt(yanWen.Timestamp, 10) +
			yanWen.Version +
			yanWen.ApiToken
	} else {
		body, _ := json.Marshal(yanWen.Data)
		signStr = yanWen.ApiToken +
			yanWen.UserId +
			string(body) +
			yanWen.Format +
			yanWen.Method +
			strconv.FormatInt(yanWen.Timestamp, 10) +
			yanWen.Version +
			yanWen.ApiToken
	}
	return dongle.Encrypt.FromString(signStr).ByMd5().ToHexString()
}

func PostReuqest(url string, body []byte) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		//SetResult(). // or SetResult(AuthSuccess{}).
		Post(url)
	if err != nil {
		response := model.YanWenResponse{
			Success: false,
			Code:    "-1",
			Message: err.Error(),
		}
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}

// 查询运单的轨迹
func GetTrackReuqest(num string, auth string) string {
	const trackingURL = "http://api.track.yw56.com.cn/api/tracking"
	reqURL := fmt.Sprintf("%s?nums=%s", trackingURL, num)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", auth).
		Get(reqURL)
	if err != nil {
		response := model.YanWenTrackResponse{
			Code:    "-1",
			Message: err.Error(),
		}
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}
