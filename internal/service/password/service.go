package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"math/big"
	"pswdmng/internal/service"
)

var _ service.PasswordService = (*passwordService)(nil)

func New() *passwordService {
	return &passwordService{}
}

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars  = lowercase + uppercase + digits + symbols
)

type passwordService struct {
}

func (g *passwordService) GetNewPassword(length int) (string, error) {
	pswd, err := g.generatePassword(length)
	if err != nil {
		return "", err
	}

	encryptedPass, err := g.EncryptPassword(nil, pswd)
	if err != nil {
		return "", err
	}

	return string(encryptedPass), nil
}

func (g *passwordService) generatePassword(length int) ([]byte, error) {
	if length < 12 {
		length = 12
	}

	result := make([]byte, 0)

	result = append(result, pickChar(lowercase))
	result = append(result, pickChar(uppercase))
	result = append(result, pickChar(digits))
	result = append(result, pickChar(symbols))

	for len(result) < length {
		c := pickChar(allChars)
		result = append(result, c)
	}

	shuffle(result)

	return result, nil
}

func (g *passwordService) EncryptPassword(key []byte, plainpass []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, plainpass, nil)
	return ciphertext, nil
}

func (g *passwordService) DecryptPassword(key, cipherPass []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherPass) < nonceSize {
		return nil, fmt.Errorf("error")
	}

	nonce, cipherText := cipherPass[:nonceSize], cipherPass[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func pickChar(chars string) byte {
	max := big.NewInt(int64(len(chars)))
	n, _ := rand.Int(rand.Reader, max)
	return chars[n.Int64()]
}

func shuffle(original []byte) {
	for i := range original {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(len(original))))
		original[i], original[j.Int64()] = original[j.Int64()], original[i]
	}
}
