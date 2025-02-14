package filetree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrEmptyTree func() error = func() error { return errEmptyTree }
	errEmptyTree error        = errors.New("Empty File Tree")
)

type FileTree struct {
	Root *Node
}

// IMPORTANT: Read Only after creation
type Node struct {
	Name     string  // File or directory name
	IsDir    bool    // Indicates if it's a directory
	Children []*Node // Nested files and directories
}

func BuildTree(rootPath string) (*FileTree, error) {
	rootInfo, err := os.Stat(rootPath)
	if err != nil {
		return nil, fmt.Errorf("Error accessing path %q: %v", rootPath, err)
	}

	root := &Node{Name: rootInfo.Name(), IsDir: rootInfo.IsDir()}
	if root.IsDir {
		buildTreeHelper(root, rootPath)
	}

	return &FileTree{Root: root}, nil
}

func buildTreeHelper(parent *Node, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory %q: %v\n", path, err)
		return
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		node := &Node{Name: entry.Name(), IsDir: entry.IsDir()}

		if entry.IsDir() {
			buildTreeHelper(node, childPath)
		}

		parent.Children = append(parent.Children, node)
	}
}
