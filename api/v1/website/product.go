package website

import (
	"fmt"
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"
	website "leiserv/models/website/types"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct{}

func (p *ProductAPI) GetAllProductList(c *gin.Context) {
	list, err := productService.GetAllProductListDB()
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}
	response.OkWithDetailed(list, "OK", c)
}

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

// get all list of product
func (p *ProductAPI) GetAllProductListsForSearch(c *gin.Context) {
	list, err := productService.GetAllProductListDBForSearch()
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return

	}
	var productIDs []string
	for _, product := range list {
		productIDs = append(productIDs, product.ProductID)
	}
	// get category list
	catagoryList, err := categoryService.GetCatagoryListForProduct(productIDs)
	if err != nil {
		response.FailWithMessage("获取商品列表中的分类失败", c)
		return
	}

	// get tags lists
	tagsList, err := tagsService.GetTagListForProduct(productIDs)
	if err != nil {
		response.FailWithMessage("获取商品列表中的标签失败", c)
		return
	}

	// get review list
	reviewList, err := productReviewService.GetProductReviewByProductIDDB(productIDs)
	if err != nil {
		response.FailWithMessage("获取产品评论失败", c)
		return
	}

	for i, product := range list {
		list[i].Category = catagoryList[product.ProductID]
		list[i].Tags = tagsList[product.ProductID]
		list[i].Review = reviewList[product.ProductID]
	}

	response.OkWithDetailed(list, "OK", c)
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
func (p *ProductAPI) GetProductSearch(c *gin.Context) {
	var search webauthReq.SearchRequest
	err := c.ShouldBindQuery(&search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	catagoryLists, err := categoryService.GetProductListBySearchDB(search)
	if err != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}
	// 1. get product id list from list
	var pids []string

	for _, product := range catagoryLists {
		pids = append(pids, product.ProductID)
	}

	lists, err := productService.GetProductListByCategoryDB(pids, 99999)
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

// get product main page hot product list
func (p *ProductAPI) GetBestSellerProductList(c *gin.Context) {
	var best webauthReq.BestRequest
	err := c.ShouldBindQuery(&best)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	tagsList, err := tagsService.GetTagListByTitleDB(best)
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

	response.OkWithDetailed(product, "OK", c)
}

// get product list by category id
// 1. catgory title get product id from category_info table
// 2. get product id list from product table
func (p *ProductAPI) GetProductListByCategory(c *gin.Context) {
	categoryID := c.Query("categoryId")
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.FailWithMessage("offset 参数格式错误", c)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.FailWithMessage("limit 参数格式错误", c)
		return
	}

	ctitle := strings.ReplaceAll(categoryID, "_", " ")
	list, err := categoryService.GetCategoryByTitleDB(ctitle)
	if err != nil {
		response.FailWithMessage("商品分类获取失败", c)
		return
	}

	var productIDs []string

	for _, product := range list {
		productIDs = append(productIDs, product.ProductID)
	}

	pids := laginatedlimitProdiuctIDS(productIDs, offset, limit)

	plist, err := productService.GetProductListByCategoryDB(pids, limit)
	if err != nil {
		response.FailWithMessage("商品列表获取失败", c)
		return
	}

	total := len(list)
	response.OkWithDetailed(response.ListsResult{
		List:  plist,
		Total: int64(total),
	}, "OK", c)
}

func laginatedlimitProdiuctIDS(productIDs []string, offset int, limit int) []string {
	// 计算切片的起始和结束位置
	start := offset
	end := offset + limit
	if start > len(productIDs) {
		start = len(productIDs)
	}
	if end > len(productIDs) {
		end = len(productIDs)
	}
	paginatedProductIDs := productIDs[start:end]

	return paginatedProductIDs
}

// get product sku by product id
func (p *ProductAPI) GetProductSkuByID(c *gin.Context) {
	pid := c.Param("id")
	skuList, err := skuService.GetSkuListForProduct(pid)
	if err != nil {
		response.FailWithMessage("商品SKU获取失败", c)
		return
	}
	response.OkWithDetailed(skuList, "OK", c)
}

func (p *ProductAPI) GetAllTagName(c *gin.Context) {
	review, err := tagsService.GetSKUAllList(1)
	if err != nil {
		response.FailWithMessage("所有标签列表获取失败", c)
		return
	}
	response.OkWithDetailed(review, "OK", c)
}
