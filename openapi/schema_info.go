package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/info"
)

type SchemaInfo interface {
	// SchemaObject returns an OpenAPI Schema Object based on the JSight schema root
	SchemaObject() SchemaObject

	// Optional returns the value for the "optional" rule of the JSight schema root
	Optional() bool

	// Annotation returns the JSight schema root annotation
	Annotation() string

	// PropertiesInfos finds the first nested object (only objects and references are processed),
	// returns information on its properties
	PropertiesInfos() PropertiesIterator
}

func NewSchemaInfo(s *jschema.JSchema) SchemaInfo {
	return newSchemaInfoImpl(info.NewInfo(s))
}
