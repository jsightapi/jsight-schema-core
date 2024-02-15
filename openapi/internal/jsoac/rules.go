package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

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

func addConst(astNode schema.ASTNode, p *Primitive) {
	if astNode.Rules.GetValue("const").Value == "true" {
		var v = true
		p.Required = &v
		var value = newBasicExample(p.OADType, astNode.Value)
		var enum [1]Example
		enum[0] = value
		p.CEnum = &enum
	}
}
