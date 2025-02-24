package website

import (
	"fmt"
	"leiserv/global"
	website "leiserv/models/website/types"
)

type CategoryService struct{}

func (s *CategoryService) GetCategoriesByProductID(pid string) (list interface{}, totel int64, err error) {
	fmt.Println("GetCategoriesByProductID", pid)
	var listObj []website.CategoryInfo
	err = global.MALL_DB.Where("product_id =?", pid).Find(&listObj).Count(&totel).Error
	if err != nil {
		return nil, 0, err
	}
	return listObj, totel, err
}

func (s *CategoryService) CreateCategory(c website.CategoryInfo) (err error) {
	err = global.MALL_DB.Create(&c).Error
	return err
}

func (s *CategoryService) UpdateCategory(id int) {}

func (s *CategoryService) DeleteCategory(id int) {}

// 商品列表时, 获取商品对应的分类
func (s *CategoryService) GetCatagoryListForProduct(pIDs []string) (cateMap map[string][]website.CategoryInfo, err error) {
	var category []website.CategoryInfo
	err = global.MALL_DB.Where("product_id IN (?)", pIDs).Find(&category).Error
	if err != nil {
		return nil, err
	}
	cateMap = make(map[string][]website.CategoryInfo)
	for _, tag := range category {
		cateMap[tag.ProductID] = append(cateMap[tag.ProductID], tag)
	}
	return cateMap, err
}

// get category list by title
// return product_id lists
func (s *CategoryService) GetCategoryByTitleDB(title string) (list []website.CategoryInfo, err error) {
	err = global.MALL_DB.Where("title LIKE ?", "%"+title+"%").Find(&list).Error
	if err != nil {
		return nil, err
	}
	// cateMap = make(map[string][]website.CategoryInfo)
	// for _, tag := range list {
	// 	cateMap[tag.ProductID] = append(cateMap[tag.ProductID], tag)
	// }
	return list, err
}
