package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EdgeLordKirito/goFTree/internal/goFTree/version"
	"github.com/EdgeLordKirito/goFTree/internal/renderswitch"
	"github.com/EdgeLordKirito/goFTree/package/filetree"
	"github.com/spf13/cobra"
)

func runTreeCommand(cmd *cobra.Command, args []string) error {
	if printVersion {
		fmt.Println(version.Version)
		return nil
	}

	if len(args) == 0 {
		cmd.Help()
		return nil
	}

	path := args[0]

	toStdout := true

	// Check if outputFile is set
	if outputFile != "" {
		// Check if the file already exists
		if _, err := os.Stat(outputFile); err == nil {
			fmt.Printf("File '%s' already exists. Overwrite? (y/N): ", outputFile)

			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("Error reading input: %v", err)
			}

			response = strings.TrimSpace(strings.ToLower(response))
			if response != "y" && response != "yes" {
				fmt.Println("Operation canceled.")
				return nil
			}
		}
		toStdout = false
	}

	engine, err := renderswitch.GetRenderEngine(style)
	if err != nil {
		return err
	}

	tree, err := filetree.BuildTree(path)
	if err != nil {
		return err
	}
	out, err := engine.Render(tree)
	if err != nil {
		return fmt.Errorf("Error Rendering Tree '%v'", err)
	}

	if toStdout {
		fmt.Print(out)
	} else {
		err = os.WriteFile(outputFile, []byte(out), 0644)
		if err != nil {
			return fmt.Errorf("Error writing to file '%s': %v", outputFile, err)
		}
		fmt.Printf("Output written to %s\n", outputFile)
	}
	return nil
}
