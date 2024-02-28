package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Null struct {
	Example     Example      `json:"example"`
	Enum        *Enum        `json:"enum,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

func newNull(astNode schema.ASTNode) Null {
	return Null{
		Example:     newExample(stringNull, OADTypeInteger),
		Enum:        newEnum(astNode, OADTypeInteger),
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}
}

func (Null) IsOpenAPINode() bool {
	return true
}
