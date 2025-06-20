package commands

func (r *Root) add(login, url string) error {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		return err
	}

	if err := r.repo.Add(account, login, url, "123"); err != nil {
		return err
	}

	return nil
}
