package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Array struct {
	jstType  schema.TokenType
	OADType  OADType    `json:"type"`
	Items    ArrayItems `json:"items"`
	Nullable *Nullable  `json:"nullable,omitempty"`
	MaxItems *int64     `json:"maxItems,omitempty"`
}

func newArray(astNode schema.ASTNode) Array {
	maxItems := int64RefByString("")
	if len(astNode.Children) == 0 {
		maxItems = int64Ref(0)
	}
	a := Array{
		OADType:  OADTypeArray,
		Items:    newArrayItems(len(astNode.Children)),
		Nullable: newNullable(astNode),
		MaxItems: maxItems,
	}
	for _, an := range astNode.Children {
		a.appendItem(an)
	}
	return a
}

func (a *Array) appendItem(astNode schema.ASTNode) {
	value := newNode(astNode)
	a.Items.append(value)
}

func (a Array) JSightTokenType() schema.TokenType {
	return a.jstType
}
