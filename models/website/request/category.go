package request

type CategoryCreate struct {
	ProductID string `json:"productId" example:"产品ID"`
	Title     string `json:"title" example:"分类标题"`
	ParentId  int    `json:"parentId" example:"父分类ID"`
	Level     int    `json:"level" example:"分类级别"`
}
