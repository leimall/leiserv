package website

import (
	"leiserv/models/common/response"
	website "leiserv/models/website/types"

	"github.com/gin-gonic/gin"
)

type ReviewAPI struct{}

// create a new review
func (r *ReviewAPI) CreateProductReview(c *gin.Context) {
	var review website.ProductReviewsItem
	err := c.ShouldBindJSON(&review)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productReviewService.CreateProductReviewDB(review)
	if err != nil {
		response.FailWithMessage("创建评价失败", c)
		return
	}
	response.OkWithMessage("创建评价成功", c)
}

// get all reviews for a product_id
func (r *ReviewAPI) GetProductReview(c *gin.Context) {
	var review website.ProductReviewsItem
	err := c.ShouldBindJSON(&review)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reviews, err := productReviewService.GetProductReviewByIDDB(review.ProductID)
	if err != nil {
		response.FailWithMessage("获取评价失败", c)
		return
	}
	response.OkWithData(reviews, c)
}

// update a review item
func (r *ReviewAPI) UpdateProductReview(c *gin.Context) {
	var review website.ProductReviewsItem
	err := c.ShouldBindJSON(&review)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productReviewService.UpdateProductReviewDB(review)
	if err != nil {
		response.FailWithMessage("更新评价失败", c)
		return
	}
	response.OkWithMessage("更新评价成功", c)
}

// delete a review item
func (r *ReviewAPI) DeleteProductReview(c *gin.Context) {
	var review website.ProductReviewsItem
	err := c.ShouldBindJSON(&review)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productReviewService.DeleteProductReviewDB(review.ID)
	if err != nil {
		response.FailWithMessage("删除评价失败", c)
		return
	}
	response.OkWithMessage("删除评价成功", c)
}
