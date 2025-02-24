package website

import (
	"fmt"
	"leiserv/models/common/response"
	"leiserv/utils"

	website "leiserv/models/website/types"

	"github.com/gin-gonic/gin"
)

type CartAPI struct{}

func (a *CartAPI) AddCart(c *gin.Context) {
	var cart website.CartItem
	if err := c.ShouldBindJSON(&cart); err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
		return
	}

	err := cartService.AddCartDB(cart)
	if err != nil {
		response.FailWithMessage("添加失败", c)
	}
	response.OkWithMessage("添加成功", c)
}

func (a *CartAPI) DeleteCart(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	err := cartService.DeleteCartDB(userId)
	if err != nil {
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)
}

func (a *CartAPI) DeleteCartOne(c *gin.Context) {
	var req website.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
		return
	}
	err := cartService.DeleteCartOneDB(req.UniqueId)
	if err != nil {
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)
}

func (a *CartAPI) UpdateCart(c *gin.Context) {
	var req website.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	}
	err := cartService.UpdateCartDB(req)
	if err != nil {
		response.FailWithMessage("更新失败", c)
	}
	response.OkWithMessage("更新成功", c)
}

func (a *CartAPI) GetCartList(c *gin.Context) {
	userId := utils.GetWebUserID(c)
	fmt.Println("userID:", userId)

	cart, err := cartService.GetCartDB(userId)

	if err != nil {
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(cart, "OK", c)
}
