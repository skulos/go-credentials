package environment

import "os"

// ResolveEnv returns the selected environment name (e.g., "staging", "production").
// Priority: CLI flag > ENV var > default ("master").
func ResolveEnv(cliFlag string, keyFile bool) string {
	if fromEnv := os.Getenv(CREDENTIALS_ENV); fromEnv != "" && fromEnv != DEFAULT_ENVIRONMENT {
		return fromEnv
	}
	if cliFlag != "" && cliFlag != DEFAULT_ENVIRONMENT {
		return cliFlag
	}

	if keyFile {
		return DEFUALT_KEY_NAME
	} else {
		return DEFAULT_ENCRYPTED_FILE
	}
}
