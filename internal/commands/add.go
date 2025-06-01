package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/skulos/go-credentials/internal/crypto"
	"github.com/skulos/go-credentials/internal/environment"
	"gopkg.in/yaml.v3"
)

func AddCredential(env, key, value string, force bool) error {
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

	/////////
	segments := strings.Split(strings.ToLower(key), ".")
	last := segments[len(segments)-1]
	current := data

	for _, seg := range segments[:len(segments)-1] {
		if next, ok := current[seg]; ok {
			if m, ok := next.(map[string]interface{}); ok {
				current = m
			} else {
				return fmt.Errorf("'%s' exists and is not a nested object", seg)
			}
		} else {
			newMap := make(map[string]interface{})
			current[seg] = newMap
			current = newMap
		}
	}

	// Check for collision
	if _, exists := current[last]; exists && !force {
		return fmt.Errorf("key '%s' already exists — use '--force' to overwrite it", key)
	}

	current[last] = value

	// Marshal and encrypt
	newYAML, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}
	recipient := identity.Recipient()
	if err := crypto.EncryptToFile(encPath, recipient, newYAML); err != nil {
		return fmt.Errorf("failed to encrypt credentials: %w", err)
	}

	fmt.Printf("🔐 Successfully added '%s' to '%s'\n", key, encPath)
	return nil

	///////

	// lowercaseKey := strings.ToLower(key)

	// // Check if key already exists
	// if _, exists := data[lowercaseKey]; exists && !force {
	// 	return fmt.Errorf("key '%s' already exists — use '--force' to overwrite it or use the 'update' and 'edit' commands instead", key)
	// }

	// data[lowercaseKey] = value

	// // Marshal and encrypt the new YAML
	// newYAML, err := yaml.Marshal(data)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal credentials: %w", err)
	// }
	// recipient := identity.Recipient()
	// if err := crypto.EncryptToFile(encPath, recipient, newYAML); err != nil {
	// 	return fmt.Errorf("failed to encrypt credentials: %w", err)
	// }

	// fmt.Printf("🔐 Successfully added '%s' to '%s'\n", lowercaseKey, encPath)
	// return nil
}
