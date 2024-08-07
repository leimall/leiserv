package request

type CommonID struct {
	CommonID int `json:"id" example:"通用 ID Int 类型"`
}

type ProductID struct {
	ProductID string `json:"productId" example:"产品ID"`
}

type IDandPID struct {
	ID        int    `json:"id" example:"ID"`
	ProductID string `json:"productId" example:"产品ID"`
}
type SortIDandPID struct {
	ID        int    `json:"id" example:"ID"`
	ProductID string `json:"productId" example:"产品ID"`
	SortID    int    `json:"sortId" example:"排序ID"`
}

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}
