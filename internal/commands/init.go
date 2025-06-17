package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) init(cmd *cobra.Command, args []string) {
	newFlag, _ := cmd.Flags().GetBool("new")
	if newFlag {
		if err := createNewAccount(r); err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		return
	}

	ok, accounts := getExistingAccounts(r)
	if !ok {
		createNewAccount(r)
	} else {
		index, err := getChosenItem(accounts)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("current account: %v\n", accounts[index].Login)

		psw, err := getPassword()
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("psw: %v\n", string(psw))

		if _, err := getLoginsAndUrls(r, accounts[index].Login); err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		return
	}
}

func createNewAccount(r *Root) error {
	fmt.Println("Passwords files not exist")
	fmt.Println("Create login")
	login, err := getInput()
	if err != nil {
		return err
	}

	if err := r.repo.CreateFile(login); err != nil {
		return err
	}
	return nil
}
