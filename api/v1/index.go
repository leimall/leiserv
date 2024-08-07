package v1

import (
	"leiserv/api/v1/example"
	"leiserv/api/v1/system"
	"leiserv/api/v1/website"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup

	WebSiteAPIPack website.APIPack
}

var ApiGroupApp = new(ApiGroup)
