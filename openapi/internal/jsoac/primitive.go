package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Example Example `json:"example"`
	// optional fields
	Pattern *Regex  `json:"pattern,omitempty"`
	Format  *string `json:"format,omitempty"`
}

func newBasicNode(t OADType, astNode schema.ASTNode) Primitive {
	var pType = t
	if astNode.SchemaType == "integer" {
		pType = OADTypeInteger
	}
	return Primitive{
		OADType: pType,
		Example: newBasicExample(t, astNode.Value),
		Pattern: getRegex(astNode),
		Format:  getFormat(astNode),
	}
}

func getRegex(astNode schema.ASTNode) *Regex {
	var regex *Regex = nil
	if astNode.Rules.Has("regex") {
		regex = newStringRegex(astNode.Rules.GetValue("regex").Value)
	}
	return regex
}

func getFormat(astNode schema.ASTNode) *string {
	var format *string = nil
	if astNode.Rules.Has("type") {
		switch astNode.Rules.GetValue("type").Value {
		case string(schema.SchemaTypeEmail):
			var v = "email"
			format = &v
		case string(schema.SchemaTypeURI):
			var v = "uri"
			format = &v
		case string(schema.SchemaTypeUUID):
			var v = "uuid"
			format = &v
		case string(schema.SchemaTypeDate):
			var v = "date"
			format = &v
		case string(schema.SchemaTypeDateTime):
			var v = "date-time"
			format = &v
		case string(schema.SchemaTypeFloat):
			var v = "float"
			format = &v
		default:
			format = nil
		}
	}
	return format
}

func (p Primitive) JSightTokenType() schema.TokenType {
	return p.jstType
}
