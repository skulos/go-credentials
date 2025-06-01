package cmd

import (
	"fmt"

	"github.com/skulos/go-credentials/internal/colours"
	"github.com/skulos/go-credentials/internal/environment"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print environment variables and their uses",
	Args:  NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printEnvVars()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func printEnvVars() {

	fmt.Print("Environment Variables and Their Uses:\n\n")

	fmt.Print(colours.KeyColor(environment.CREDENTIALS_ENV + ": "))
	fmt.Print(colours.ValueColor("This is the environment used for selecting the appropriate credentials file (e.g., 'development', 'staging', 'production').\n\n"))

	fmt.Print(colours.KeyColor(environment.MASTER_KEY + ": "))
	fmt.Print(colours.ValueColor("This is the path to the 'age' private key used for encryption/decryption. If set, it overrides the default .credentials/*.key file.\n\n"))

}
