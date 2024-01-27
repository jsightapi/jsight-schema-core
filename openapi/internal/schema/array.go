package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Array struct {
	Type_ Type   `json:"type"`
	Items []Node `json:"items"`
}

func newArray(astNode schema.ASTNode) Array {
	a := Array{
		Type_: TypeArray,
		Items: make([]Node, 0, len(astNode.Children)),
	}

	for _, an := range astNode.Children {
		a.Items = append(a.Items, newNode(an))
	}

	return a
}

func (a Array) Type() Type {
	return a.Type_
}
