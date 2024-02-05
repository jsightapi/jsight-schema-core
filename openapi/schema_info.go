package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/info"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/rsoac"
)

type SchemaInfo interface {
	// SchemaObject returns an OpenAPI Schema Object based on the JSight schema root
	SchemaObject() SchemaObject

	// Optional returns the value for the "optional" rule of the JSight schema root
	Optional() bool

	// Annotation returns the JSight schema root annotation
	Annotation() string

	// SetAnnotation sets an annotation for the JSight schema root
	SetAnnotation(s string)

	// NestedObjectProperties finds the first nested object (only objects and references are processed),
	// returns information on its properties
	NestedObjectProperties() []SchemaInfo // returns an empty slice if the object is not found
}

var _ SchemaObject = jsoac.JSOAC{}
var _ SchemaObject = &jsoac.JSOAC{}

var _ SchemaObject = rsoac.RSOAC{}
var _ SchemaObject = &rsoac.RSOAC{}

func NewSchemaInfo(s *jschema.JSchema) SchemaInfo {
	return newSchemaInfoImpl(info.New(s))
}
