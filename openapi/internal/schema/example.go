package schema

import (
	"encoding/json"
)

type Example struct {
	value []byte
}

var _ json.Marshaler = Example{}
var _ json.Marshaler = &Example{}

// newBasicExample creates an example for: integer, number, boolean values
func newBasicExample(ex string) Example {
	return Example{value: []byte(ex)}
}

// newStringExample creates an example for: string value
func newStringExample(ex string) Example {
	return Example{value: quotedBytes(ex)}
}

func (ex Example) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
