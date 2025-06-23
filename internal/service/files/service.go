package files

import (
	"pswdmng/internal/service"
)

var _ service.FileService = (*fileService)(nil)

type fileService struct {
}

func (s *fileService) ReadEncryptedFile(path string, key []byte) ([]byte, error) {
	return nil, nil
}

func (s *fileService) WriteEncryptedFile(path string, data, key []byte) error {
	return nil
}
