package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Or struct {
	AnyOf       []Node       `json:"anyOf,omitempty"`
	Example     *Example     `json:"example,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

func newOr(astNode schema.ASTNode) Or {
	rule := astNode.Rules.GetValue("or")

	var ex *Example = nil
	if astNode.TokenType != schema.TokenTypeShortcut {
		t := oadTypeFromASTNode(astNode)
		ex = newExample(astNode.Value, t == OADTypeString)
	}

	or := Or{
		AnyOf:       newAnyOf(rule.Items),
		Example:     ex,
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}

	return or
}

func newAnyOf(rr []schema.RuleASTNode) []Node {
	nn := make([]Node, 0, len(rr))

	for _, r := range rr {
		ast := ruleToASTNode(r)
		node := newNode(ast)

		if p, ok := node.(Primitive); ok { // fix empty string Example. See JSight {or: [ {type: "integer"} ]}
			p.Example = nil
			node = p
		}

		nn = append(nn, node)
	}

	return nn
}

func (Or) IsOpenAPINode() bool {
	return true
}
