package website

import (
	"errors"
	"leiserv/global"
	website "leiserv/models/website/types"

	"gorm.io/gorm"
)

type CartService struct{}

func (s *CartService) AddCartDB(cart website.CartItem) (err error) {
	err = global.MALL_DB.Where("user_id =? AND product_id =?", cart.UserId, cart.ProductID).Updates(cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	err = global.MALL_DB.Create(&cart).Error
	return err
}

func (s *CartService) GetCartDB(userId string) (cart []website.CartItem, err error) {
	err = global.MALL_DB.Where("user_id =? AND status = 1", userId).Find(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cart, nil
		}
		return cart, err
	}
	return cart, err
}

// UpdateCartDB updates the cart item in the database

func (s *CartService) UpdateCartDB(cart website.CartItem) (err error) {
	err = global.MALL_DB.Where("user_id =? AND product_id =?", cart.UserId, cart.ProductID).Updates(cart).Error
	return err
}

func (s *CartService) DeleteCartDB(userId string) (err error) {
	err = global.MALL_DB.Model(&website.CartItem{}).Where("user_id =?", userId).Update("status", 0).Error
	return err
}

func (s *CartService) DeleteCartOneDB(productId string) (err error) {
	err = global.MALL_DB.Model(&website.CartItem{}).Where("product_id =?", productId).Update("status", 0).Error
	return err
}
