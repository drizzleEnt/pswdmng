package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func (r *Root) init(cmd *cobra.Command, args []string) {
	//check exit
	fmt.Println("Check existing passwords files")

	ok, logins, err := r.repo.CheckExist()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	if !ok {
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
	} else {
		fmt.Println("founded accounts:")
		for i, l := range logins {
			fmt.Printf("%v: %v\n", i+1, l)
		}

		fmt.Println("enter number:")
		chosenLogin, err := getInput()
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}

		index, err := strconv.Atoi(chosenLogin)
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("current account: %v\n", logins[index-1])

		// psw, err := getPassword()
		// if err != nil {
		// 	fmt.Printf("err.Error(): %v\n", err.Error())
		// 	os.Exit(1)
		// }
		entries, err := r.repo.List(logins[index-1])
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			os.Exit(1)
		}

		for i, e := range entries {
			fmt.Printf("%v: %v\n", i+1, e)
		}
		return
	}

	fmt.Println("Passwords files founded")
	fmt.Println("Enter master password for access to account:")
	fmt.Println("init cmd")
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
