package openapi

type ElementType int

const (
	ElementTypeRegex ElementType = iota
	ElementTypeObject
	ElementTypeArray
	ElementTypeScalar // string, integer, boolean, null
	ElementTypeAny
)
