package artisan_template

type ValidatorTemplate struct {
	BasePath     string
	ClassName    string
	SubClassName string
	Content      string
	Pack         string
	FileName     string
}

func NewValidatorTemplate(basePath, packName, className, subClassName, fileName string) *ValidatorTemplate {
	return &ValidatorTemplate{
		BasePath:     basePath,
		Pack:         packName,
		ClassName:    className,
		SubClassName: subClassName,
		FileName:     fileName,
	}
}

func (v *ValidatorTemplate) GetContent() string {
	return `package ` + v.Pack + `

import (
	"gin-pro/app/http"
	"gin-pro/app/utils/response"
	"github.com/gin-gonic/gin"
)

func New` + v.ClassName + `() *` + v.ClassName + ` {
	return &` + v.ClassName + `{}
}

type ` + v.ClassName + ` struct {
}

func (i ` + v.ClassName + `) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&i); err != nil {
		response.ValidatorError(context, err)
		return
	}

	if http.BindContext(i, context) == nil {
		response.ErrorSystem(context, " index 表单验证器json化失败", "")
	} else {
		context.Next()
	}
}`
}
