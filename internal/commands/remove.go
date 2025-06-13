package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removecmd = &cobra.Command{
	Use:   "remove",
	Short: "remove row from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove cmd")
	},
}
