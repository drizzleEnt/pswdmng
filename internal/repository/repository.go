package repository

type Repository interface {
	CheckExist() (bool, []string, error)
	Add()
	Get()
	List(login string) ([]string, error)
	Remove()
	CreateFile(login string) error
}
