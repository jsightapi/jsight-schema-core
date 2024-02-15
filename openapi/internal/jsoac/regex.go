package jsoac

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
)

type Regex struct {
	value []byte
}

var _ json.Marshaler = Regex{}
var _ json.Marshaler = &Regex{}

// regex creates an example for: string regex value
func newRegex(astNode schema.ASTNode) *Regex {
	if astNode.Rules.Has("regex") {
		return &Regex{value: quotedBytes(astNode.Rules.GetValue("regex").Value)}
	}
	return nil
}

func (ex Regex) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
