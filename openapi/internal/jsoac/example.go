package jsoac

import (
	"encoding/json"
)

type Example struct {
	value []byte
}

var _ json.Marshaler = Example{}
var _ json.Marshaler = &Example{}

// newExample creates an example for: integer, number, boolean, string values
func newExample(ex string, t OADType) Example {
	if t == OADTypeString {
		return Example{value: quotedBytes(ex)}
	} else {
		return Example{value: []byte(ex)}
	}
}

func (ex Example) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
