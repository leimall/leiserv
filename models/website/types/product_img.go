package website

import "leiserv/global"

// product_id varchar(16)  NOT NULL COMMENT '商品ID',
//
//	imgurl VARCHAR(255) NOT NULL COMMENT '商品图片地址',
//	sort_id int UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片排序序号, 0表示主图, 1..n表示其他图片',
//	type tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '图片类型 1:商品展示图, 2:商品详情图, 3: 商品轮播图, 9: 其他图片',
//	status tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态 0:删除, 1: 上架',
//	name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '图片名称',
//	tag VARCHAR(32) NOT NULL DEFAULT '' COMMENT '图片标签',
//	uuid VARCHAR(64) NOT NULL DEFAULT '' COMMENT '图片UUID',
type ProductImg struct {
	global.DATE_MODEL
	ProductID string `json:"product_id" gorm:"index;comment:商品ID"`
	ImgURL    string `json:"img_url" gorm:"comment:商品图片地址"`
	SortID    uint   `json:"sort_id" gorm:"default:0;comment:图片排序序号, 0表示主图, 1..n表示其他图片"`
	Type      uint   `json:"type" gorm:"default:1;comment:图片类型 1:商品展示图, 2:商品详情图, 3: 商品轮播图, 9: 其他图片"`
	Status    uint   `json:"status" gorm:"default:1;comment:商品状态 0:删除, 1: 上架"`
	Name      string `json:"name" gorm:"default:'';comment:图片名称"`
	Tag       string `json:"tag" gorm:"default:'';comment:图片标签"`
	UUID      string `json:"uuid" gorm:"default:'';comment:图片UUID"`
	IsMain    int    `json:"is_main" gorm:"default:0;comment:是否主图"`
}

func (ProductImg) TableName() string {
	return "product_img"
}
