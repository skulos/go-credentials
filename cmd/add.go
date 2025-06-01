package cmd

import (
	"github.com/skulos/go-credentials/internal/commands"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add -k [key] -v [value]",
	Short:   "Add a key-value pair into the encrypted credentials store",
	GroupID: "management",
	Args:    NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commands.AddCredential(env, addKey, addValue, force)
	},
}

func init() {
	addCmd.Flags().StringVarP(&addKey, "key", "k", "", "Key to add to the credentials store")
	addCmd.Flags().StringVarP(&addValue, "value", "v", "", "Value for the key being added")
	addCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite of existing key")

	addCmd.MarkFlagRequired("key")
	addCmd.MarkFlagRequired("value")

	rootCmd.AddCommand(addCmd)
}
