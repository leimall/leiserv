package website

import (
	middleware "leiserv/middleware"
	"leiserv/router"

	"github.com/gin-gonic/gin"
)

type SetPath struct{}

func (s *SetPath) InitSetPath(Router *gin.Engine) {
	websiteRouter := router.RouterPackApp.Website

	PublicPack := Router.Group("/api/web")
	PrivatePack := Router.Group("/api/web")

	PrivatePack.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	// signup new client myself
	websiteRouter.InitAuthRouter(PublicPack, PrivatePack)
	websiteRouter.InitAddressRouter(PrivatePack)

	// public and common api
	websiteRouter.InitProductRouter(PublicPack, PrivatePack)

	websiteRouter.InitCommonRouter(PublicPack)

	// category api
	websiteRouter.InitCategoryRouter(PublicPack)

	// cart api
	websiteRouter.InitCartRouter(PublicPack, PrivatePack)

	// order api
	websiteRouter.InitOrdersRouter(PublicPack, PrivatePack)

	// document api
	websiteRouter.InitDocumentRouter(PublicPack)

	// lianlian pay api
	websiteRouter.InitLianLiangPayRouter(PublicPack)

}
