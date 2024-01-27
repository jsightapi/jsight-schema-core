package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/schema"
)

type SchemaObject interface {
	MarshalJSON() (b []byte, err error)
}

func NewSchemaObject(j *jschema.JSchema) SchemaObject {
	return schema.New(j)
}
