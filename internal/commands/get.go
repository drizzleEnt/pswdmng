package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) get(cmd *cobra.Command, args []string) {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	entries, err := r.repo.List(account)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Printf("You did not add any login\n")
		return
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

	pswd, err := r.repo.Get(account, entries[rowIndex][1], entries[rowIndex][0])
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("db pswd: %v\n", pswd)
	return
}
