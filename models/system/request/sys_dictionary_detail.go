package request

import (
	"leiserv/models/common/request"
	"leiserv/models/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
