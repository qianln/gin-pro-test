package main

import (
	"gin-pro/app/core/system"
	_ "gin-pro/bootstrap"
	"gin-pro/routers"
)

func main() {
	system.Config.ConfigFileChangeListen()

	r := routers.NewRouter()

	_ = r.Run(system.Config.GetString("HttpServer.Port"))
}
