package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetPreviousVersion(path string) ([]byte, error) {
	cmd := exec.Command("git", "show", fmt.Sprintf("HEAD:%s", path))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("git show error: %s", stderr.String())
	}
	return out.Bytes(), nil
}
