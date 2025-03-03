package website

import (
	"errors"
	"leiserv/global"
	websitetypes "leiserv/models/website/types"

	"gorm.io/gorm"
)

type SkuService struct{}

func (s *SkuService) CreateSkuTitleDB(sku websitetypes.SkuItem) (skuItem websitetypes.SkuItem, err error) {
	err = global.MALL_DB.Create(&sku).Error

	return sku, err
}
func (s *SkuService) UpdateSkuTitleDB(sku websitetypes.SkuItem) (err error) {
	err = global.MALL_DB.Save(&sku).Error
	return err
}

func (s *SkuService) CreateSkuValueDB(sku []websitetypes.SkuItem) (err error) {
	err = global.MALL_DB.Create(&sku).Error
	return err
}
func (s *SkuService) UpdateSkuValueDB(sku []websitetypes.SkuItem) (err error) {
	err = global.MALL_DB.Save(&sku).Error
	return err
}

func (s *SkuService) GetSkuListForProduct(productID string) (sku websitetypes.SkuInfo, err error) {
	var skuList []websitetypes.Tag

	err = global.MALL_DB.Where("product_id =? and  parent_id = 0", productID).First(&sku).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		sku.List = []websitetypes.Tag{}
		return sku, nil
	}
	if err != nil {
		return sku, err
	}

	err = global.MALL_DB.Where("type=? and parent_id =?", 2, sku.TagID).Find(&skuList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || len(skuList) == 0 {
		sku.List = skuList
		return sku, nil
	}
	if err != nil {
		return sku, err
	}

	sku.List = skuList
	return sku, nil
}
