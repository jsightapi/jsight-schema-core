package constraint

import (
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errors"
)

const colonTrue = ": true"
const colonFalse = ": false"

func parseUint(v bytes.Bytes, c Type) uint {
	u, err := v.ParseUint()
	if err != nil {
		panic(errors.Format(errors.ErrInvalidValueOfConstraint, c.String()))
	}
	return u
}
