package router

import (
	"leiserv/router/website"
)

type RouterPack struct {
	Website website.RouterPack
}

var RouterPackApp = new(RouterPack)
