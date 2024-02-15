package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Example Example `json:"example"`
	// optional fields
	Pattern  *Regex      `json:"pattern,omitempty"`
	Format   *string     `json:"format,omitempty"`
	Required *bool       `json:"required,omitempty"`
	CEnum    *[1]Example `json:"enum,omitempty"`
}

func newBasicNode(t OADType, astNode schema.ASTNode) Primitive {
	var pType = t
	if astNode.SchemaType == "integer" {
		pType = OADTypeInteger
	}
	var p = Primitive{
		OADType: pType,
		Example: newBasicExample(t, astNode.Value),
		Pattern: getRegex(astNode),
		Format:  getFormat(astNode),
	}
	if astNode.Rules.Has("const") {
		addConst(astNode, &p)
	}
	return p
}

func (p Primitive) JSightTokenType() schema.TokenType {
	return p.jstType
}
