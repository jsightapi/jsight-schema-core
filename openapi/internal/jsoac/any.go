package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Any struct {
	Example     *Example     `json:"example,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

func newAny(astNode schema.ASTNode) Any {
	a := Any{
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}

	switch astNode.TokenType {
	case schema.TokenTypeString:
		a.Example = newExample(astNode.Value, true)
	case schema.TokenTypeNumber, schema.TokenTypeBoolean, schema.TokenTypeNull:
		a.Example = newExample(astNode.Value, false)
	default:
		a.Example = nil
	}

	return a
}

func (Any) IsOpenAPINode() bool {
	return true
}
