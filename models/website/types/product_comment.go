package website

import "leiserv/global"

type Comment struct {
	global.DATE_MODEL
	ProductID   string `json:"productId" gorm:"index;comment:产品ID"` // 产品ID
	UserID      string `json:"userId" gorm:"index;comment:用户ID"`    // 用户ID
	OrderID     string `json:"orderId" gorm:"comment:订单ID"`         // 订单ID
	UserName    string `json:"userName" gorm:"comment:用户名"`         // 用户名
	Title       string `json:"title" gorm:"comment:标题"`             // 标题
	Content     string `json:"content" gorm:"comment:评论内容"`         // 评论内容
	Star        int    `json:"star" gorm:"comment:评分"`              // 评分
	IsImg       bool   `json:"isImg" gorm:"comment:是否有图片"`          // 是否有图片
	ImgUrl      string `json:"imgUrl" gorm:"comment:图片地址"`          // 图片地址
	ShopContent string `json:"shopContent" gorm:"comment:商家回复内容"`   // 商家回复内容
	Date        string `json:"date" gorm:"comment:日期"`              // 日期
}

func (Comment) TableName() string {
	return "product_comment"
}
