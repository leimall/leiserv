package v1

import (
	"leiserv/api/v1/website"
)

type ApiGroup struct {
	WebSiteAPIPack website.APIPack
}

var ApiGroupApp = new(ApiGroup)
