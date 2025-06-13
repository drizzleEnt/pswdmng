package repository

type Repository interface {
	CheckExist() (bool, error)
	Add()
	Get()
	List()
	Remove()
}