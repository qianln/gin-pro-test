package artisan_template

type ControllerTemplate struct {
	BasePath     string
	ClassName    string
	SubClassName string
	Content      string
	Pack         string
	FileName     string
}

func NewControllerTemplate(basePath, packName, className, subClassName, fileName string) *ControllerTemplate {
	return &ControllerTemplate{
		BasePath:     basePath,
		Pack:         packName,
		ClassName:    className,
		SubClassName: subClassName,
		FileName:     fileName,
	}
}

func (c *ControllerTemplate) GetContent() string {
	return `package ` + c.Pack + `

import (
	"github.com/gin-gonic/gin"
)

func New` + c.ClassName + `() *` + c.ClassName + ` {
	return &` + c.ClassName + `{}
}

type ` + c.ClassName + ` struct {
}

func (` + c.SubClassName + ` ` + c.ClassName + `) Index(c *gin.Context) {

}
`
}
