package asciitree

import (
	"github.com/EdgeLordKirito/goFTree/package/filetree"
	"strings"
)

type Engine struct{}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {
	if tree == nil {
		return "", filetree.ErrEmptyTree()
	}

	var sb strings.Builder

	sb.WriteString(tree.Root.Name + "/\n")

	for i, child := range tree.Root.Children {
		render(&sb, child, "", i == len(tree.Root.Children)-1)
	}

	return sb.String(), nil
}

func render(sb *strings.Builder, node *filetree.Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "+-- "
	if isLast {
		connector = "`-- "
	}

	sb.WriteString(prefix + connector + node.Name)
	if node.IsDir {
		sb.WriteString("/")
	}
	sb.WriteString("\n")

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "|   "
	}

	// Recursively render children
	for i, child := range node.Children {
		render(sb, child, newPrefix, i == len(node.Children)-1)
	}
}
