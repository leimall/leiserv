package website

import (
	"fmt"
	"leiserv/models/common/response"
	"leiserv/models/website/request"
	website "leiserv/models/website/types"

	"github.com/gin-gonic/gin"
)

type CategoryAPI struct{}

func (st *CategoryAPI) GetCategory(c *gin.Context) {
	var pid request.ProductID
	err := c.ShouldBindJSON(&pid)

	fmt.Println("productId", pid)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := categoryService.GetCategoriesByProductID(pid.ProductID)
	if err != nil {
		response.FailWithMessage("获取产品分类失败", c)
		return
	}

	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (st *CategoryAPI) GetCategoryList(c *gin.Context) {
	list, err := tagsService.GetTagByTypeDB(0)
	if err != nil {
		response.FailWithMessage("获取产品分类失败", c)
		return
	}

	response.OkWithDetailed(list, "OK", c)
}

func (st *CategoryAPI) GetStyleList(c *gin.Context) {
	list, err := tagsService.GetTagByTypeDB(4)
	if err != nil {
		response.FailWithMessage("获取产品分类失败", c)
		return
	}

	response.OkWithDetailed(list, "OK", c)
}

func (st *CategoryAPI) GetShapeList(c *gin.Context) {
	list, err := tagsService.GetTagByTypeDB(5)
	if err != nil {
		response.FailWithMessage("获取产品分类失败", c)
		return
	}
	response.OkWithDetailed(list, "OK", c)
}

func (st *CategoryAPI) GetMenuList(c *gin.Context) {
	list, err := tagsService.GetTagByTypeForMenuDB(7)
	if err != nil {
		response.FailWithMessage("获取产品分类失败", c)
		return
	}
	response.OkWithDetailed(list, "OK", c)
}

func (st *CategoryAPI) CreateCategory(c *gin.Context) {
	var req request.CategoryCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = categoryService.CreateCategory(website.CategoryInfo{
		Title:     req.Title,
		ProductID: req.ProductID,
		Value:     req.ParentId,
		Level:     req.Level,
	})
	if err != nil {
		response.FailWithMessage("创建产品分类失败", c)
		return
	}

	response.OkWithMessage("创建产品分类成功", c)

}

func (st *CategoryAPI) GetTags(c *gin.Context) {
}
