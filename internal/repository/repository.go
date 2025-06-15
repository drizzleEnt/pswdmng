package repository

type Repository interface {
	CheckExist() (bool, []string, error)
	Add(account string, login string, url string) error
	Get(account string, url string, login string) (string, error)
	List(login string) ([][]string, error)
	Remove()
	CreateFile(login string) error
}
