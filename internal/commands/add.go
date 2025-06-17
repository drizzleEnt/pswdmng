package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) add(cmd *cobra.Command, args []string) {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	login, err := cmd.Flags().GetString("login")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	if err := r.repo.Add(account, login, url, "123"); err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
