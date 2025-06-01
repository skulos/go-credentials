package commands

import (
	"fmt"

	"github.com/skulos/go-credentials/internal/crypto"
	"github.com/skulos/go-credentials/internal/environment"
	"github.com/skulos/go-credentials/internal/git"
)

func SetupEnvironment(env string) (string, string, bool, error) {

	// if err := gitignore.AddIgnoreLine(keyPath); err != nil {
	// 	fmt.Println("⚠️  Could not update .gitignore for key:", err)
	// }
	// if err := gitignore.AddIgnoreLine(encPath); err != nil {
	// 	fmt.Println("⚠️  Could not update .gitignore for yaml:", err)
	// }

	const keyDir = ".credentials"
	key := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)
	keyName := fmt.Sprintf("%s.key", key)
	keyPath := fmt.Sprintf("%s/%s", keyDir, keyName)
	encPath := fmt.Sprintf("%s/%s.yml.enc", keyDir, encName)

	if crypto.FileExists(keyPath) {
		return keyPath, encPath, false, fmt.Errorf("🔒 Master key already exists at %s", keyPath)
	}

	identity, err := crypto.GenerateIdentity()
	if err != nil {
		return keyPath, encPath, false, err
	}

	if err := crypto.CreateDirIfNotExists(keyDir); err != nil {
		return keyPath, encPath, false, err
	}

	if err := crypto.WriteFileSecure(keyPath, []byte(identity.String())); err != nil {
		return keyPath, encPath, false, err
	}

	recipient := identity.Recipient()

	if err := crypto.EncryptToFile(encPath, recipient, crypto.DefaultYaml()); err != nil {
		return keyPath, encPath, false, fmt.Errorf("failed to create encrypted YAML: %w", err)
	}

	fileExists, err := git.AddIgnoreLine()

	if err != nil {
		fmt.Println("⚠️  Could not update .gitignore for key:", err)
	}

	return keyPath, encPath, fileExists, nil
}
