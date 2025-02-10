package jsontree

import (
	"encoding/json"
	"fmt"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
)

type Engine struct{}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {
	if tree.Root == nil {
		return "", filetree.ErrEmptyTree()
	}
	jsonData, err := json.MarshalIndent(tree.Root, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error generating JSON: %v", err)
	}
	return string(jsonData), nil
}
