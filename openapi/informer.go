package openapi

type ElementInformer interface {
	Type() ElementType
}

type ObjectInformer interface {
	Type() ElementType
	PropertiesInfos() []PropertyInformer
}

type PropertyInformer interface {
	// SchemaObject returns an OpenAPI Schema Object based on the JSight schema root
	SchemaObject() SchemaObject

	// Key return the object key name
	Key() string

	// Optional returns the value for the "optional" rule of the JSight schema root
	Optional() bool

	// Annotation returns the JSight schema root annotation
	Annotation() string
}
