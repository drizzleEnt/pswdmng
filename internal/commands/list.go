package commands

func (r *Root) list(_ []string) error {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		return err
	}

	if _, err := getLoginsAndUrls(r, account); err != nil {
		return err
	}

	return nil
}
