package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/skulos/go-credentials/internal/crypto"
	"github.com/skulos/go-credentials/internal/environment"
	"gopkg.in/yaml.v3"
)

func RemoveCredential(env, key string) error {
	keyName := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)

	keyPath := filepath.Join(".credentials", fmt.Sprintf("%s.key", keyName))
	encPath := filepath.Join(".credentials", fmt.Sprintf("%s.yml.enc", encName))

	// Read private key
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}
	identity, err := crypto.ParseIdentity(string(keyBytes))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	// Decrypt
	plaintext, err := crypto.DecryptFromFile(encPath, identity)
	if err != nil {
		return fmt.Errorf("failed to decrypt credentials: %w", err)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(plaintext, &data); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	lowerKey := strings.ToLower(key)
	if _, exists := data[lowerKey]; !exists {
		return fmt.Errorf("key '%s' not found", lowerKey)
	}

	delete(data, lowerKey)

	// Marshal and re-encrypt
	updatedYAML, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to re-marshal YAML: %w", err)
	}

	recipient := identity.Recipient()
	if err := crypto.EncryptToFile(encPath, recipient, updatedYAML); err != nil {
		return fmt.Errorf("failed to re-encrypt updated credentials: %w", err)
	}

	fmt.Printf("✅ Successfully removed '%s' from '%s'\n", lowerKey, encPath)
	return nil
}
