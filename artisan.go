package main

import (
	"gin-pro/app/console"
	"gin-pro/app/core/system"
	_ "gin-pro/bootstrap"
	"gin-pro/library/cobra_command"
)

func init() {
	cobra_command.Command()
	console.Command()
}

func main() {

	if err := system.CobraCommand.Execute(); err != nil {
		panic(err)
	}
}
