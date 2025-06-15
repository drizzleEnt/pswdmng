package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (r *Root) add(cmd *cobra.Command, args []string) {
	ok, accounts := getExistingLogins(r)
	if !ok {
		fmt.Println("At first you need to initialize your first account")
		return
	}

	index := getChoosenAccount(accounts)

	fmt.Printf("current account: %v\n", accounts[index-1])

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

	if err := r.repo.Add(accounts[index-1], login, url); err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
