package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func (r *Root) init(cmd *cobra.Command, args []string) {
	newFlag, _ := cmd.Flags().GetBool("new")
	if newFlag {
		createNewAccount(r)
		return
	}

	//check exit
	ok, accounts := getExistingAccounts(r)

	if !ok {
		createNewAccount(r)
	} else {
		index := getChosenAccount(accounts)

		fmt.Printf("current account: %v\n", accounts[index])

		psw, err := getPassword()
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("psw: %v\n", string(psw))

		entries, err := r.repo.List(accounts[index])
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Url | Login\n")
		for i, e := range entries {
			fmt.Printf("%v: %v - %v\n", i+1, e[1], e[0])
		}

		return
	}
}

func createNewAccount(r *Root) {
	fmt.Println("psw files not exist")
	fmt.Println("Create login")
	login, err := getInput()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	if err := r.repo.CreateFile(login); err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}
}

func getChosenAccount(accounts []string) int {
	if len(accounts) == 1 {
		return 0
	}

	fmt.Println("founded accounts:")
	for i, l := range accounts {
		fmt.Printf("%v: %v\n", i+1, l)
	}

	fmt.Print("Enter account number: ")
	chosenLogin, err := getInput()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	inputIndex, err := strconv.Atoi(chosenLogin)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	return inputIndex - 1
}

func getChosenLogin(logins [][]string) int {
	if len(logins) == 1 {
		return 0
	}

	fmt.Print("Enter login number: ")
	chosenLogin, err := getInput()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	inputIndex, err := strconv.Atoi(chosenLogin)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	return inputIndex - 1
}

func getExistingAccounts(r *Root) (bool, []string) {
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
