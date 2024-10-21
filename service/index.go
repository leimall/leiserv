package service

import (
	"leiserv/service/website"
)

type ServiceGroup struct {
	WebsiteServiceGroup website.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
