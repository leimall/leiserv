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
	ProductCommentAPI
	ReviewAPI
	CartAPI
	OrdersApi
	DocumentAPI
	ShippingAPI
	LLPayAPI
}

var (
	jwtWebService         = service.ServiceGroupApp.WebsiteServiceGroup.JWTWebService
	userService           = service.ServiceGroupApp.WebsiteServiceGroup.UserService
	commonService         = service.ServiceGroupApp.WebsiteServiceGroup.CommonService
	categoryService       = service.ServiceGroupApp.WebsiteServiceGroup.CategoryService
	addressService        = service.ServiceGroupApp.WebsiteServiceGroup.AddressService
	billingAddressService = service.ServiceGroupApp.WebsiteServiceGroup.BillingAddressService
	productService        = service.ServiceGroupApp.WebsiteServiceGroup.ProductService
	productImgService     = service.ServiceGroupApp.WebsiteServiceGroup.ProductImgService
	tagsService           = service.ServiceGroupApp.WebsiteServiceGroup.TagsService
	skuService            = service.ServiceGroupApp.WebsiteServiceGroup.SkuService
	productDetailService  = service.ServiceGroupApp.WebsiteServiceGroup.ProductDetailService
	productReviewService  = service.ServiceGroupApp.WebsiteServiceGroup.ProductReviewService
	productBrandService   = service.ServiceGroupApp.WebsiteServiceGroup.ProductBrandService
	productCommentService = service.ServiceGroupApp.WebsiteServiceGroup.ProductCommentService
	cartService           = service.ServiceGroupApp.WebsiteServiceGroup.CartService
	ordersService         = service.ServiceGroupApp.WebsiteServiceGroup.OrdersService
	documentService       = service.ServiceGroupApp.WebsiteServiceGroup.DocumentService
	yanwenService         = service.ServiceGroupApp.WebsiteServiceGroup.YanWenService
	lianlianpayService    = service.ServiceGroupApp.WebsiteServiceGroup.LLPayService
	PaymentLlPayService   = service.ServiceGroupApp.WebsiteServiceGroup.PaymentLLPayService
)
