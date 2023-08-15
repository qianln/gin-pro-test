package demo

import (
	"gin-pro/app/global/consts"
	"gin-pro/app/http"
	"gin-pro/app/http/validator/common_data"
	"gin-pro/app/utils/response"
	"github.com/gin-gonic/gin"
)

func NewIndexValidator() *IndexValidator {
	return &IndexValidator{}
}

type IndexValidator struct {
	common_data.IntId
}

func (c IndexValidator) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&c); err != nil {
		response.ValidatorError(context, err)
		return
	}

	if http.BindContext(c, consts.Param, context) == nil {
		response.ErrorSystem(context, " index 表单验证器json化失败", "")
	} else {
		context.Next()
	}
}
