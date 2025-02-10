package markdowntree

import (
	"strings"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
)

const (
	level        string = "  "
	markdownList string = "- "
)

type Engine struct{}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {
	if tree == nil {
		return "", filetree.ErrEmptyTree()
	}
	var sb strings.Builder
	render(&sb, tree.Root, 0)
	return sb.String(), nil
}

func render(sb *strings.Builder, node *filetree.FileNode, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat(level, depth)
	prefix := markdownList

	sb.WriteString(indent + prefix + node.Name)
	if node.IsDir {
		sb.WriteString("/")
	}
	sb.WriteString("\n")

	for _, child := range node.Children {
		render(sb, child, depth+1)
	}
}
