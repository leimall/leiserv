package utils

import (
	"leiserv/global"

	"github.com/jaevor/go-nanoid"
)

func GetNanoID() string {

	uuid, err := nanoid.CustomASCII(global.MALL_CONFIG.NanoID.ASCII, global.MALL_CONFIG.NanoID.Length)
	if err != nil {
		return err.Error()
	}
	return uuid()
}
