package path

import "leiserv/path/website"

type PathRouter struct {
	WebsitePathPack website.PathPack
}

var PathRouterAPP = new(PathRouter)
