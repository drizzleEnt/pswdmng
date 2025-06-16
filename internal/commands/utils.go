package commands

import (
	"fmt"
	"os"
	"pswdmng/internal/domain"
	"strconv"

	"golang.org/x/term"
)

func getLoginsAndUrls(r *Root, account string) ([]domain.UserInfo, error) {
	entries, err := r.repo.List(account)
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("You did not add any login\n")
	}

	fmt.Printf("Url | Login\n")
	for i, e := range entries {
		fmt.Printf("%v: %v - %v\n", i+1, e.Url, e.Login)
	}

	return entries, nil
}

func getAccountAndPassword(r *Root) (string, string, error) {
	ok, accounts := getExistingAccounts(r)
	if !ok {
		return "", "", fmt.Errorf("At first you need to initialize your first account")
	}

	accountIndex, err := getChosenItem(accounts)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("current account: %v\n", accounts[accountIndex].Login)

	psw, err := getPassword()
	if err != nil {
		return "", "", err
	}

	return accounts[accountIndex].Login, string(psw), nil
}

func getChosenItem(items []domain.UserInfo) (int, error) {
	if len(items) == 1 {
		return 0, nil
	}

	fmt.Println("founded rows:")
	for i, l := range items {
		fmt.Printf("%v: %v\n", i+1, l.Login)
	}

	fmt.Print("Enter row number: ")
	chosenItem, err := getInput()
	if err != nil {
		return 0, err
	}

	itemIndex, err := strconv.Atoi(chosenItem)
	if err != nil {
		return 0, err
	}

	return itemIndex - 1, nil
}

func getExistingAccounts(r *Root) (bool, []domain.UserInfo) {
	fmt.Println("Check existing passwords files")

	ok, logins, err := r.repo.CheckExist()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}
	return ok, logins
}

func getInput() (string, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}

func getPassword() ([]byte, error) {
	fmt.Println("Enter master password for access to account:")
	psw, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return psw, nil
}
