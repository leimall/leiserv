package website

import (
	"errors"
	"fmt"
	"leiserv/global"
	websitetypes "leiserv/models/website/types"

	"gorm.io/gorm"
)

type SkuService struct{}

func (s *SkuService) GetSkuListForProduct(productID string) (skus []websitetypes.SkuInfo, err error) {
	// 1. 查询父 SKU（parent_id = 0 且 product_id 匹配）
	var parentSkus []websitetypes.SkuItem
	err = global.MALL_DB.Model(&websitetypes.SkuItem{}).
		Where("product_id = ? AND parent_id = 0", productID).
		Find(&parentSkus).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return skus, nil
		}
		return skus, fmt.Errorf("查询父 SKU 失败: %v", err)
	}

	// 2. 遍历父 SKU，查询对应的子 SKU
	for _, parent := range parentSkus {
		var childSkus []websitetypes.SkuItem
		err = global.MALL_DB.Model(&websitetypes.SkuItem{}).
			Where("product_id = ? AND parent_id = ?", productID, parent.ID).
			Find(&childSkus).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("查询子 SKU 失败: %v", err)
		}

		// 3. 组装 SkuInfo：父 SKU 信息 + 子 SKU 列表
		skuInfo := websitetypes.SkuInfo{
			SkuItem: parent,    // 父 SKU 自身信息
			List:    childSkus, // 子 SKU 列表
		}
		skus = append(skus, skuInfo)
	}

	return skus, nil
}
