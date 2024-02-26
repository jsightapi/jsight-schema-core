package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type AnyOf struct {
	Items []Node `json:"anyOf"`
}

type Array struct {
	jstType  schema.TokenType
	OADType  OADType    `json:"type"`
	Items    ArrayItems `json:"items"`
	Nullable *Nullable  `json:"nullable,omitempty"`
	MinItems *int64     `json:"minItems,omitempty"`
	MaxItems *int64     `json:"maxItems,omitempty"`
}

func newArray(astNode schema.ASTNode) Array {
	maxItems := newMaxItems(astNode)
	if len(astNode.Children) == 0 {
		maxItems = int64Ref(0)
	}
	a := Array{
		OADType:  OADTypeArray,
		Items:    newArrayItems(len(astNode.Children)),
		Nullable: newNullable(astNode),
		MinItems: newMinItems(astNode),
		MaxItems: maxItems,
	}
	for _, an := range astNode.Children {
		a.appendItem(an)
	}
	return a
}

func (a *Array) appendItem(astNode schema.ASTNode) {
	a.Items.append(newNode(astNode))
}

func (a Array) JSightTokenType() schema.TokenType {
	return a.jstType
}
