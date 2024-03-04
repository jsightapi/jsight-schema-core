package jsoac

import (
	"encoding/json"
)

type Example struct {
	value []byte
}

var _ json.Marshaler = Example{}
var _ json.Marshaler = &Example{}

// newExample creates an example value for primitive types
func newExample(ex string, isString bool) *Example {
	if isString {
		return &Example{value: quotedBytes(ex)}
	} else {
		return &Example{value: []byte(ex)}
	}
}

func (ex Example) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
