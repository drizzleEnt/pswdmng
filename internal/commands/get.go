package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getcmd = &cobra.Command{
	Use:   "get",
	Short: "get row from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get cmd")
	},
}
