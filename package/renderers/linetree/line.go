package linetree

import "github.com/EdgeLordKirito/goFTree/package/filetree"

type Engine struct{}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {
	if tree.Root == nil {
		return "", filetree.ErrEmptyTree()
	}

	displayName := tree.Root.Name
	if tree.Root.IsDir {
		displayName += "/"
	}

	result := displayName + "\n"

	for i, child := range tree.Root.Children {
		isLastChild := i == len(tree.Root.Children)-1
		result += render(child, "", isLastChild)
	}

	return result, nil
}

func render(node *filetree.FileNode, prefix string, isLast bool) string {
	if node == nil {
		return ""
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	displayName := node.Name
	if node.IsDir {
		displayName += "/"
	}

	result := prefix + connector + displayName + "\n"

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	// Iterate through children
	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		result += render(child, newPrefix, isLastChild)
	}

	return result
}
