package cobra_command

import (
	"fmt"
	"gin-pro/app/core/system"
	"gin-pro/library/cobra_command/system_command"
	"github.com/spf13/cobra"
)

func NewCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gin-pro-cli",
		Short: "gin-pro-cli start",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gin-pro-cli %s args:%v \n", cmd.Name(), args)
		},
	}
}

// 系统注册 artisan 命令

func Command() {
	system.CobraCommand.AddCommand(
		system_command.NewMakeController().AddCommand().CobraCommand,
		system_command.NewMakeValidator().AddCommand().CobraCommand,
	)
}
