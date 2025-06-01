package crypto

import (
	"os"
)

// FileExists returns true if the file at path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// CreateDirIfNotExists ensures a directory exists.
func CreateDirIfNotExists(path string) error {
	return os.MkdirAll(path, 0700)
}

// WriteFileSecure writes data to a file with secure permissions.
func WriteFileSecure(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}
