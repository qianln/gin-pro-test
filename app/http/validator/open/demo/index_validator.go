package demo

import (
	"gin-pro/app/http"
	"gin-pro/app/utils/response"
	"github.com/gin-gonic/gin"
)

func NewIndexValidator() *IndexValidator {
	return &IndexValidator{}
}

type IndexValidator struct {
}

func (i IndexValidator) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&i); err != nil {
		response.ValidatorError(context, err)
		return
	}

	if http.BindContext(i, context) == nil {
		response.ErrorSystem(context, " index 表单验证器json化失败", "")
	} else {
		context.Next()
	}
}
