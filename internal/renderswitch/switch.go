package renderswitch

import (
	"fmt"

	"github.com/EdgeLordKirito/goFTree/package/renderer"
	"github.com/EdgeLordKirito/goFTree/package/renderers/asciitree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/dashedtree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/jsontree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/linetree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/markdowntree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/xmltree"
)

func GetRenderEngine(identifier string) (renderer.Renderer, error) {
	switch identifier {
	case "ascii":
		return asciitree.Engine{}, nil
	case "markdown":
		return markdowntree.Engine{}, nil
	case "xml":
		return xmltree.Engine{}, nil
	case "json":
		return jsontree.Engine{}, nil
	case "line":
		return linetree.Engine{}, nil
	case "dashed":
		return dashedtree.Engine{}, nil
	default:
		return nil, fmt.Errorf("Unrecognised Identifier: '%s'", identifier)
	}
}
