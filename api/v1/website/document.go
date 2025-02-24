package website

import (
	"leiserv/models/common/response"
	webauthReq "leiserv/models/website/request"

	"github.com/gin-gonic/gin"
)

type DocumentAPI struct{}

func (p *DocumentAPI) GetDocumentList(c *gin.Context) {
	var doc webauthReq.DocResponse
	err := c.ShouldBindQuery(&doc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	document, err1 := documentService.GetDocumentByTitleFromDB(doc.Title)
	if err1 != nil {
		response.FailWithMessage(err1.Error(), c)
		return
	}
	response.OkWithDetailed(document, "OK", c)
}
