package request

import (
	"leiserv/models/common/request"
	"leiserv/models/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
