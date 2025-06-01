package cmd

import (
	"fmt"

	"github.com/skulos/go-credentials/internal/commands"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit the credentials store for an environment",
	GroupID: "core",
	Args:    NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		encryptedFile, edited, err := commands.Editor(env)
		if err != nil {
			return fmt.Errorf("failed to edit credentials: %w", err)
		}
		if edited {
			fmt.Printf("🔐 Successfully saved updates to '%s'\n", encryptedFile)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
