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

func UpdateCredentials(env, key, value string) error {

	keyName := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)

	keyPath := fmt.Sprintf(".credentials/%s.key", keyName)
	encPath := fmt.Sprintf(".credentials/%s.yml.enc", encName)

	// Read the private key
	identityBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}
	identity, err := crypto.ParseIdentity(string(identityBytes))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	// Decrypt the YAML file
	var data map[string]interface{}
	plaintext, err := crypto.DecryptFile(encPath, identity)
	if err != nil {
		// If file doesn't exist or decryption fails, start with empty data
		fmt.Println("⚠️  No existing credentials found; creating a new encrypted file.")
		data = make(map[string]interface{})
	} else {
		if len(plaintext) > 0 {
			if err := yaml.Unmarshal(plaintext, &data); err != nil {
				return fmt.Errorf("failed to parse existing credentials: %w", err)
			}
		}
		// Ensure map is initialized even if unmarshal succeeded with nil
		if data == nil {
			data = make(map[string]interface{})
		}
	}

	lowerKey := strings.ToLower(key)
	data[lowerKey] = value

	updated, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal updated credentials: %w", err)
	}

	if err := crypto.EncryptToFile(encPath, identity.Recipient(), updated); err != nil {
		return fmt.Errorf("failed to encrypt updated credentials: %w", err)
	}

	fmt.Printf("🔄 Successfully updated '%s' in '%s'\n", lowerKey, encPath)
	return nil
}

// PeekCredential checks if a key exists and returns its value if it does.
func PeekCredential(env, key string) (exists bool, value interface{}, err error) {
	keyName := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)

	keyPath := filepath.Join(".credentials", fmt.Sprintf("%s.key", keyName))
	encPath := filepath.Join(".credentials", fmt.Sprintf("%s.yml.enc", encName))

	// Load private key
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return false, nil, fmt.Errorf("failed to read private key: %w", err)
	}

	identity, err := crypto.ParseIdentity(string(keyBytes))
	if err != nil {
		return false, nil, fmt.Errorf("failed to parse identity: %w", err)
	}

	// Decrypt file
	plaintext, err := crypto.DecryptFromFile(encPath, identity)
	if err != nil || len(plaintext) == 0 {
		return false, nil, nil // no data or unable to decrypt, treated as "not found"
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(plaintext, &data); err != nil {
		return false, nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	val, ok := data[strings.ToLower(key)]
	return ok, val, nil
}
