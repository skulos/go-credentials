package commands

import (
	"fmt"
	"os"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/skulos/go-credentials/internal/crypto"
	"github.com/skulos/go-credentials/internal/environment"
	"github.com/skulos/go-credentials/internal/git"
)

func ShowGitDifference(env string) (string, error) {
	keyName := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)
	keyPath := fmt.Sprintf(".credentials/%s.key", keyName)
	encPath := fmt.Sprintf(".credentials/%s.yml.enc", encName)

	// oldEncrypted, err := getPreviousVersion(".credentials/credentials.yml.enc")

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read private key: %w", err)
	}

	identity, err := crypto.ParseIdentity(string(keyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to parse identity: %w", err)
	}

	oldEncrypted, err := git.GetPreviousVersion(encPath)

	if err != nil {
		return "", fmt.Errorf("⚠️  No previous version of %s found in git", encPath)
	}

	oldDecrypted, err := crypto.DecryptFromBytes(oldEncrypted, identity)
	if err != nil {
		return "", fmt.Errorf("⚠️  Failed to decrypt previous version: %w", err)
	}

	// Load current version
	currentEncrypted, err := os.ReadFile(encPath)
	if err != nil {
		return "", fmt.Errorf("failed to read current encrypted file: %w", err)
	}

	currentDecrypted, err := crypto.DecryptFromBytes(currentEncrypted, identity)
	if err != nil {
		return "", fmt.Errorf("⚠️  Failed to decrypt current version: %w", err)
	}

	// Diff the decrypted files
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(oldDecrypted), string(currentDecrypted), false)

	return dmp.DiffPrettyText(diffs), nil
}
