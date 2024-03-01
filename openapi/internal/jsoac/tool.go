package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/internal/sync"

	"strconv"
)

const stringTrue = "true"
const stringFalse = "false"
const stringNull = "null"
const stringAny = "any"
const stringEnum = "enum"
const stringArray = "array"

var bufferPool = sync.NewBufferPool(1024)

func quotedBytes(s string) []byte {
	bb := make([]byte, 0, len(s)+2)

	bb = append(bb, '"')
	bb = append(bb, []byte(s)...)
	bb = append(bb, '"')

	return bb
}

func isNullable(astNode schema.ASTNode) bool {
	if astNode.Rules.Has("nullable") && astNode.Rules.GetValue("nullable").Value == stringTrue {
		return true
	}
	return false
}

func isString(astNode schema.ASTNode) bool {
	return astNode.TokenType == schema.TokenTypeString
}

func strRef(s string) *string {
	return &s
}

func int64Ref(i int64) *int64 {
	return &i
}

func int64RefByString(s string) *int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return int64Ref(value)
}
