package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skulos/go-credentials/internal/commands"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a key-value pair from the credentials store",
	GroupID: "management",
	Args:    noArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the existing value
		exists, val, err := commands.PeekCredential(env, addKey)
		if err != nil {
			return err
		}
		if !exists {
			fmt.Printf("❌ Key '%s' does not exist in %s credentials.\n", addKey, env)
			return nil
		}

		// Confirm removal unless forced
		if !force {
			fmt.Printf("🗑️  Found '%s' with value: %v\n", addKey, val)
			fmt.Print("⚠️  Are you sure you want to delete it? [y/n]: ")
			reader := bufio.NewReader(os.Stdin)
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSpace(strings.ToLower(confirm))
			if confirm != "y" && confirm != "yes" {
				fmt.Println("❌ Delete cancelled.")
				return nil
			}
		}

		return commands.RemoveCredential(env, addKey)
	},
}

func init() {
	removeCmd.Flags().StringVarP(&addKey, "key", "k", "", "Credential key to remove (required)")
	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal without confirmation")

	_ = removeCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(removeCmd)
}
