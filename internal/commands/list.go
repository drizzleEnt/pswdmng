package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listcmd = &cobra.Command{
	Use:   "list",
	Short: "get rows list from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list cmd")
	},
}
