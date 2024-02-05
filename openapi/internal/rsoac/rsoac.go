package rsoac

// Regex schema to OpenAPi converter

import (
	"github.com/jsightapi/jsight-schema-core/notations/regex"
)

type RSOAC struct {
	// TODO props
}

func New(j *regex.RSchema) RSOAC {
	panic("TODO regex.RSchema") // TODO method
}

func (o RSOAC) JSON() (b []byte, err error) {
	return []byte("TODO"), nil // TODO method
}
