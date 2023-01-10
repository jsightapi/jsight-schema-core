package constraint

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/json"
)

type KeysCaseInsensitive struct {
	value bool
}

var (
	_ Constraint = KeysCaseInsensitive{}
	_ Constraint = (*KeysCaseInsensitive)(nil)
	_ BoolKeeper = KeysCaseInsensitive{}
	_ BoolKeeper = (*KeysCaseInsensitive)(nil)
)

func NewKeysCaseInsensitive(b bool) *KeysCaseInsensitive {
	return &KeysCaseInsensitive{
		value: b,
	}
}

func (KeysCaseInsensitive) IsJsonTypeCompatible(t json.Type) bool {
	return t == json.TypeObject
}

func (KeysCaseInsensitive) Type() Type {
	return KeysCaseInsensitiveConstraintType
}

func (c KeysCaseInsensitive) String() string {
	str := "[ UNVERIFIABLE CONSTRAINT ] " + KeysCaseInsensitiveConstraintType.String()
	if c.value {
		str += colonTrue
	} else {
		str += colonFalse
	}
	return str
}

func (c KeysCaseInsensitive) Bool() bool {
	return c.value
}

func (KeysCaseInsensitive) ASTNode() schema.RuleASTNode {
	panic(errs.ErrRuntimeFailure)
}
