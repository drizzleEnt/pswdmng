package password

import (
	"crypto/rand"
	"testing"

	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	s := &passwordService{}
	key := make([]byte, 32)
	rand.Read(key)

	original := "test_password1"

	cipher, err := s.EncryptPassword(key, []byte(original))
	require.NoError(t, err)

	decrypted, err := s.DecryptPassword(key, cipher)
	require.NoError(t, err)

	assert.Equal(t, original, string(decrypted))
}
