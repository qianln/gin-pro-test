package console

import (
	"gin-pro/app/console/commands"
	"gin-pro/app/core/system"
	"github.com/robfig/cron/v3"
)

func Command() {
	system.CobraCommand.AddCommand(
		commands.NewPrintLog().AddCommand().CobraCommand,
	)
}

func Schedule() {
	c := cron.New()

	//	执行规则
	//	@every 5s 每几秒执行一次
	//	支持linux crontab 命令

	c.AddFunc("@every 2s", commands.NewPrintLog().Handle)

	if system.Config.GetBool("AppDebug") == false {
		go c.Start()
		defer c.Stop()
	}
}
