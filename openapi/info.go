package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/info"

type ElementInfo interface {
	Type() info.ElementType
}

type ObjectInfo interface {
	Type() info.ElementType
	PropertiesInfos() []PropertyInfo
}

type PropertyInfo interface {
	// SchemaObject returns an OpenAPI Schema Object based on the JSight schema root
	SchemaObject() SchemaObject

	// Key return the object key name
	Key() string

	// Optional returns the value for the "optional" rule of the JSight schema root
	Optional() bool

	// Annotation returns the JSight schema root annotation
	Annotation() string
}

var _ ElementInfo = info.ElementInfo{}
var _ ElementInfo = (*info.ElementInfo)(nil)

var _ ObjectInfo = ObjectInfoImpl{}
var _ ObjectInfo = (*ObjectInfoImpl)(nil)

var _ PropertyInfo = PropertyInfoImpl{}
var _ PropertyInfo = (*PropertyInfoImpl)(nil)
