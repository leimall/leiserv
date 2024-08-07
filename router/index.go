package router

import (
	"leiserv/router/example"
	"leiserv/router/system"
	"leiserv/router/website"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

type RouterPack struct {
	Website website.RouterPack
}

var RouterPackApp = new(RouterPack)
