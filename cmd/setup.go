package cmd

import (
	"fmt"

	"github.com/skulos/go-credentials/internal/commands"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:     "setup",
	Short:   "Initialize a master key and credentials store for an environment",
	GroupID: "core",
	Args:    cobra.NoArgs,
	Aliases: []string{"init"},
	RunE: func(cmd *cobra.Command, args []string) error {

		keyPath, encPath, fileExists, err := commands.SetupEnvironment(env)

		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			fmt.Println("✅ Credential key generated and stored.")
			fmt.Println("🗝️  Key stored at:", keyPath)
			fmt.Println("📝 Encrypted YAML created at:", encPath)

			if fileExists {
				fmt.Println("📝 No .gitignore file")
			} else {
				fmt.Println("🔥 Added", keyPath, "to .gitignore")
			}
			fmt.Println("📢 Encryption and decryption will use this key.")
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
