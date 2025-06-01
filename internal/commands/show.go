package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/skulos/go-credentials/internal/colours"
	"github.com/skulos/go-credentials/internal/crypto"
	"gopkg.in/yaml.v3"
)

func ShowCredentials(keyName, encName, specificKey string, colouredOutput bool) (string, string, error) {
	keyPath := filepath.Join(".credentials", fmt.Sprintf("%s.key", keyName))
	encPath := filepath.Join(".credentials", fmt.Sprintf("%s.yml.enc", encName))

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return "", encPath, fmt.Errorf("Failed to read private key: %s", colours.ErrorColor(err))
	}

	identity, err := crypto.ParseIdentity(string(keyBytes))
	if err != nil {
		return "", encPath, fmt.Errorf("Failed to parse identity: %s", colours.ErrorColor(err))
	}

	plaintext, err := crypto.DecryptFromFile(encPath, identity)
	if err != nil {
		return "", encPath, fmt.Errorf("Failed to decrypt credentials: %s", colours.ErrorColor(err))
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(plaintext, &data); err != nil {
		return "", encPath, fmt.Errorf("Failed to parse YAML: %s", colours.ErrorColor(err))
	}

	// Show only the specific key if requested
	if specificKey != "" {
		value, found := traverseKey(data, specificKey)
		if !found {
			if colouredOutput {
				return "", encPath, fmt.Errorf("Key %s not found in credentials", colours.ErrorColor(specificKey))
			}
			return "", encPath, fmt.Errorf("Key %s not found in credentials", specificKey)
		}
		if colouredOutput {
			return fmt.Sprintf("%s : %v\n", colours.KeyColor(specificKey), colours.ValueColor(value)), encPath, nil
		}
		return fmt.Sprintf("%s : %v\n", specificKey, value), encPath, nil
	}

	// Pretty print top-level keys (flat output)
	keys := make([]string, 0, len(data))
	maxKeyLen := 0
	for k := range data {
		keys = append(keys, k)
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}
	sort.Strings(keys)

	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return "", encPath, fmt.Errorf("Failed to format YAML: %w", err)
	}

	if colouredOutput {
		lines := strings.Split(string(yamlBytes), "\n")
		var b strings.Builder
		for _, line := range lines {
			if strings.Contains(line, ":") {
				parts := strings.SplitN(line, ":", 2)
				keyPart := colours.KeyColor(parts[0])
				valPart := colours.ValueColor(parts[1])
				b.WriteString(fmt.Sprintf("%s:%s\n", keyPart, valPart))
			} else {
				b.WriteString(line + "\n")
			}
		}
		return b.String(), encPath, nil
	}

	return string(yamlBytes), encPath, nil
}

func traverseKey(data map[string]interface{}, dottedKey string) (interface{}, bool) {
	parts := strings.Split(dottedKey, ".")
	var current interface{} = data
	for _, part := range parts {
		if m, ok := current.(map[string]interface{}); ok {
			if val, exists := m[part]; exists {
				current = val
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
	return current, true
}
