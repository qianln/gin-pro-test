package open

import (
	"gin-pro/app/global/consts"
	"gin-pro/app/http/controllers"
	"gin-pro/app/modules/models"
	"gin-pro/app/utils/response"
	"github.com/gin-gonic/gin"
)

func NewDemoController() *DemoController {
	return &DemoController{}
}

type DemoController struct {
	controllers.BaseController
	uModel  models.UsersModel
	usModel []models.UsersModel
}

func (d DemoController) Index(c *gin.Context) {
	version := c.GetString(consts.Param + "version")
	response.Success(c, consts.HttpStatusOkMsg, gin.H{
		"say":     "Hello GinPro",
		"version": version,
	})
}
