package commands

func (r *Root) remove(_ []string) error {
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

	if err := r.repo.Remove(account, entries[rowIndex].Url, entries[rowIndex].Login); err != nil {
		return err
	}

	return nil
}
