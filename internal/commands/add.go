package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (r *Root) add(cmd *cobra.Command, args []string) {
	ok, accounts := getExistingAccounts(r)
	if !ok {
		fmt.Println("At first you need to initialize your first account")
		return
	}

	index := getChosenAccount(accounts)

	fmt.Printf("current account: %v\n", accounts[index])

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

	if err := r.repo.Add(accounts[index], login, url); err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
