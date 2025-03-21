package request

type CommentPost struct {
	ID          int    `json:"id" example:"ID"`
	OrderID     string `json:"orderId" example:"订单ID"`
	ProductID   string `json:"productId" example:"产品ID"`
	UserID      string `json:"userId" example:"用户ID"`
	UserName    string `json:"userName" example:"用户名"`
	Title       string `json:"title" example:"标题"`
	Content     string `json:"content" example:"评论内容"`
	Star        int    `json:"star" example:"评分"`
	IsImg       bool   `json:"isImg" example:"是否有图片"`
	ImgUrl      string `json:"imgUrl" example:"图片地址"`
	ShopContent string `json:"shopContent" example:"商家回复内容"`
}
