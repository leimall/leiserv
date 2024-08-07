package initialize

import (
	_ "leiserv/source/example"
	_ "leiserv/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
