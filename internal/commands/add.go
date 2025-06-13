package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use:   "add",
	Short: "add row in password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add cmd")
		fmt.Printf("args: %v\n", args)
	},
}
