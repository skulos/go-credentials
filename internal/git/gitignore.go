package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/skulos/go-credentials/internal/colours"
)

// AddIgnoreLine ensures the given path is listed in .gitignore
func AddIgnoreLine() (bool, error) {
	const gitignorePath = ".gitignore"
	var line = ".credentials/*.key"
	line = strings.TrimSpace(line)

	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		// return os.WriteFile(gitignorePath, []byte(line+"\n"), 0644)
		return true, nil
	}

	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		return false, err
	}

	if !containsLine(string(content), line) {
		f, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return false, err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Printf("failed to close .gitignore file: %s", colours.WarnColor(err))
			}
		}(f)
		_, err = f.WriteString(line + "\n")
		return false, err
	}
	return false, nil
}

func containsLine(content, line string) bool {
	for _, l := range strings.Split(content, "\n") {
		if strings.TrimSpace(l) == line {
			return true
		}
	}
	return false
}
