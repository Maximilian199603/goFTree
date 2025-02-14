package generaltree

import (
	"errors"
	"strings"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
)

var (
	ErrDirPrependerNotSet error = errors.New(
		"Cannot apply Dir prepender as it is not Set")
	ErrDirAppenderNotSet error = errors.New(
		"Cannot apply Dir appender as it is not Set")
	ErrFilePrependerNotSet error = errors.New(
		"Cannot apply File prepender as it is not Set")
	ErrFileAppenderNotSet error = errors.New(
		"Cannot apply File appender as it is not Set")
	Noop func(string, *filetree.Node) (string, *filetree.Node) = produceNoop()
)

func produceNoop() func(string, *filetree.Node) (string, *filetree.Node) {
	return func(s string,
		n *filetree.Node) (string, *filetree.Node) {
		return s, n
	}

}

type RenderSettings struct {
	DirTJunction  string //"├── "
	DirLJunction  string //"└── "
	FileTJunction string
	FileLJunction string
	NoJunction    string //"│   "
	Empty         string

	DirPrepender  func(string, *filetree.Node) (string, *filetree.Node)
	DirAppender   func(string, *filetree.Node) (string, *filetree.Node)
	FilePrepender func(string, *filetree.Node) (string, *filetree.Node)
	FileAppender  func(string, *filetree.Node) (string, *filetree.Node)
}

func (self RenderSettings) HasError() error {
	if self.DirPrepender == nil {
		return ErrDirPrependerNotSet
	}
	if self.DirAppender == nil {
		return ErrDirAppenderNotSet
	}
	if self.FilePrepender == nil {
		return ErrFilePrependerNotSet
	}
	if self.FileAppender == nil {
		return ErrFileAppenderNotSet
	}
	return nil
}

func applyDir(set *RenderSettings, target string, node *filetree.Node) string {
	str, _ := set.DirAppender(set.DirPrepender(target, node))
	return str
}

func applyFile(set *RenderSettings, target string, node *filetree.Node) string {
	str, _ := set.FileAppender(set.FilePrepender(target, node))
	return str
}

func buildLine(sb *strings.Builder, prefix, connector string,
	node *filetree.Node, settings *RenderSettings) {
	if node.IsDir {
		sb.WriteString(prefix + connector + applyDir(settings, node.Name, node) + "\n")
	} else {
		sb.WriteString(prefix + connector + applyFile(settings, node.Name, node) + "\n")
	}
}

func Render(tree *filetree.FileTree, settings *RenderSettings) (string, error) {
	if tree == nil {
		return "", filetree.ErrEmptyTree()
	}
	if err := settings.HasError(); err != nil {
		return "", err
	}

	var sb strings.Builder

	sb.WriteString(applyDir(settings, tree.Root.Name, tree.Root) + "\n")
	for i, child := range tree.Root.Children {
		render(&sb, child, "", i == len(tree.Root.Children)-1, settings)
	}
	return sb.String(), nil
}

func render(sb *strings.Builder, node *filetree.Node,
	prefix string, isLast bool, settings *RenderSettings) {
	if node == nil {
		return
	}

	connector := ""
	if node.IsDir {
		connector = settings.DirTJunction
		if isLast {
			connector = settings.DirLJunction
		}
	} else {
		connector = settings.FileTJunction
		if isLast {
			connector = settings.FileLJunction
		}
	}

	buildLine(sb, prefix, connector, node, settings)

	newPrefix := prefix
	if isLast {
		newPrefix += settings.Empty
	} else {
		newPrefix += settings.NoJunction
	}

	// Recursively render children
	for i, child := range node.Children {
		render(sb, child, newPrefix, i == len(node.Children)-1, settings)
	}
}
