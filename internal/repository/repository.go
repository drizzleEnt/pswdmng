package repository

import "pswdmng/internal/domain"

type Repository interface {
	CheckExist() (bool, []domain.UserInfo, error)
	Add(account string, userInfo domain.UserInfo) error
	Get(account string, url string, login string) (string, error)
	List(login string) ([]domain.UserInfo, error)
	Remove(account string, url string, login string) error
	CreateFile(login string) error
}
