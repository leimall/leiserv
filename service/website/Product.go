package website

import (
	"leiserv/global"
	websiteReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"leiserv/utils"
	"time"
)

type ProductService struct{}

func (p *ProductService) CreateProductOne(product website.Produce) (err error) {
	product.ProductID = utils.SnowflakeID()
	err = global.MALL_DB.Create(&product).Error
	if err != nil {
		return err
	}
	return err
}

func (p *ProductService) GetProductList(info websiteReq.PageInfo) (list []website.AllProduct, total int64, err error) {
	var listObj []website.AllProduct
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	err = global.MALL_DB.Order("id desc").Limit(limit).Offset(offset).Find(&listObj).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return listObj, total, nil
}

func (p *ProductService) CreateProduct() error {
	return nil
}

func (p *ProductService) UpdateProductOne(product website.Produce) error {
	product.UpdatedAt = time.Now()
	return global.MALL_DB.Save(&product).Error
}

func (p *ProductService) DeleteProduct(id int) error {
	return nil
}

func (p *ProductService) SetMainImageUrl(pid string, url string) (err error) {
	err = global.MALL_DB.Model(&website.Produce{}).Where("product_id =?", pid).Update("main_img", url).Error
	return err
}
