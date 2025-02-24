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

// update order detail of product api
func (p *OrdersService) UpdateOrdersProductDB(orderDetail website.OrdersType) (err error) {
	err = global.MALL_DB.Save(&orderDetail).Error
	return err
}

// get myself order list api
func (p *OrdersService) GetMyOrdersListDB(pageinfo webauthReq.PageInfo, user_id string) (list []website.OrdersType, total int64, err error) {
	offset := (pageinfo.Page - 1) * pageinfo.PageSize
	limit := pageinfo.PageSize
	db := global.MALL_DB.Model(&website.OrdersType{})
	err = db.Where("user_id =?", user_id).Order("updated_at desc").Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

// get one order by id api
func (p *OrdersService) GetOneOrderByIDDB(order_id string) (order website.OrdersType, err error) {
	err = global.MALL_DB.Model(&website.OrdersType{}).Where("order_id =?", order_id).First(&order).Error
	return order, err
}

// order_product DB
// get one order by id in the product api
func (p *OrdersService) GetOneOrderByIDProductsDB(order_id string) (order []website.OrdersProduct, err error) {
	err = global.MALL_DB.Model(&website.OrdersProduct{}).Where("order_id =?", order_id).Find(&order).Error
	return order, err
}

// create order detail of product api
func (p *OrdersService) CreateOrdersProductDB(orderDetail []website.OrdersProduct) (err error) {
	err = global.MALL_DB.Create(&orderDetail).Error
	return err
}

// get product detail by order_id []
func (p *OrdersService) GetMyOrdersProductDB(oIDs []string) (productMap map[string][]website.OrdersProduct, err error) {
	var product []website.OrdersProduct
	err = global.MALL_DB.Model(&website.OrdersProduct{}).Where("order_id IN (?)", oIDs).Find(&product).Error
	if err != nil {
		return nil, err
	}
	productMap = make(map[string][]website.OrdersProduct)
	for _, v := range product {
		productMap[v.OrderID] = append(productMap[v.OrderID], v)
	}
	return productMap, err
}
