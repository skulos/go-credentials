package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Hardcode the LICENSE text as a string
const licenseText = "MIT License © 2025 Hendre Hayman <hendrehayman@gmail.com>"

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Display the software license",
	Args:  NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(licenseText)
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}
