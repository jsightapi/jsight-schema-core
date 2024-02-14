package jsoac

import (
	"encoding/json"
)

type Regex struct {
	value []byte
}

var _ json.Marshaler = Regex{}
var _ json.Marshaler = &Regex{}

// newStringRegex creates an example for: string regex value
func newStringRegex(ex string) *Regex {
	return &Regex{value: quotedBytes(ex)}
}

func (ex Regex) MarshalJSON() (b []byte, err error) {
	return ex.value, nil
}
