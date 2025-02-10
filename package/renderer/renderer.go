package renderer

import "github.com/EdgeLordKirito/goFTree/package/filetree"

type Renderer interface {
	Render(tree *filetree.FileTree) (string, error)
}
