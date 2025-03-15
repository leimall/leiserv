package website

import (
	"errors"
	"leiserv/global"
	websiteReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"leiserv/utils"
	"time"

	"gorm.io/gorm"
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

// get product lastest list
func (p *ProductService) GetLastestProductListDB(limit int) (list interface{}, err error) {
	var listObj []website.ProductListItme
	if limit == 0 {
		limit = 8
	}
	err = global.MALL_DB.Where("status = ?", 1).Order("updated_at desc").Limit(limit).Find(&listObj).Error
	if err != nil {
		return nil, err
	}
	return listObj, nil
}

// get product by product_id
func (p *ProductService) GetProductDetailByPidServ(id string) (product website.AllProduct, err error) {
	err = global.MALL_DB.Where("product_id =?", id).First(&product).Error
	return product, err
}

func (p *ProductService) GetProductListByCategoryDB(pId []string, limit int) (list []website.ProductListItme, err error) {
	var listObj []website.ProductListItme
	if limit == 0 {
		limit = 8
	}
	err = global.MALL_DB.Where("product_id in (?)", pId).Order("updated_at desc").Limit(limit).Find(&listObj).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return list, nil
		}
		return list, err
	}
	return listObj, nil
}
