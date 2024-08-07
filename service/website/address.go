package website

import (
	"leiserv/global"
	website "leiserv/models/website/types"
)

type AddressService struct{}

func (st *AddressService) GetAddressLists(userId string) (list interface{}, total int64, err error) {
	var addressList []website.ClientAddress
	err = global.MALL_DB.Where("user_id=?", userId).Find(&addressList).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	return addressList, total, err
}

func (st *AddressService) CreateAddressOne(address website.ClientAddress) (res website.ClientAddress, err error) {

	err = global.MALL_DB.Create(&address).Error

	return address, err
}
