package jsoac

import (
	"encoding/json"
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

type AdditionalProperties struct {
	boolean bool
}

var _ json.Marshaler = AdditionalProperties{}
var _ json.Marshaler = &AdditionalProperties{}

func newAdditionalProperties(astNode schema.ASTNode) *AdditionalProperties {
	if astNode.Rules.Has("additionalProperties") {
		r := astNode.Rules.GetValue("additionalProperties")
		if r.TokenType == schema.TokenTypeBoolean && r.Value == stringFalse {
			return newAdditionalPropertiesFalse()
		}
	} else { // rule not set
		return newAdditionalPropertiesFalse()
	}
	return nil
}

func newAdditionalPropertiesFalse() *AdditionalProperties {
	return &AdditionalProperties{boolean: false}
}

func (a AdditionalProperties) MarshalJSON() ([]byte, error) {
	if a.boolean == false {
		return []byte(stringFalse), nil
	}
	return nil, errs.ErrRuntimeFailure.F()
}
