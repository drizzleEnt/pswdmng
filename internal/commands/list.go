package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) list(cmd *cobra.Command, args []string) {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	entries, err := r.repo.List(account)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Url | Login\n")
	for i, e := range entries {
		fmt.Printf("%v: %v - %v\n", i+1, e[1], e[0])
	}
}

func getAccountAndPassword(r *Root) (string, string, error) {
	ok, accounts := getExistingAccounts(r)
	if !ok {
		return "", "", fmt.Errorf("At first you need to initialize your first account")
	}

	accountIndex := getChosenAccount(accounts)

	fmt.Printf("current account: %v\n", accounts[accountIndex])

	psw, err := getPassword()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("psw: %v\n", string(psw))

	return accounts[accountIndex], string(psw), nil
}
