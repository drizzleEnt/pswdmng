package commands

import (
	"fmt"
	"os"
)

func (r *Root) list(_ []string) {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	if _, err := getLoginsAndUrls(r, account); err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
}
