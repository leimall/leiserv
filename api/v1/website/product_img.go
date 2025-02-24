package website

import (
	"leiserv/models/common/response"
	websiteReq "leiserv/models/website/request"
	websiteRes "leiserv/models/website/response"
	website "leiserv/models/website/types"

	"github.com/gin-gonic/gin"
)

type ProductImgAPI struct{}

func (p *ProductImgAPI) GetImage(c *gin.Context) {
	var pid websiteReq.ProductID
	err := c.ShouldBindJSON(&pid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, mysqlErr := productImgService.GetProductImgList(pid.ProductID)
	if mysqlErr != nil {
		response.FailWithMessage("获取产品列表失败", c)
		return
	}
	response.OkWithDetailed(response.ListsResult{
		List:  list,
		Total: total,
	}, "OK", c)
}

func (p *ProductImgAPI) UploadImage(c *gin.Context) {
	pid := c.Param("pid")
	var file website.ProductImg
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = productImgService.SaveProductImg(header, noSave, pid)
	if err != nil {
		response.FailWithMessage("上传文件失败", c)
		return
	}
	response.OkWithDetailed(websiteRes.ProductImgResponse{File: file}, "上传成功", c)
}

func (p *ProductImgAPI) DeleteImage(c *gin.Context) {

}

func (p *ProductImgAPI) SetMainImage(c *gin.Context) {
	var file website.ProductImg

	var idandpid websiteReq.IDandPID
	err := c.ShouldBindJSON(&idandpid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productImgService.CancelMainImageDB(idandpid.ProductID)
	if err != nil {
		response.FailWithMessage("清除原主图失败", c)
		return
	}
	file, err = productImgService.SetMainImageDB(idandpid.ID, idandpid.ProductID)
	if err != nil {
		response.FailWithMessage("设置主图失败", c)
		return
	}
	err = productService.SetMainImageUrl(idandpid.ProductID, file.ImgURL)
	if err != nil {
		response.FailWithMessage("更新产品主图URL失败", c)
		return
	}
	response.OkWithMessage("设置主图, 并更新商品主图URL成功", c)
}

func (p *ProductImgAPI) SetSortIdforImage(c *gin.Context) {
	var spid websiteReq.SortIDandPID
	err := c.ShouldBindJSON(&spid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productImgService.SetSortIDforImageDB(spid.ID, spid.ProductID, spid.SortID)
	if err != nil {
		response.FailWithMessage("设置图片排序失败", c)
		return
	}
	response.OkWithMessage("设置图片排序成功", c)
}
