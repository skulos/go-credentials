package main

import (
	"github.com/skulos/go-credentials/cmd"
	// "github.com/spf13/cobra/doc"
)

func main() {
	// generate()

	cmd.Execute()
}

// func generate() {
// 	outputDir := "./docs"
// 	err := os.MkdirAll(outputDir, os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("❌ Failed to create docs directory: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// err = doc.GenMarkdownTree(cmd.RootCmd, outputDir) // Replace 'cmd.RootCmd' with your actual root command
// 	// if err != nil {
// 	// 	fmt.Printf("❌ Failed to generate docs: %v\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// fmt.Printf("✅ Documentation generated at %s\n", outputDir)
// }
