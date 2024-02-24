package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

type OADType int // OpenAPI Data Types

//go:generate stringer -type=OADType -linecomment
const (
	OADTypeString  OADType = iota // string
	OADTypeInteger                // integer
	OADTypeNumber                 // number
	OADTypeBoolean                // boolean
	OADTypeArray                  // array
	OADTypeObject                 // object
)

func (t OADType) MarshalJSON() (b []byte, err error) {
	return quotedBytes(t.String()), nil
}

func newOADType(astNode schema.ASTNode) OADType {
	if astNode.SchemaType == "integer" {
		return OADTypeInteger
	}
	switch astNode.TokenType {
	case schema.TokenTypeNumber:
		return OADTypeNumber
	case schema.TokenTypeString:
		return OADTypeString
	case schema.TokenTypeBoolean:
		return OADTypeBoolean
	case schema.TokenTypeArray:
		return OADTypeArray
	case schema.TokenTypeObject:
		return OADTypeObject
	default:
		// schema.TokenTypeShortcut:
		// schema.TokenTypeObject
		panic(errs.ErrRuntimeFailure.F())
	}
}
