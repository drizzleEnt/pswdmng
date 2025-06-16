package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Root) remove(cmd *cobra.Command, args []string) {
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

	rowIndex, err := getChosenItem(entries)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	if err := r.repo.Remove(account, entries[rowIndex].Url, entries[rowIndex].Login); err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
}
