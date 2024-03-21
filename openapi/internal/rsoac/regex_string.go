package rsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

var _ Node = (*RegexString)(nil)

type RegexString struct {
	Pattern *Pattern `json:"pattern"`
	Type    string   `json:"type"`
}

func newRegexString(astNode schema.ASTNode) Node {
	var p = RegexString{
		Pattern: newPattern(astNode.Value),
		Type:    "string",
	}
	return &p
}
