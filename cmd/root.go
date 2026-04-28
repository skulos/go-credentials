package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	// Global Flags
	env      string
	noColour bool

	// Add
	force    bool
	addKey   string
	addValue string

	// Show
	key string

	// Filesystem
	filesystem afero.Fs

	// Bootstrap Sync
	boot sync.Once
)

var noArgs = func(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		argList := strings.Join(args, " ")
		// _ = cmd.Usage()
		return fmt.Errorf(
			"\n❌ No argument(s) '%s' are expected for command '%s'.\n\nℹ️  Use '%s --help' for more information",
			argList, cmd.CommandPath(), cmd.CommandPath(),
		)
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:           "credentials",
	Short:         "Manage encrypted credentials",
	Long:          "A lightweight, Rails-inspired encrypted credentials system for Go.",
	Version:       version,
	SilenceErrors: true,
}

func bootstrap() *cobra.Command {
	boot.Do(func() {
		filesystem = afero.NewOsFs()

		rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "development", "environment name (e.g. development, production, staging, test, ...)")
		rootCmd.PersistentFlags().BoolVarP(&noColour, "no-color", "c", false, "disable colourized output")

		rootCmd.AddGroup(&cobra.Group{
			ID:    "management",
			Title: "Management Commands:",
		})

		rootCmd.AddGroup(&cobra.Group{
			ID:    "core",
			Title: "Core Commands:",
		})
	})

	return rootCmd
}

func GetEnv() string {
	return env
}

func main() {
	bootstrap()
	//generate()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//func generate() {
//	outputDir := "./docs"
//	err := os.MkdirAll(outputDir, os.ModePerm)
//	if err != nil {
//		fmt.Printf("❌ Failed to create docs directory: %v\n", err)
//		os.Exit(1)
//	}
//
//	err = doc.GenMarkdownTree(cmd.RootCmd, outputDir) // Replace 'cmd.RootCmd' with your actual root command
//	if err != nil {
//		fmt.Printf("❌ Failed to generate docs: %v\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Printf("✅ Documentation generated at %s\n", outputDir)
//}
