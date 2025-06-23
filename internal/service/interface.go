package service

type FileService interface {
}

type PasswordService interface {
	GetNewPassword(length int) (string, error)
}
