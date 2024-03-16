package openapi

type ReferenceTargetType int

const (
	ReferenceTargetTypeRegex ReferenceTargetType = iota
	ReferenceTargetTypeObject
	ReferenceTargetTypeArray
	ReferenceTargetTypeScalar // string, integer, boolean
)

type ReferenceTargetInfo interface { // TODO rename to ElementInfo?
	ReferenceTargetType() ReferenceTargetType
}
