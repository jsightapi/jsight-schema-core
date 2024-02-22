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
	if 0 < len(astNode.Comment) {
		comment := regexp.MustCompile(`\s+`).ReplaceAllString(astNode.Comment, " ")
		return &Description{value: quotedBytes(comment)}
	}
	return nil
}

func (ex Description) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
