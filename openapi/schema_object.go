package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/rsoac"
)

type SchemaObject interface {
	JSON() (b []byte, err error)
}

type SchemaKeeper interface {
	*jschema.JSchema | *regex.RSchema | schema.ASTNode
}

func NewSchemaObject[T SchemaKeeper](s T) SchemaObject {
	switch st := any(s).(type) {
	case schema.ASTNode:
		return jsoac.NewFromASTNode(st)
	case *jschema.JSchema:
		return jsoac.New(st)
	case *regex.RSchema:
		return rsoac.New(st)
	}

	panic(errs.ErrRuntimeFailure.F())
}
