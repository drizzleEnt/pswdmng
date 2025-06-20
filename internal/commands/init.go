package commands

import (
	"fmt"
)

func (r *Root) init(isNew bool) error {
	if isNew {
		if err := createNewAccount(r); err != nil {
			return err
		}
		return nil
	}

	ok, accounts, err := getExistingAccounts(r)
	if err != nil {
		return err
	}
	if !ok {
		err := createNewAccount(r)
		if err != nil {
			return err
		}
	} else {
		outputItems(accounts)
		index, err := getChosenItem(accounts)
		if err != nil {
			return err
		}

		fmt.Printf("current account: %v\n", accounts[index].Login)

		psw, err := getMasterPassword()
		if err != nil {
			return err
		}
		fmt.Printf("psw: %v\n", string(psw))

		if _, err := getLoginsAndUrls(r, accounts[index].Login); err != nil {
			return err
		}

	}
	return nil
}

func createNewAccount(r *Root) error {
	fmt.Println("Passwords files not exist")
	fmt.Println("Create login")
	login, err := getInput()
	if err != nil {
		return err
	}

	pswd, err := getMasterPassword()
	if err != nil {
		return err
	}

	fmt.Printf("pswd: %v\n", pswd)

	if err := r.repo.CreateFile(login); err != nil {
		return err
	}
	return nil
}
