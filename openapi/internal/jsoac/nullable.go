package jsoac

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
)

type Nullable struct {
	value []byte
}

var _ json.Marshaler = Nullable{}
var _ json.Marshaler = &Nullable{}

func newNullable(astNode schema.ASTNode) *Nullable {
	if astNode.Rules.Has("nullable") && astNode.Rules.GetValue("nullable").Value == stringTrue {
		return &Nullable{[]byte(`true`)}
	}
	return nil
}

func (n Nullable) MarshalJSON() (b []byte, err error) {
	return n.value, nil
}
