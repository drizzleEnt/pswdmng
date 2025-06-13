package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "pass",
	Short: "password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func InitCommands() {
	rootcmd.AddCommand(initcmd)
	rootcmd.AddCommand(logincmd)
	rootcmd.AddCommand(addcmd)
	rootcmd.AddCommand(getcmd)
	rootcmd.AddCommand(listcmd)
	rootcmd.AddCommand(removecmd)
}
