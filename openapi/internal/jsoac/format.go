package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newFormat(astNode schema.ASTNode) *string {
	if astNode.Rules.Has("type") {
		return newFormatFromString(astNode.Rules.GetValue("type").Value)
	}
	return nil
}

func newFormatFromString(s string) *string {
	switch s {
	case string(schema.SchemaTypeEmail), string(schema.SchemaTypeURI),
		string(schema.SchemaTypeUUID), string(schema.SchemaTypeDate):
		return strRef(s)
	case string(schema.SchemaTypeDateTime):
		return strRef("date-time")
	default:
		return nil
	}
}
