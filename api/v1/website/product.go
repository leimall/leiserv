package website

import (
	"fmt"
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"
	website "leiserv/models/website/types"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct{}

func (p *ProductAPI) GetProduct(c *gin.Context) {
	var pageinfo webauthReq.PageInfo
	err := c.ShouldBindQuery(&pageinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println("GetProduct", pageinfo)

	list, total, err := productService.GetProductList(pageinfo)
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}

	var productIDs []string
	for _, product := range list {
		productIDs = append(productIDs, product.ProductID)
	}

	fmt.Println("productIDs", productIDs)

	// get image list
	imageList, err := productImgService.GetImageListForProduct(productIDs)
	if err != nil {
		response.FailWithMessage("获取商品列表中的图片失败", c)
		return
	}

	// get category list
	catagoryList, err := categoryService.GetCatagoryListForProduct(productIDs)
	if err != nil {
		response.FailWithMessage("获取商品列表中的分类失败", c)
		return
	}

	for i, product := range list {
		list[i].Category = catagoryList[product.ProductID]
		list[i].ImageList = imageList[product.ProductID]
	}

	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (p *ProductAPI) CreateProduct(c *gin.Context) {
	var cp webauthReq.CreateProductRequest
	if err := c.ShouldBindJSON(&cp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// TODO: create product
	np := &website.Produce{
		Title:          cp.Title,
		Desction:       cp.Description,
		SeoKeywords:    cp.SeoKeywords,
		SeoDescription: cp.SeoDescription,
		Price:          cp.Price,
		PriceOff:       cp.PriceOff,
		MainImg:        cp.MainImg,
		Stock:          cp.Stock,
	}
	err := productService.CreateProductOne(*np)
	if err != nil {
		response.FailWithMessage("商品创建失败", c)
		return
	}
	response.OkWithMessage("商品创建成功", c)

}

func (p *ProductAPI) UpdateProduct(c *gin.Context) {
	var up website.Produce
	if err := c.ShouldBindJSON(&up); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := productService.UpdateProductOne(up)

	if err != nil {
		response.FailWithMessage("商品更新失败", c)
		return
	}
	response.OkWithMessage("商品更新成功", c)

}

func (p *ProductAPI) DeleteProduct(c *gin.Context) {

}

func (p *ProductAPI) GetProductDetail(c *gin.Context) {

}
func (p *ProductAPI) GetProductSearch(c *gin.Context) {}
