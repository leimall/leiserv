package website

import (
	"fmt"
	"leiserv/models/common/response"
	YanWenModel "leiserv/models/website/yanwen"
	"leiserv/service/website"

	"github.com/gin-gonic/gin"
)

type ShippingAPI struct{}

var yanwen = &website.YanWenService{
	Url:       "https://ejf-fat.yw56.com.cn",
	UserId:    "100000",
	Format:    "json",
	Method:    "",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
	Data:      "",
}

// 查询通达国家列表
func (p *ShippingAPI) GetCountryLists(c *gin.Context) {
	res := yanwen.Country()
	if !res.Success {
		response.FailWithMessage("Error", c)
		return
	}
	data := res.Data
	response.OkWithDetailed(response.ListsResult{List: data}, "OK", c)
}

// get order information
func (p *ShippingAPI) GetOrderInfo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id:", id)
	var orderInfoRequest = YanWenModel.YanWenOrderInfoRequest{}
	orderInfoRequest.WaybillNumber = "LR039896495CN"
	yanwen.Data = orderInfoRequest
	res := yanwen.OrderInfo()
	data := res.Data
	response.OkWithDetailed(data, "OK", c)
}

// print shipping

func (p *ShippingAPI) GetTrackingInfo(c *gin.Context) {
	id := c.Param("id")
	res := yanwenService.YangwenTracking(id, "20202218")
	data := res.Result
	response.OkWithDetailed(data, "OK", c)
}
