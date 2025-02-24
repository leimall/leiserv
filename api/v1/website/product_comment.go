package website

import (
	"fmt"
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"

	"github.com/gin-gonic/gin"
)

type ProductCommentAPI struct{}

func (a *ProductCommentAPI) GetCommentListByProductID(c *gin.Context) {
	var pageinfo webauthReq.CommentPage
	err := c.ShouldBindQuery(&pageinfo)
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, totle, err := productCommentService.GetProductReviewByProductID(pageinfo)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    totle,
		Page:     pageinfo.Page,
		PageSize: pageinfo.PageSize,
	}, "获取评论列表成功", c)

}
func (a *ProductCommentAPI) PostCommetByOrderID(c *gin.Context) {

	var comment webauthReq.CommentPost

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productCommentService.PostCommentByOrderID(comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("评论成功", c)

}
