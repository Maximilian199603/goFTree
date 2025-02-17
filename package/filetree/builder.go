package filetree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func BuildTree(rootPath string, options *HiddenOptions) (*FileTree, error) {
	rootInfo, err := os.Stat(rootPath)
	if err != nil {
		return nil, fmt.Errorf("Error accessing path %q: %v", rootPath, err)
	}

	root := &Node{Name: rootInfo.Name(), IsDir: rootInfo.IsDir()}
	if root.IsDir {
		buildTreeHelper(root, rootPath, options)
	}

	return &FileTree{Root: root}, nil
}

func buildTreeHelper(parent *Node, path string, options *HiddenOptions) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory %q: %v\n", path, err)
		return
	}

	for _, entry := range entries {
		if !options.CanAdd(entry) {
			continue
		}
		childPath := filepath.Join(path, entry.Name())
		node := &Node{Name: entry.Name(), IsDir: entry.IsDir()}

		if entry.IsDir() {
			buildTreeHelper(node, childPath, options)
		}

		parent.Children = append(parent.Children, node)
	}
}

type HiddenOptions struct {
	All   bool
	Dirs  bool
	Files bool
}

func NewHiddenOption(all, dir, file bool) *HiddenOptions {
	return &HiddenOptions{
		All:   all,
		Dirs:  dir,
		Files: file,
	}
}

func (self HiddenOptions) CanAdd(entry os.DirEntry) bool {
	hidden := isHidden(entry.Name())
	if !hidden {
		return true
	}
	if self.All {
		return true
	}
	if entry.IsDir() {
		if self.Dirs {
			return true
		}
	} else {
		if self.Files {
			return true
		}
	}
	return false
}

func isHidden(path string) bool {
	return strings.HasPrefix(path, ".")
}
