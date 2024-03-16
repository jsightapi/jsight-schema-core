package openapi

type PropertyInfo interface {
	// SchemaObject returns an OpenAPI Schema Object based on the JSight schema root
	SchemaObject() SchemaObject

	// Optional returns the value for the "optional" rule of the JSight schema root
	Optional() bool

	// Annotation returns the JSight schema root annotation
	Annotation() string
}
