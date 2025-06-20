package commands

import (
	"fmt"
)

func (r *Root) get(_ []string) error {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		return err
	}

	entries, err := getLoginsAndUrls(r, account)
	if err != nil {
		return err
	}

	rowIndex, err := getChosenItem(entries)
	if err != nil {
		return err
	}

	pswd, err := r.repo.Get(account, entries[rowIndex].Url, entries[rowIndex].Login)
	if err != nil {
		return err
	}

	fmt.Printf("db pswd: %v\n", pswd)
	return nil
}
