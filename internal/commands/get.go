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

	entries, err := getLoginsAndUrls(r, account)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	rowIndex := getChosenLogin(entries)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	pswd, err := r.repo.Get(account, entries[rowIndex].Url, entries[rowIndex].Login)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("db pswd: %v\n", pswd)
	return
}
