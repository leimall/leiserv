package website

import (
	v1 "leiserv/api/v1"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (a *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("product")

	productAPI := v1.ApiGroupApp.WebSiteAPIPack.ProductAPI
	productImgAPI := v1.ApiGroupApp.WebSiteAPIPack.ProductImgAPI
	productCommentAPI := v1.ApiGroupApp.WebSiteAPIPack.ProductCommentAPI
	{
		productRouter.GET("list", productAPI.GetProduct)
		productRouter.GET("detail", productAPI.GetProductDetail)
		productRouter.GET("search", productAPI.GetProductSearch)
	}

	// server api for admin create, update, delete product
	{
		productRouter.POST("create", productAPI.CreateProduct)
		productRouter.POST("update", productAPI.UpdateProduct)
		productRouter.POST("delete", productAPI.DeleteProduct)
	}

	// upload image for product
	{
		productRouter.POST("/image/:pid", productImgAPI.UploadImage)
		productRouter.POST("image", productImgAPI.GetImage)
		productRouter.DELETE("image", productImgAPI.DeleteImage)
		productRouter.POST("/image/setmain", productImgAPI.SetMainImage)
		productRouter.POST("/image/setsortid", productImgAPI.SetSortIdforImage)
	}

	// get product main page lastest product list
	{
		productRouter.GET("lastest", productAPI.GetLastestProductList)
		productRouter.GET("detail/:id", productAPI.GetProductDetailById)
	}

	// get product category list
	{
		productRouter.GET("category/:id", productAPI.GetProductListByCategory)
	}

	// get product comment list
	{
		productRouter.GET("comment", productCommentAPI.GetCommentListByProductID)
	}

}
