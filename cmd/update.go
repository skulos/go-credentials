package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skulos/go-credentials/internal/commands"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing key with a new value in the credentials store",
	GroupID: "management",
	Args:    noArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch existing value first
		exists, oldValue, err := commands.PeekCredential(env, addKey)
		if err != nil {
			return err
		}
		if !exists {
			fmt.Printf("❌ Key '%s' does not exist in %s credentials.\n", addKey, env)
			return nil
		}
		if exists && !force {
			fmt.Printf("🔁 Key '%s' already exists with value: %v\n", addKey, oldValue)
			fmt.Print("⚠️  Do you want to update it? [y/N]: ")
			reader := bufio.NewReader(os.Stdin)
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSpace(strings.ToLower(confirm))
			if confirm != "y" && confirm != "yes" {
				fmt.Println("❌ Update cancelled.")
				return nil
			}
		}

		return commands.UpdateCredentials(env, addKey, addValue, filesystem)
	},
}

func init() {
	updateCmd.Flags().StringVar(&env, "env", "development", "Environment name")
	updateCmd.Flags().StringVar(&addKey, "key", "", "Key to update")
	updateCmd.Flags().StringVar(&addValue, "value", "", "New value")
	updateCmd.Flags().BoolVar(&force, "force", false, "Force update without confirmation")

	_ = updateCmd.MarkFlagRequired("key")
	_ = updateCmd.MarkFlagRequired("value")

	rootCmd.AddCommand(updateCmd)
}
