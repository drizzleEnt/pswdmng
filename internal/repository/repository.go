package repository

type Repository interface {
	CheckExist() (bool, []string, error)
	Add(account string, login string, url string) error
	Get()
	List(login string) ([]string, error)
	Remove()
	CreateFile(login string) error
}
