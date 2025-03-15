package website

import (
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
	var reqadds webReq.ClientAddress
	if err := c.ShouldBindJSON(&reqadds); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetWebUserID(c)

	address := &website.ClientAddress{
		UserId:      userId,
		FirstName:   reqadds.FirstName,
		LastName:    reqadds.LastName,
		Line1:       reqadds.Line1,
		Email:       reqadds.Email,
		City:        reqadds.City,
		State:       reqadds.State,
		CountryName: reqadds.CountryName,
		Country:     reqadds.Country,
		PostalCode:  reqadds.PostalCode,
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

func (a *AddressAPI) UpdateAddress(c *gin.Context) {
	var reqadds website.ClientAddress
	if err := c.ShouldBindJSON(&reqadds); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetWebUserID(c)
	reqadds.UserId = userId
	err := addressService.UpdateAddressOne(reqadds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}

func (a *AddressAPI) DeleteAddress(c *gin.Context) {
	var addressId webReq.AddressID
	userId := utils.GetWebUserID(c)
	if err := c.ShouldBindJSON(&addressId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := addressService.DeleteAddressOne(userId, addressId.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}

// set default address
func (a *AddressAPI) PutDefaultAddress(c *gin.Context) {
	var addressId webReq.AddressID
	userId := utils.GetWebUserID(c)
	if err := c.ShouldBindJSON(&addressId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := addressService.SetDefaultAddress(userId, addressId.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}

// billing address
func (a *AddressAPI) GetBillingAddress(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	adds, err := billingAddressService.GetBillingAddressByUserID(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(adds, "OK", c)
}

func (a *AddressAPI) CreateBillingAddress(c *gin.Context) {
	var address website.BillingAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetWebUserID(c)
	address.UserId = userId

	err := billingAddressService.CreateBillingAddress(address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(address, "OK", c)
}
func (a *AddressAPI) UpdateBillingAddress(c *gin.Context) {
	var reqadds website.BillingAddress
	if err := c.ShouldBindJSON(&reqadds); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := billingAddressService.UpdateBillingAddress(reqadds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}

func (a *AddressAPI) DeleteBillingAddress(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	err := billingAddressService.DeleteBillingAddress(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OK", c)
}
