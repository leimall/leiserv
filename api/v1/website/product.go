package website

import (
	"fmt"
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"strings"

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

// get product main page hot product list
func (p *ProductAPI) GetBestSellerProductList(c *gin.Context) {
	num := 0
	tagsList, err := tagsService.GetTagListByTitleDB(num, "Best Seller")
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}
	var pids []string

	for _, product := range tagsList {
		pids = append(pids, product.ProductID)
	}

	lists, err := productService.GetProductListByCategoryDB(pids, 0)
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}

	reviewList, err := productReviewService.GetProductReviewByProductIDDB(pids)
	if err != nil {
		response.FailWithMessage("获取产品评论失败", c)
		return
	}
	for i, product := range lists {
		lists[i].Review = reviewList[product.ProductID]
	}

	response.OkWithDetailed(lists, "OK", c)
}

// get product main page sale product list

func (p *ProductAPI) GetSaleProductList(c *gin.Context) {
	num := 0
	tagsList, err := tagsService.GetTagListByTitleDB(num, "sale")
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}
	var pids []string

	for _, product := range tagsList {
		pids = append(pids, product.ProductID)
	}

	lists, err := productService.GetProductListByCategoryDB(pids, 0)
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}

	reviewList, err := productReviewService.GetProductReviewByProductIDDB(pids)
	if err != nil {
		response.FailWithMessage("获取产品评论失败", c)
		return
	}
	for i, product := range lists {
		lists[i].Review = reviewList[product.ProductID]
	}

	response.OkWithDetailed(lists, "OK", c)
}

// get product main page lastest product list
func (p *ProductAPI) GetLastestProductList(c *gin.Context) {
	num := 0
	lists, err := productService.GetLastestProductListDB(num)
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}

	// 1. get product id list from list
	pid := make([]string, 0, len(lists.([]website.ProductListItme)))
	for _, product := range lists.([]website.ProductListItme) {
		pid = append(pid, product.ProductID)
	}

	reviewList, err := productReviewService.GetProductReviewByProductIDDB(pid)
	if err != nil {
		response.FailWithMessage("获取产品评论失败", c)
		return
	}

	for i, product := range lists.([]website.ProductListItme) {
		lists.([]website.ProductListItme)[i].Review = reviewList[product.ProductID]
	}

	response.OkWithDetailed(lists, "OK", c)
}

func (p *ProductAPI) GetProductDetailById(c *gin.Context) {
	pid := c.Param("id")
	product, err := productService.GetProductDetailByPidServ(pid)
	if err != nil {
		response.FailWithMessage("商品获取失败", c)
		return
	}

	imageList, err := productImgService.GetImageListForProduct([]string{product.ProductID})
	if err != nil {
		response.FailWithMessage("商品图片获取失败", c)
		return
	}
	product.ImageList = imageList[product.ProductID]

	catagoryList, err := categoryService.GetCatagoryListForProduct([]string{product.ProductID})
	if err != nil {
		response.FailWithMessage("商品分类获取失败", c)
		return
	}
	product.Category = catagoryList[product.ProductID]

	tagsList, err := tagsService.GetTagListForProduct([]string{product.ProductID})
	if err != nil {
		response.FailWithMessage("商品标签获取失败", c)
		return
	}
	product.Tags = tagsList[product.ProductID]

	skuList, err := skuService.GetSkuListForProduct(pid)
	if err != nil {
		response.FailWithMessage("商品SKU获取失败", c)
		return
	}
	product.Sku = skuList

	productDetail, err := productDetailService.GetProductDetailByPidDB(pid)
	if err != nil {
		response.FailWithMessage("商品详情获取失败", c)
		return
	}
	product.Detail = productDetail

	productReview, err := productReviewService.GetProductReviewByIDDB(pid)
	if err != nil {
		response.FailWithMessage("商品评价获取失败", c)
		return
	}
	product.Review = productReview

	productBrand, err := productBrandService.GetProductBrandByIDDB(pid)
	if err != nil {
		response.FailWithMessage("商品品牌获取失败", c)
		return
	}
	product.Brand.Info = productBrand

	productBrandTags, err := tagsService.GetProductBrandTagsByIDDB(productBrand.ShapeID)
	if err != nil {
		response.FailWithMessage("商品品牌标签获取失败", c)
		return
	}
	for i, tag := range productBrandTags {
		children, err := tagsService.GetProductBrandChildrenTagsByIDDB(tag.ID)
		if err != nil {
			response.FailWithMessage("商品品牌Size获取失败", c)
			return
		}
		productBrandTags[i].Children = children

	}
	product.Brand.Tags = productBrandTags

	response.OkWithDetailed(product, "OK", c)
}

// get product list by category id
// 1. catgory title get product id from category_info table
// 2. get product id list from product table
func (p *ProductAPI) GetProductListByCategory(c *gin.Context) {
	cid := c.Param("id")
	ctitle := strings.ReplaceAll(cid, "_", " ")
	list, err := categoryService.GetCategoryByTitleDB(ctitle)
	if err != nil {
		response.FailWithMessage("商品分类获取失败", c)
		return
	}

	var productIDs []string

	for _, product := range list {
		productIDs = append(productIDs, product.ProductID)
	}

	plist, err := productService.GetProductListByCategoryDB(productIDs, 0)
	if err != nil {
		response.FailWithMessage("商品列表获取失败", c)
		return
	}

	reviewList, err := productReviewService.GetProductReviewByProductIDDB(productIDs)
	if err != nil {
		response.FailWithMessage("获取产品评论失败", c)
		return
	}

	for i := range plist {
		plist[i].Review = reviewList[plist[i].ProductID]
	}

	response.OkWithDetailed(plist, "OK", c)
}
