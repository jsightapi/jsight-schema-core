package jsoac

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"

	"strings"
)

type Description struct {
	value []byte
}

var _ json.Marshaler = Description{}
var _ json.Marshaler = &Description{}

func newDescription(astNode schema.ASTNode) *Description {
	if 0 < len(astNode.Comment) {
		comment := astNode.Comment
		comment = strings.ReplaceAll(comment, "\n", "\\n")
		comment = strings.ReplaceAll(comment, "\t", "\\t")
		comment = strings.ReplaceAll(comment, "\r", "\\r")
		return &Description{value: quotedBytes(comment)}
	}
	return nil
}

func (ex Description) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
