package artisan_template

import (
	"fmt"
	"strings"
)

type MakeController struct {
	BasePath     string
	ClassName    string
	SubClassName string
	Content      string
}

func DoClassName(className string) (string, string) {
	fistName := strings.ToUpper(string(className[0]))
	className = fmt.Sprintf("%s%s", fistName, className[1:])
	if strings.HasSuffix(className, "Controller") == false {
		className = fmt.Sprintf("%sController", className)
	}

	subClassName := strings.ToLower(string(className[0]))
	return className, subClassName
}

func NewMakeController(basePath, name string) *MakeController {
	className, subClassName := DoClassName(name)
	return &MakeController{
		BasePath:     basePath,
		ClassName:    className,
		SubClassName: subClassName,
	}
}

func (m *MakeController) GetContent() string {
	return `package open

import (
	"gin-pro/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func New` + m.ClassName + `() *` + m.ClassName + ` {
	return &` + m.ClassName + `{}
}

type ` + m.ClassName + ` struct {
	controllers.BaseController
}

func (` + m.SubClassName + ` ` + m.ClassName + `) Index(c *gin.Context) {

}
`
}
