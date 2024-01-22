package openapi

import (
	"encoding/json"
)

type Example struct {
	value []byte
}

var _ json.Marshaler = Example{}
var _ json.Marshaler = &Example{}

func newExample(t Type, s string) Example {
	b := []byte(s)

	if t == TypeString {
		ex := Example{
			value: make([]byte, 0, len(b)+2),
		}
		var q byte = '"'
		ex.value = append(ex.value, q)
		ex.value = append(ex.value, b...)
		ex.value = append(ex.value, q)

		return ex
	} else {
		return Example{
			value: b,
		}
	}
}

func (ex Example) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
