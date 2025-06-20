package commands

import (
	"fmt"
	"os"
)

func (r *Root) add(login, url string) {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	if err := r.repo.Add(account, login, url, "123"); err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
