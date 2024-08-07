package website

import (
	"leiserv/global"
	website "leiserv/models/website/types"
)

type CommonService struct{}

func (st *CommonService) GetCountriesLists() (list interface{}, total int64, err error) {
	var clist []website.DataRegion

	err = global.MALL_DB.Where("level=?", 2).Find(&clist).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return clist, total, nil
}

func (st *CommonService) GetRegionLists(countryId int) (list interface{}, total int64, err error) {
	var clist []website.DataRegion

	err = global.MALL_DB.Where("level=? AND pid=?", 3, countryId).Find(&clist).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}
	return clist, total, nil
}

func (st *CommonService) GetCityLists(regionId int) (list interface{}, total int64, err error) {
	var clist []website.DataRegion

	err = global.MALL_DB.Where("level=? AND pid=?", 4, regionId).Find(&clist).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}
	return clist, total, nil
}
