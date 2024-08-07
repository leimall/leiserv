package website

import (
	"fmt"
	"leiserv/models/common/response"
	webReq "leiserv/models/website/request"
	webRes "leiserv/models/website/response"
	website "leiserv/models/website/types"
	"leiserv/utils"

	"github.com/gin-gonic/gin"
)

type AddressAPI struct{}

func (a *AddressAPI) GetAddress(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	fmt.Println("userID:", userId)
	lists, total, err := addressService.GetAddressLists(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  lists,
		Total: total,
	}, "OK", c)
}

func (a *AddressAPI) CreateAddress(c *gin.Context) {
	var reqadds webReq.Address
	if err := c.ShouldBindJSON(&reqadds); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetWebUserID(c)

	address := &website.ClientAddress{
		UserId:      userId,
		FirstName:   reqadds.FirstName,
		LastName:    reqadds.LastName,
		Street1:     reqadds.Street1,
		Street2:     reqadds.Street2,
		City:        reqadds.City,
		Region:      reqadds.Region,
		CountryCode: reqadds.CountryCode,
		Country:     reqadds.Country,
		ZipCode:     reqadds.ZipCode,
		Phone:       reqadds.Phone,
		IsDefault:   reqadds.IsDefault,
	}

	addsReturn, err := addressService.CreateAddressOne(*address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(webRes.Address{Address: addsReturn}, "OK", c)

}
