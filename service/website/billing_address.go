package website

import (
	"errors"
	"leiserv/global"
	website "leiserv/models/website/types"
	"time"

	"gorm.io/gorm"
)

type BillingAddressService struct{}

// create a new billing address
func (s *BillingAddressService) CreateBillingAddress(address website.BillingAddress) error {
	var existingAddress website.BillingAddress

	if address.CreatedAt.IsZero() {
		address.CreatedAt = time.Now()
	}
	err := global.MALL_DB.Where("user_id = ?", address.UserId).First(&existingAddress).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			address.ID = 0
			return global.MALL_DB.Create(&address).Error
		}
		return err
	}
	address.ID = existingAddress.ID
	return global.MALL_DB.Save(&address).Error
}

func (s *BillingAddressService) UpdateBillingAddress(baddress website.BillingAddress) error {
	err := global.MALL_DB.Save(&baddress).Error
	return err
}

func (s *BillingAddressService) GetBillingAddressByUserID(user_id string) (website.BillingAddress, error) {
	var baddress website.BillingAddress
	err := global.MALL_DB.Where("user_id = ?", user_id).First(&baddress).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baddress, nil
		}
		return baddress, err
	}
	return baddress, err
}

func (s *BillingAddressService) DeleteBillingAddress(user_id string) error {
	err := global.MALL_DB.Delete(&website.BillingAddress{}, user_id).Error
	return err
}
