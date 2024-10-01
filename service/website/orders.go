package website

import (
	"leiserv/global"
	webauthReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
)

type OrdersService struct{}

func (p *OrdersService) CreateOrdersDB(orders website.OrdersType) (err error) {
	err = global.MALL_DB.Create(&orders).Error
	return err
}

func (p *OrdersService) GetOrdersListDB(pageinfo webauthReq.PageInfo) (list interface{}, total int64, err error) {
	offset := (pageinfo.Page - 1) * pageinfo.PageSize
	limit := pageinfo.PageSize
	db := global.MALL_DB.Model(&website.OrdersType{})
	var orders []website.OrdersType
	err = db.Count(&total).Offset(offset).Limit(limit).Find(&orders).Error
	list = orders
	return list, total, err
}

func (p *OrdersService) GetOrdersList(pageinfo webauthReq.PageInfo) (list interface{}, total int64, err error) {
	offset := (pageinfo.Page - 1) * pageinfo.PageSize
	limit := pageinfo.PageSize
	db := global.MALL_DB.Model(&website.OrdersType{})
	var orders []website.OrdersType
	err = db.Count(&total).Offset(offset).Limit(limit).Find(&orders).Error
	list = orders
	return list, total, err
}

// create order detail of product api
func (p *OrdersService) CreateOrdersProductDB(orderDetail []website.OrdersProduct) (err error) {
	err = global.MALL_DB.Create(&orderDetail).Error
	return err
}
