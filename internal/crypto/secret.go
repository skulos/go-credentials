package crypto

import (
	"crypto/rand"
	"fmt"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateSecret(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	for i := range b {
		b[i] = letters[int(b[i])%len(letters)]
	}
	return string(b)
}

func DefaultYaml() []byte {
	defaultYAML := fmt.Sprintf(`---
secret_key_base: %s
`, generateSecret(64))

	return []byte(defaultYAML)
}
