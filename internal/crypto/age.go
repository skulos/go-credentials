package crypto

import (
	"bytes"
	"io"
	"os"
	"strings"

	"filippo.io/age"
)

// GenerateIdentity generates a new X25519 identity.
func GenerateIdentity() (*age.X25519Identity, error) {
	return age.GenerateX25519Identity()
}

// EncryptToFile encrypts the plaintext with the recipient and writes it to the file
func EncryptToFile(path string, recipient age.Recipient, data []byte) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	encWriter, err := age.Encrypt(out, recipient)
	if err != nil {
		return err
	}

	_, err = bytes.NewReader(data).WriteTo(encWriter)
	if err != nil {
		return err
	}
	return encWriter.Close()
}

func DecryptFile(path string, identity *age.X25519Identity) ([]byte, error) {
	in, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	r, err := age.Decrypt(in, identity)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}

func ParseIdentity(s string) (*age.X25519Identity, error) {
	id, err := age.ParseX25519Identity(strings.TrimSpace(s))
	if err != nil {
		return nil, err
	}
	return id, nil
}

func DecryptFromFile(path string, identity age.Identity) ([]byte, error) {
	in, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	r, err := age.Decrypt(in, identity)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}

func DecryptFromBytes(data []byte, identity age.Identity) ([]byte, error) {
	r, err := age.Decrypt(bytes.NewReader(data), identity)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}
