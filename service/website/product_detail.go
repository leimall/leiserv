package website

import (
	"leiserv/global"
	website "leiserv/models/website/types"
)

type ProductDetailService struct{}

// product detail create
func (p *ProductDetailService) CreateAndSaveProductDetailDB(productDetail []website.ProductDetail) (err error) {
	err = global.MALL_DB.Save(&productDetail).Error
	if err != nil {
		return err
	}
	return err
}

// get product detail by product_id
func (p *ProductDetailService) GetProductDetailByPidDB(pid string) (productDetail []website.ProductDetail, err error) {
	err = global.MALL_DB.Where("product_id =?", pid).Find(&productDetail).Error
	return productDetail, err
}

// product detail delete
func (p *ProductDetailService) DeleteProductDetailDB(pid uint) (err error) {
	err = global.MALL_DB.Where("product_id =?", pid).Error
	return err
}
