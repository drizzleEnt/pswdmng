package generator

import "pswdmng/internal/service"

var _ service.Generator = (*generator)(nil)

func New() *generator {
	return &generator{}
}

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars  = lowercase + uppercase + digits + symbols
)

type generator struct {
}

func (g *generator) GetPassword() string {
	//generate
	//hash

	return ""
}
