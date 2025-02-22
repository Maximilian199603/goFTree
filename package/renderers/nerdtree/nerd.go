package nerdtree

import (
	"path/filepath"
	"strings"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/generaltree"
)

const (
	// directoryIcon is the Nerd Font icon for directories.
	directoryIcon = " "
	// genericFileIcon is used when no specific file icon is found.
	genericFileIcon = " "
)

type Engine struct{}

// Render generates the Nerd Font file tree as a string.
func (e Engine) Render(tree *filetree.FileTree) (string, error) {

	set := &generaltree.RenderSettings{
		DirTJunction:  "├── ",
		DirLJunction:  "└── ",
		FileTJunction: "├── ",
		FileLJunction: "└── ",
		NoJunction:    "│   ",
		Empty:         "    ",

		DirPrepender:  func(s string, n *filetree.Node) (string, *filetree.Node) { return getIcon(n) + s, n },
		DirAppender:   func(s string, n *filetree.Node) (string, *filetree.Node) { return s + "/", n },
		FilePrepender: func(s string, n *filetree.Node) (string, *filetree.Node) { return getIcon(n) + s, n },
		FileAppender:  generaltree.Noop,
	}
	return generaltree.Render(tree, set)
}

// getIcon returns the appropriate Nerd Font icon for the given node.
func getIcon(node *filetree.Node) string {
	if node.IsDir {
		return directoryIcon
	}
	nameAsLower := strings.ToLower(node.Name)
	icon := fileNameToIcon[nameAsLower]
	if icon != "" {
		return icon
	}
	suffix := filepath.Ext(nameAsLower)
	icon = fileExtensionToIcon[suffix]
	if icon != "" {
		return icon
	}
	return genericFileIcon
}
