package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logincmd = &cobra.Command{
	Use:   "login",
	Short: "login in password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login cmd")
	},
}
