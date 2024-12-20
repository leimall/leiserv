package api

import (
	"leiserv/global"
	"leiserv/models/common/response"
{{ if .NeedModel }}	"leiserv/plugin/{{ .Snake}}/model" {{ end }}
	"leiserv/plugin/{{ .Snake}}/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type {{ .PlugName}}Api struct{}

// @Tags {{ .PlugName}}
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /{{ .RouterGroup}}/routerName [post]
func (p *{{ .PlugName}}Api) ApiName(c *gin.Context) {
    {{ if .HasRequest}}
        var plug model.Request
        _ = c.ShouldBindJSON(&plug)
    {{ end }}
        if {{ if .HasResponse }} res, {{ end }} err:= service.ServiceGroupApp.PlugService({{ if .HasRequest }}plug{{ end -}}); err != nil {
		global.MALL_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
	{{if .HasResponse }}
	    response.OkWithDetailed(res,"成功",c)
	{{else}}
	    response.OkWithData("成功", c)
	{{ end -}}

	}
}
