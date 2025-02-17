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

	path := ""
	if len(args) == 0 {
		p, err := os.Getwd()
		if err != nil {
			return err
		}
		path = p
	} else {
		path = args[0]
	}

	hOptions := parseHiddenOptions()

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

	tree, err := filetree.BuildTree(path, hOptions)
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

func parseHiddenOptions() *filetree.HiddenOptions {
	dec := strings.ToLower(includes)
	switch dec {
	case "all":
		return filetree.NewHiddenOption(true, true, true)
	case "dirs":
		return filetree.NewHiddenOption(false, true, false)
	case "files":
		return filetree.NewHiddenOption(false, false, true)
	default:
		return filetree.NewHiddenOption(false, false, false)
	}
}
