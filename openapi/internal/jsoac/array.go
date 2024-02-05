package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Array struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Items   []Node  `json:"items"`
}

func newArray(astNode schema.ASTNode) Array {
	a := Array{
		OADType: OADTypeArray,
		Items:   make([]Node, 0, len(astNode.Children)),
	}

	for _, an := range astNode.Children {
		a.Items = append(a.Items, newNode(an))
	}

	return a
}

func (a Array) JSightTokenType() schema.TokenType {
	return a.jstType
}
