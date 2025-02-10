package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	style      string
	outputFile string
)

// Execute runs the CLI tool
func Execute() {
	rootCmd := buildRoot()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func buildRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "goFTree [path]",
		Short: "goFTree - A CLI tool to display filesystem structure",
		Args:  cobra.ExactArgs(1), // Enforce exactly one argument (path)
		RunE:  runTreeCommand,
	}

	rootCmd.PersistentFlags().StringVarP(&style, "style", "s", "ascii",
		"Choose output style: ascii, markdown, xml, json, line, dashed")

	rootCmd.PersistentFlags().StringVarP(&outputFile, "file", "f", "",
		"Write output to a file instead of stdout")

	return rootCmd
}
