package service

import (
	"leiserv/service/example"
	"leiserv/service/system"
	"leiserv/service/website"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	WebsiteServiceGroup website.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
