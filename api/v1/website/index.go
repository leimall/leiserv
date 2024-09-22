package website

import "leiserv/service"

type APIPack struct {
	AuthAPI
	AddressAPI
	CaptchaAPI
	CategoryAPI
	CommonAPI
	ProductAPI
	ProductImgAPI
	ReviewAPI
	CartAPI
}

var (
	jwtWebService        = service.ServiceGroupApp.WebsiteServiceGroup.JWTWebService
	userService          = service.ServiceGroupApp.WebsiteServiceGroup.UserService
	commonService        = service.ServiceGroupApp.WebsiteServiceGroup.CommonService
	categoryService      = service.ServiceGroupApp.WebsiteServiceGroup.CategoryService
	addressService       = service.ServiceGroupApp.WebsiteServiceGroup.AddressService
	productService       = service.ServiceGroupApp.WebsiteServiceGroup.ProductService
	productImgService    = service.ServiceGroupApp.WebsiteServiceGroup.ProductImgService
	tagsService          = service.ServiceGroupApp.WebsiteServiceGroup.TagsService
	skuService           = service.ServiceGroupApp.WebsiteServiceGroup.SkuService
	productDetailService = service.ServiceGroupApp.WebsiteServiceGroup.ProductDetailService
	productReviewService = service.ServiceGroupApp.WebsiteServiceGroup.ProductReviewService
	cartService          = service.ServiceGroupApp.WebsiteServiceGroup.CartService
)
