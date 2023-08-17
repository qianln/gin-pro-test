package cobra_command

import (
	"fmt"
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
