package website

import (
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"

	"github.com/gin-gonic/gin"
)

type CommonAPI struct{}

func (st *CommonAPI) GetCountries(c *gin.Context) {
	list, total, err := commonService.GetCountriesLists()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (st *CommonAPI) GetRegion(c *gin.Context) {
	var cid webauthReq.CommonID
	err := c.ShouldBindJSON(&cid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := commonService.GetRegionLists(cid.CommonID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (st *CommonAPI) GetCity(c *gin.Context) {
	var rid webauthReq.CommonID
	err := c.ShouldBindJSON(&rid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := commonService.GetCityLists(rid.CommonID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)

}
