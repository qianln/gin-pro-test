package open

import (
	"gin-pro/app/http/controllers"
	"gin-pro/app/modules/models"
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

}
