package rsoac

// Regex schema to OpenAPi converter

import (
	"github.com/jsightapi/jsight-schema-core/notations/regex"
)

type RSOAC struct {
	description string
}

func New(j *regex.RSchema) *RSOAC {
	panic("TODO regex.RSchema") // TODO method
}

func (o *RSOAC) SetDescription(s string) {
	o.description = s
}

func (o RSOAC) MarshalJSON() (b []byte, err error) {
	return []byte("TODO"), nil // TODO method
}
