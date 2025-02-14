package xmltree

import (
	"encoding/xml"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
)

type Engine struct{}

type xmlNode struct {
	XMLName  xml.Name  `xml:"node"`
	Name     string    `xml:"name,attr"`
	IsDir    bool      `xml:"isDir,attr"`
	Children []xmlNode `xml:"children>node,omitempty"`
}

func (e Engine) Render(tree *filetree.FileTree) (string, error) {
	if tree == nil {
		return "", filetree.ErrEmptyTree()
	}

	rootXML := convertToXMLNode(tree.Root)
	output, err := xml.MarshalIndent(rootXML, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(output), nil
}

func convertToXMLNode(node *filetree.Node) xmlNode {
	if node == nil {
		return xmlNode{}
	}

	children := make([]xmlNode, len(node.Children))
	for i, child := range node.Children {
		children[i] = convertToXMLNode(child)
	}

	return xmlNode{
		Name:     node.Name,
		IsDir:    node.IsDir,
		Children: children,
	}
}
