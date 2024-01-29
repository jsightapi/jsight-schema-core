package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	oschema "github.com/jsightapi/jsight-schema-core/openapi/internal/schema"
)

type SchemaObject interface {
	MarshalJSON() (b []byte, err error)
}

type SchemaKeeper interface {
	*jschema.JSchema | *regex.RSchema | schema.ASTNode
}

func NewSchemaObject[T SchemaKeeper](s T) SchemaObject {
	switch ss := any(s).(type) {
	case *jschema.JSchema:
		return oschema.NewFromJSchema(ss)
	case schema.ASTNode:
		return oschema.NewFromASTNode(ss)
	case *regex.RSchema:
		panic("TODO regex.RSchema") // TODO regex schema
	}
	panic(errs.ErrRuntimeFailure.F())
}
