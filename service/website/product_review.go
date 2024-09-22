package website

import (
	"errors"
	"leiserv/global"
	websitetypes "leiserv/models/website/types"

	"gorm.io/gorm"
)

type ProductReviewService struct{}

// create reivew item
func (p *ProductReviewService) CreateProductReviewDB(productReview websitetypes.ProductReviewsItem) (err error) {
	err = global.MALL_DB.Create(&productReview).Error
	return err
}

// get review item by id
func (p *ProductReviewService) GetProductReviewByIDDB(id string) (r websitetypes.ProductReviewsItem, err error) {
	var productReview websitetypes.ProductReviewsItem

	err = global.MALL_DB.Where("product_id =?", id).First(&productReview).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return productReview, nil
		}
		return productReview, err
	}

	return productReview, err
}

// update review item
func (p *ProductReviewService) UpdateProductReviewDB(productReview websitetypes.ProductReviewsItem) (err error) {
	err = global.MALL_DB.Save(&productReview).Error
	return err
}

// delete review item
func (p *ProductReviewService) DeleteProductReviewDB(id uint) (err error) {
	err = global.MALL_DB.Delete(&websitetypes.ProductReviewsItem{}, id).Error
	return err
}

// get review from list product id
func (p *ProductReviewService) GetProductReviewByProductIDDB(productID []string) (newReviews map[string]websitetypes.ProductReviewsItem, err error) {
	var reviews []websitetypes.ProductReviewsItem

	err = global.MALL_DB.Where("product_id IN ?", productID).Find(&reviews).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	reviewMaps := make(map[string]websitetypes.ProductReviewsItem)
	for _, item := range reviews {
		reviewMaps[item.ProductID] = item
	}
	return reviewMaps, nil
}
