package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newFormat(astNode schema.ASTNode) *string {
	if astNode.Rules.Has("type") {
		value := astNode.Rules.GetValue("type").Value
		switch value {
		case string(schema.SchemaTypeEmail), string(schema.SchemaTypeURI),
			string(schema.SchemaTypeUUID), string(schema.SchemaTypeDate):
			return strRef(value)
		case string(schema.SchemaTypeDateTime):
			return strRef("date-time")
		default:
			return nil
		}
	}
	return nil
}
