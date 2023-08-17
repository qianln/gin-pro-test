package system_command

import (
	"fmt"
	"gin-pro/app/core/system"
	"gin-pro/app/utils/helps"
	"gin-pro/library/cobra_command/artisan_template"
	"github.com/spf13/cobra"
	"strings"
)

func NewMakeController() *MakeController {
	return &MakeController{}
}

type MakeController struct {
	CobraCommand *cobra.Command
}

func (m *MakeController) AddCommand() *MakeController {
	m.CobraCommand = &cobra.Command{
		Use:   "make:controller",
		Short: "创建一个控制器",
		Long:  "创建一一一一一一个个个个个个个个控制器",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			m.Handle(args[0])
		},
	}
	return m
}

func (m *MakeController) Handle(args string) {

	basePath := "app/http/controllers" + "/" + args
	pathSlice := strings.Split(basePath, "/")
	basePathStr := strings.Join(pathSlice[:len(pathSlice)-1], "/")
	className, subClassName, fileName := m.doClassName(pathSlice[len(pathSlice)-1])
	c := artisan_template.NewControllerTemplate(
		basePathStr,
		pathSlice[len(pathSlice)-2],
		className,
		subClassName,
		fileName,
	)

	err := helps.CreateFile(c.BasePath, c.FileName, c.GetContent())
	if err != nil {
		system.ZapLog.Error(err.Error())
		return
	}

	fmt.Println("File created and content written successfully : " + c.BasePath + "/" + c.ClassName + ".go")
}
func (m *MakeController) doClassName(className string) (string, string, string) {
	className = helps.CamelString(className)
	subClassName := strings.ToLower(string(className[0]))

	if strings.HasSuffix(className, "Controller") == false {
		className = fmt.Sprintf("%sController", className)
	}
	fileName := helps.SnakeString(className)
	return className, subClassName, fileName
}
