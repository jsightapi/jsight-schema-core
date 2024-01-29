package schema

type OADType int // OpenAPI Data Types

//go:generate stringer -type=OADType -linecomment
const (
	OADTypeString  OADType = iota // string
	OADTypeInteger                // integer
	OADTypeNumber                 // number
	OADTypeBoolean                // boolean
	OADTypeArray                  // array
	OADTypeObject                 // object
)

func (t OADType) MarshalJSON() (b []byte, err error) {
	return quotedBytes(t.String()), nil
}
