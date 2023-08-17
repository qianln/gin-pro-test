package commands

import (
	"gin-pro/app/core/system"
	"github.com/spf13/cobra"
)

func NewPrintLog() *PrintLog {
	return &PrintLog{}
}

type PrintLog struct {
	CobraCommand *cobra.Command
}

func (p *PrintLog) AddCommand() *PrintLog {
	p.CobraCommand = &cobra.Command{
		Use:   "command:PrintLog", // 执行命令
		Short: "打印日志",             // 注释
		Run: func(cmd *cobra.Command, args []string) {
			p.Handle()
		},
	}
	return p
}

func (p *PrintLog) Handle() {
	system.ZapLog.Info("NewPrintLog sleep print")
}
