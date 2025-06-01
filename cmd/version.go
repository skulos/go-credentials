package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  NoArgs,
	Short: "Display the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("credentials version", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
