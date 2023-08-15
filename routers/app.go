package routers

import (
	"gin-pro/app/core/system"
	"gin-pro/app/http/middleware/cors"
	"gin-pro/library/release_router"
	"gin-pro/routers/open"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templateFiles []string

func NewRouter() *gin.Engine {
	router := initRouter()

	router.Static("/public", "./public")               // 定义静态资源路由与实际目录映射关系
	router.StaticFile("/readme", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	loadTemplateFiles("./template", ".html")
	router.LoadHTMLFiles(templateFiles...)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "QIANLN_GIN_PRO",
		})
	})

	registerRouter(router)

	return router
}

func initRouter() *gin.Engine {
	var router *gin.Engine

	if system.Config.GetBool("AppDebug") == true {
		router = gin.Default()
	} else {
		router = release_router.NewReleaseRouter()
	}

	if system.Config.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	return router
}

func registerRouter(e *gin.Engine) {
	// Todo 注册其他分组下的路由
	open.LoadDemoRouter(e)
}

func loadTemplateFiles(path, stuffix string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, stuffix) {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
}
