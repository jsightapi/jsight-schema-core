package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

func ruleToASTNode(r schema.RuleASTNode) schema.ASTNode {
	switch r.TokenType {
	case schema.TokenTypeString, schema.TokenTypeShortcut:
		return stringRuleToASTNode(r)
	case schema.TokenTypeObject:
		return objectRuleToASTNode(r)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

// stringRuleToASTNode returns the ASTNode for "OR" rule elements. JSight example: // {or: [ "email", "integer" ]}
func stringRuleToASTNode(r schema.RuleASTNode) schema.ASTNode {
	a := schema.ASTNode{
		Rules: schema.MakeRuleASTNodes(1),
	}

	format := formatFromSchemaType(r.Value)

	if format != nil { // JSight example: // {or: [ "email"... ]}
		a.TokenType = schema.TokenTypeString
		a.Rules.Set("type", schema.RuleASTNode{
			TokenType: schema.TokenTypeString,
			Value:     *format,
		})
	} else { // JSight example: // {or: [ "integer", "@cat" ]}
		a.TokenType = tokenType(r.Value)
		a.SchemaType = r.Value
	}

	return a
}

// objectRuleToASTNode returns the ASTNode for "OR" rule elements. JSight example: // {or: [ {type: "integer"}, ... ]}
func objectRuleToASTNode(r schema.RuleASTNode) schema.ASTNode {
	a := schema.ASTNode{
		Rules: schema.MakeRuleASTNodes(r.Properties.Len()),
	}

	if r.Properties != nil { // or: [ {...} ]
		_ = r.Properties.Each(func(k string, v schema.RuleASTNode) error {
			a.Rules.Set(k, v)
			return nil
		})
	}

	if typeRule, ok := r.Properties.Get("type"); ok { // or: [ { type: ...} ]
		a.TokenType = tokenType(typeRule.Value)
		a.SchemaType = typeRule.Value
	}

	return a
}

func tokenType(s string) schema.TokenType {
	if s[0] == '@' {
		return schema.TokenTypeShortcut
	}

	switch s {
	case "string":
		return schema.TokenTypeString
	case "boolean":
		return schema.TokenTypeBoolean
	case "float", "integer", "decimal":
		return schema.TokenTypeNumber
	case "object":
		return schema.TokenTypeObject
	case "array":
		return schema.TokenTypeArray
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}
