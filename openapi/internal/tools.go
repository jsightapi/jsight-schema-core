package internal

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/internal/sync"

	"strconv"
)

const StringTrue = "true"
const StringFalse = "false"
const StringNull = "null"
const StringAny = "any"
const StringEnum = "enum"
const StringArray = "array"

var BufferPool = sync.NewBufferPool(1024)

// ToJSONString returns JSON quoted string data
func ToJSONString(s string) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return b
}

func IsNullable(astNode schema.ASTNode) bool {
	if astNode.Rules.Has("nullable") && astNode.Rules.GetValue("nullable").Value == StringTrue {
		return true
	}
	return false
}

func IsString(astNode schema.ASTNode) bool {
	return astNode.TokenType == schema.TokenTypeString
}

func Int64Ref(i int64) *int64 {
	return &i
}

func Int64RefByString(s string) *int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return Int64Ref(value)
}
