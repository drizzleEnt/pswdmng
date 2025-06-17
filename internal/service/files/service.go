package files

import (
	"pswdmng/internal/service"
)

var _ service.FileService = (*checkerService)(nil)

type checkerService struct {
}

func (s *checkerService) ReadEncryptedFile(path string, key []byte) ([]byte, error) {
	return nil, nil
}

func (s *checkerService) WriteEncryptedFile(path string, data, key []byte) error {
	return nil
}
