package schema

type Type int

//go:generate stringer -type=Type -linecomment
const (
	TypeString  Type = iota // string
	TypeInteger             // integer
	TypeNumber              // number
	TypeBoolean             // boolean
	TypeArray               // array
	TypeObject              // object
)

func (t Type) MarshalJSON() (b []byte, err error) {
	return quotedBytes(t.String()), nil
}
