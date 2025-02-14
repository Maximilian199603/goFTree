package dashedtree

import (
	"github.com/EdgeLordKirito/goFTree/package/filetree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/generaltree"
)

type Engine struct{}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {

	set := &generaltree.RenderSettings{
		DirTJunction:  "├── ",
		DirLJunction:  "└── ",
		FileTJunction: "├╌╌ ",
		FileLJunction: "└╌╌ ",
		NoJunction:    "│   ",
		Empty:         "    ",

		DirPrepender:  generaltree.Noop,
		DirAppender:   func(s string, n *filetree.Node) (string, *filetree.Node) { return s + "/", n },
		FilePrepender: generaltree.Noop,
		FileAppender:  generaltree.Noop,
	}
	return generaltree.Render(tree, set)
}
