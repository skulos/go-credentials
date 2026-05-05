package cmd

import (
	"fmt"

	"github.com/skulos/go-credentials/internal/commands"
	"github.com/skulos/go-credentials/internal/environment"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show the credentials store for an environment",
	GroupID: "core",
	Args:    noArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var colourOuput bool

		keyName := environment.ResolveEnv(env, true)
		encName := environment.ResolveEnv(env, false)

		// Check if --diff flag is passed
		diffFlag, _ := cmd.Flags().GetBool("diff")
		rawFlag, _ := cmd.Flags().GetBool("raw")

		if diffFlag {
			// Show the diff
			diff, err := commands.ShowGitDifference(env)
			if err != nil {
				return err
			}
			fmt.Printf("🔓 Decrypted credentials diff for %s (%s):\n\n", env, encName)
			fmt.Println(diff)
			return nil
		}

		if rawFlag {
			colourOuput = false
		} else {
			colourOuput = !noColour
		}

		// If --diff is not passed, show the decrypted credentials as normal
		contents, encryptedFile, err := commands.ShowCredentials(keyName, encName, key, colourOuput, rawFlag)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !rawFlag {
			fmt.Printf("🔓 Decrypted credentials for %s (%s):\n\n", env, encryptedFile)
		}

		fmt.Println(contents)
		return nil
	},
}

func init() {
	showCmd.Flags().StringVarP(&key, "key", "k", "", "Show value for specific key")
	showCmd.Flags().BoolP("diff", "d", false, "Show the diff between current and previous encrypted credentials")
	showCmd.Flags().BoolP("raw", "", false, "Shows only the raw value for the key")

	rootCmd.AddCommand(showCmd)
}
