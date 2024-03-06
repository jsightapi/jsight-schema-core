package jsoac

import (
	"encoding/json"
	"regexp"

	schema "github.com/jsightapi/jsight-schema-core"
)

type Description struct {
	value []byte
}

var _ json.Marshaler = Description{}
var _ json.Marshaler = &Description{}

func newDescription(astNode schema.ASTNode) *Description {
	return newDescriptionFromString(astNode.Comment)
}

func newDescriptionFromString(s string) *Description {
	if 0 < len(s) {
		s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
		return &Description{value: quotedBytes(s)}
	}
	return nil
}

func (ex Description) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
