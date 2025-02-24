package website

import (
	"leiserv/global"
	webauthReq "leiserv/models/website/request"
	webtype "leiserv/models/website/types"
)

type ProductCommentService struct{}

func (a *ProductCommentService) GetProductReviewByProductID(pageinfo webauthReq.CommentPage) (list interface{}, total int64, err error) {
	var comment []webtype.Comment
	offset := (pageinfo.Page - 1) * pageinfo.PageSize
	limit := 3
	db := global.MALL_DB.Model(&webtype.Comment{})
	err = db.Where("product_id =?", pageinfo.ProductID).Order("date desc").Count(&total).Offset(offset).Limit(limit).Find(&comment).Error
	return comment, total, err
}

func (a *ProductCommentService) PostCommentByOrderID(comment webauthReq.CommentPost) (err error) {
	err = global.MALL_DB.Create(comment).Error
	return err
}
