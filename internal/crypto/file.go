package crypto

import (
	"github.com/spf13/afero"
)

// FileExists returns true if the file at path exists.
func FileExists(path string, filesystem afero.Fs) bool {
	_, err := filesystem.Stat(path)
	return err == nil
}

// CreateDirIfNotExists ensures a directory exists.
func CreateDirIfNotExists(path string, filesystem afero.Fs) error {
	return filesystem.MkdirAll(path, 0700)
}

// WriteFileSecure writes data to a file with secure permissions.
func WriteFileSecure(path string, filesystem afero.Fs, data []byte) error {
	return afero.WriteFile(filesystem, path, data, 0600)
}
