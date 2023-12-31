package open

import (
	"gin-pro/app/http/controllers/open"
	"gin-pro/app/http/middleware/cors"
	"gin-pro/app/http/validator/open/demo"
	"github.com/gin-gonic/gin"
)

func LoadDemoRouter(e *gin.Engine) {

	d := e.Group("/open/demo").Use(cors.Next())
	{
		d.GET("/index", demo.NewIndexValidator().CheckParams, open.NewDemoController().Index)
	}
}
