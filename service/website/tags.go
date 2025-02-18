package website

import (
	"errors"
	"leiserv/global"
	websitetypes "leiserv/models/website/types"

	"gorm.io/gorm"
)

type TagsService struct{}

func (t *TagsService) CreateTagDB(tag websitetypes.Tag) (err error) {
	err = global.MALL_DB.Create(&tag).Error
	return err
}

func (t *TagsService) GetTagByIDDB(tagID uint) (tag websitetypes.Tag, err error) {
	err = global.MALL_DB.Where("id =?", tagID).First(&tag).Error
	return tag, err
}
func (t *TagsService) GetTagByTypeDB(tagID uint) (tag []websitetypes.Tag, err error) {
	err = global.MALL_DB.Where("type =?", tagID).Find(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tag, nil
		}
		return tag, err
	}
	return tag, err
}

func (t *TagsService) GetTagsListDB() (tags []websitetypes.Tag, err error) {
	err = global.MALL_DB.Order("type desc").Find(&tags).Error
	return tags, err
}

func (t *TagsService) UpdateTagDB(tag websitetypes.Tag) (err error) {
	err = global.MALL_DB.Save(&tag).Error
	return err
}

func (t *TagsService) DeleteTagDB(tag websitetypes.Tag) (err error) {
	return err
}

func (t *TagsService) GetTagListForProduct(pIDS []string) (tagMap map[string][]websitetypes.TagInfo, err error) {
	var tags []websitetypes.TagInfo
	err = global.MALL_DB.Where("product_id in (?)", pIDS).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	tagMap = make(map[string][]websitetypes.TagInfo)
	for _, tag := range tags {
		tagMap[tag.ProductID] = append(tagMap[tag.ProductID], tag)
	}
	return tagMap, err
}

// sku get sku all list
func (s *TagsService) GetSKUAllList(typeID uint) (list []websitetypes.Tag, err error) {
	err = global.MALL_DB.Where("type =?", typeID).Find(&list).Error
	return list, err
}

func (s *TagsService) GetProductBrandTagsByIDDB(pID uint) (tags []websitetypes.TagList, err error) {
	err = global.MALL_DB.Model(&websitetypes.Tag{}).Where("type = ? and parent_id in (?)", 4, pID).Find(&tags).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tags, err
	}
	if err != nil {
		return nil, err
	}
	return tags, err
}

func (s *TagsService) GetProductBrandChildrenTagsByIDDB(pID uint) (tags []websitetypes.Tag, err error) {
	err = global.MALL_DB.Model(&websitetypes.Tag{}).Where("type = ? and parent_id in (?)", 4, pID).Find(&tags).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tags, err
	}
	if err != nil {
		return nil, err
	}
	return tags, err
}
