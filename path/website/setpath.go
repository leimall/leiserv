package website

import (
	middleware "leiserv/middleware/website"
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
	websiteRouter.InitProductRouter(PublicPack)
	websiteRouter.InitCommonRouter(PublicPack)

	// category api
	websiteRouter.InitCategoryRouter(PublicPack)

}