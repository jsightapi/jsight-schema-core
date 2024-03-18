package info

type ElementType int

const (
	ElementTypeRegex ElementType = iota
	ElementTypeObject
	ElementTypeArray
	ElementTypeScalar // string, integer, boolean, null
	ElementTypeAny
)
