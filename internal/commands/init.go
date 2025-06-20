package commands

import (
	"fmt"
	"os"
)

func (r *Root) init(isNew bool) {
	if isNew {
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
		outputItems(accounts)
		index, err := getChosenItem(accounts)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("current account: %v\n", accounts[index].Login)

		psw, err := getMasterPassword()
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
