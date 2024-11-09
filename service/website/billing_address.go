package website

import (
	"errors"
	"leiserv/global"
	website "leiserv/models/website/types"

	"gorm.io/gorm"
)

type BillingAddressService struct{}

// create a new billing address
func (s *BillingAddressService) CreateBillingAddress(baddress website.BillingAddress) error {
	err := global.MALL_DB.Create(&baddress).Error
	return err
}

func (s *BillingAddressService) UpdateBillingAddress(baddress website.BillingAddress) error {
	err := global.MALL_DB.Save(&baddress).Error
	return err
}

func (s *BillingAddressService) GetBillingAddressByUserID(user_id string) (website.BillingAddress, error) {
	var baddress website.BillingAddress
	err := global.MALL_DB.First(&baddress, user_id).Error
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
