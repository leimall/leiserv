package website

import (
	"errors"
	"leiserv/global"
	website "leiserv/models/website/types"

	"gorm.io/gorm"
)

type ProductBrandService struct{}

// get brand by product id
func (p *ProductBrandService) GetProductBrandByIDDB(productID string) (brand website.ProductBrand, err error) {
	err = global.MALL_DB.Where("product_id = ?", productID).First(&brand).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return brand, nil
	}
	return brand, err
}
