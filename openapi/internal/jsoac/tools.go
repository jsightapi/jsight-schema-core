package jsoac

import (
	"encoding/json"

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
const stringString = "string"
const stringInteger = "integer"
const stringBoolean = "boolean"
const stringDate = "date"
const stringFloat = "float"
const stringObject = "object"
const stringNumber = "number"
const stringDatetime = "datetime"
const stringEmail = "email"
const stringUri = "uri"
const stringUuid = "uuid"
const stringAdditionalProperties = "additionalProperties"

var bufferPool = sync.NewBufferPool(1024)

// toJSONString returns JSON quoted string data
func toJSONString(s string) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return b
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

func stringRef(s string) *string {
	return &s
}

func boolRef(b bool) *bool {
	return &b
}

func refRuleASTNode(r schema.RuleASTNode) *schema.RuleASTNode {
	return &r
}
