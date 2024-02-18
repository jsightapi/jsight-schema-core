package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Example Example `json:"example"`
	// optional fields
	Pattern          *Pattern `json:"pattern,omitempty"`
	Format           *string  `json:"format,omitempty"`
	Enum             *Enum    `json:"enum,omitempty"`
	Minimum          *Example `json:"minimum,omitempty"`
	Maximum          *Example `json:"maximum,omitempty"`
	ExclusiveMinimum *bool    `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum *bool    `json:"exclusiveMaximum,omitempty"`
	MinLength        *Example `json:"minLength,omitempty"`
	MaxLength        *Example `json:"maxLength,omitempty"`
}

func newPrimitive(t OADType, astNode schema.ASTNode) Primitive {
	if astNode.SchemaType == "integer" {
		t = OADTypeInteger
	}
	var p = Primitive{
		OADType:          t,
		Example:          newExample(astNode.Value, t),
		Pattern:          newPattern(astNode),
		Format:           newFormat(astNode),
		Enum:             newEnum(astNode, t),
		Minimum:          newMinimum(astNode, t),
		Maximum:          newMaximum(astNode, t),
		ExclusiveMinimum: newExclusiveMinimum(astNode, t),
		ExclusiveMaximum: newExclusiveMaximum(astNode, t),
		MinLength:        newMinLength(astNode, t),
		MaxLength:        newMaxLength(astNode, t),
	}
	return p
}

func (p Primitive) JSightTokenType() schema.TokenType {
	return p.jstType
}
