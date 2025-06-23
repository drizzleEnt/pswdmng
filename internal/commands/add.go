package commands

import "pswdmng/internal/domain"

func (r *Root) add(login, url string) error {
	account, _, err := getAccountAndPassword(r)
	if err != nil {
		return err
	}

	userInfo := domain.UserInfo{
		Login: login,
		Url:   url,
	}

	userInfo.Password, err = r.pswSrv.GetNewPassword(10)
	if err != nil {
		return err
	}

	if err := r.repo.Add(account, userInfo); err != nil {
		return err
	}

	return nil
}
