package constraint

import (
	"net/mail"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errors"
	"github.com/jsightapi/jsight-schema-core/json"
)

type Email struct{}

var (
	_ Constraint       = Email{}
	_ Constraint       = (*Email)(nil)
	_ LiteralValidator = Email{}
	_ LiteralValidator = (*Email)(nil)
)

func NewEmail() *Email {
	return &Email{}
}

func (Email) IsJsonTypeCompatible(t json.Type) bool {
	return t == json.TypeString
}

func (Email) Type() Type {
	return EmailConstraintType
}

func (Email) String() string {
	return EmailConstraintType.String()
}

func (Email) Validate(email bytes.Bytes) {
	email = email.Unquote()

	if email.Len() == 0 {
		panic(errors.ErrEmptyEmail)
	}

	char := email.FirstByte()
	if char == ' ' || char == '<' {
		panic(errors.Format(errors.ErrInvalidEmail, email.String()))
	}

	char = email.LastByte()
	if char == ' ' || char == '>' {
		panic(errors.Format(errors.ErrInvalidEmail, email.String()))
	}

	emailStr := email.String()

	_, err := mail.ParseAddress(emailStr)
	if err != nil {
		panic(errors.Format(errors.ErrInvalidEmail, emailStr))
	}
}

func (Email) ASTNode() schema.RuleASTNode {
	return newEmptyRuleASTNode()
}
