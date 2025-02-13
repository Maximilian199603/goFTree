package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	style        string
	outputFile   string
	printVersion bool
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
		Args:  validateArgs,
		RunE:  runTreeCommand,
	}

	rootCmd.PersistentFlags().StringVarP(&style, "style", "s", "ascii",
		"Choose output style: ascii, markdown, xml, json, line, dashed, nerd")

	rootCmd.PersistentFlags().StringVarP(&outputFile, "file", "f", "",
		"Write output to a file instead of stdout")

	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false,
		"Print Current Version")

	return rootCmd
}

func validateArgs(cmd *cobra.Command, args []string) error {
	_ = cmd

	if len(args) == 1 && printVersion {
		return fmt.Errorf("cannot provide an argument and use --version (-v)")
	}

	if len(args) > 1 {
		return fmt.Errorf("too many arguments: expected 0 or 1")
	}

	return nil
}
