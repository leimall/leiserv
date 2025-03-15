package website

import (
	"leiserv/global"
	website "leiserv/models/website/types"
	"time"
)

type AddressService struct{}

func (st *AddressService) GetAddressLists(userId string) (list interface{}, total int64, err error) {
	var addressList []website.ClientAddress
	err = global.MALL_DB.Where("user_id=?", userId).Order("is_default DESC").Find(&addressList).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	return addressList, total, err
}

func (st *AddressService) CreateAddressOne(address website.ClientAddress) (res website.ClientAddress, err error) {
	if address.IsDefault == 1 {
		err := global.MALL_DB.Model(&website.ClientAddress{}).Where("user_id=?", address.UserId).Update("is_default", 0).Error
		if err != nil {
			return address, err
		}
	}
	err = global.MALL_DB.Create(&address).Error

	return address, err
}

func (st *AddressService) DeleteAddressOne(userid string, addressId int64) (err error) {
	err = global.MALL_DB.Where("user_id=? AND id=?", userid, addressId).Delete(&website.ClientAddress{}).Error
	return err
}
func (st *AddressService) UpdateAddressOne(address website.ClientAddress) (err error) {
	if address.CreatedAt.IsZero() {
		address.CreatedAt = time.Now()
	}
	if address.IsDefault == 1 {
		err := global.MALL_DB.Model(&website.ClientAddress{}).Where("user_id=?", address.UserId).Update("is_default", 0).Error
		if err != nil {
			return err
		}
	}
	err = global.MALL_DB.Save(&address).Error
	return err
}

// set defaule address
func (st *AddressService) SetDefaultAddress(userId string, addressId int64) (err error) {
	err = global.MALL_DB.Model(&website.ClientAddress{}).Where("user_id=?", userId).Update("is_default", 0).Error
	if err != nil {
		return err
	}
	err = global.MALL_DB.Model(&website.ClientAddress{}).Where("id=?", addressId).Update("is_default", 1).Error
	return err
}

// get address by id
func (st *AddressService) GetAddressById(userId string, addressId uint64) (address website.ClientAddress, err error) {
	err = global.MALL_DB.Where("user_id=? AND id=?", userId, addressId).First(&address).Error
	return address, err
}

// get address by id for order
func (st *AddressService) GetMyOrdersAddressDB(aIDs []uint64) (addressMap map[uint]website.ClientAddress, err error) {
	var addressList []website.ClientAddress
	err = global.MALL_DB.Where("id in (?)", aIDs).Find(&addressList).Error
	if err != nil {
		return addressMap, err
	}
	addressMap = make(map[uint]website.ClientAddress)
	for _, v := range addressList {
		addressMap[v.ID] = v
	}
	return addressMap, nil
}
