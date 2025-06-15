package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) get(cmd *cobra.Command, args []string) {
	ok, accounts := getExistingAccounts(r)
	if !ok {
		fmt.Println("At first you need to initialize your first account")
		return
	}

	accountIndex := getChosenAccount(accounts)

	fmt.Printf("current account: %v\n", accounts[accountIndex])

	psw, err := getPassword()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("psw: %v\n", string(psw))

	entries, err := r.repo.List(accounts[accountIndex])
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Url | Login\n")
	for i, e := range entries {
		fmt.Printf("%v: %v - %v\n", i+1, e[1], e[0])
	}

	rowIndex := getChosenLogin(entries)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	pswd, err := r.repo.Get(accounts[accountIndex], entries[rowIndex][1], entries[rowIndex][0])
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("pswd: %v\n", pswd)
	return
}
