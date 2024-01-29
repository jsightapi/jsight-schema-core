package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Null struct {
	jstType schema.TokenType
}

func newNull() Null { // TODO astNode rules etc.
	return Null{
		jstType: schema.TokenTypeNull,
	}
}

func (n Null) JSightTokenType() schema.TokenType {
	return n.jstType
}

func (n Null) MarshalJSON() (b []byte, err error) {
	return []byte(`{"enum": [null], "example": null}`), nil
}
