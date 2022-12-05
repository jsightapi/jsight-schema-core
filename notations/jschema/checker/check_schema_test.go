package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/errs"

	"github.com/jsightapi/jsight-schema-core/notations/jschema/loader"

	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/scanner"
)

func TestCheckRootSchema(t *testing.T) {
	type typ struct {
		name   string
		schema string
	}

	check := func(schema string, types []typ) {
		schemaFile := fs.NewFile("schema", schema)

		rootSchema := loader.LoadSchema(scanner.New(schemaFile), nil)

		for _, datum := range types {
			f := fs.NewFile(datum.name, datum.schema)
			typ := loader.LoadSchema(scanner.New(f), rootSchema)
			rootSchema.AddNamedType(datum.name, typ, f, 0)
		}

		loader.CompileAllOf(rootSchema)
		loader.AddUnnamedTypes(rootSchema)
		CheckRootSchema(rootSchema)
	}

	t.Run("positive", func(t *testing.T) {
		var tests = []struct {
			schema string
			types  []typ
		}{
			{
				`{}`,
				[]typ{},
			},
			{
				`[1,2,3]`,
				[]typ{},
			},
			{
				`123`,
				[]typ{},
			},
			{
				`"qwerty"`,
				[]typ{},
			},
			{
				`{} // note`,
				[]typ{},
			},
			{
				`123 // note`,
				[]typ{},
			},
			{
				`"qwerty" // note`,
				[]typ{},
			},
			{
				"5 // {min: 1}", // the rule without quotes
				[]typ{},
			},
			{
				`5 // {"min": 1}`, // the rule with quotes
				[]typ{},
			},
			{
				`5 // {min: 1, max: 5}`, // a few rules
				[]typ{},
			},
			{
				"5 // {}", // without rules
				[]typ{},
			},
			{
				"5 // {min: 1} - some comment", // text after rule
				[]typ{},
			},
			{
				"5 // some comment", // text without rules
				[]typ{},
			},
			{
				"5 // - some comment", // text without rules
				[]typ{},
			},
			{
				"5 // [ some comment ]", // text without rules
				[]typ{},
			},

			// typeConstraint: explicit type definition
			{
				`{} // {type: "object"}`,
				[]typ{},
			},
			{
				`true // {type: "boolean"}`,
				[]typ{},
			},
			{
				`"abc" // {type: "string"}`,
				[]typ{},
			},
			{
				`null // {type: "null"}`,
				[]typ{},
			},
			{
				`[ // {type: "array"}
					1
				]`,
				[]typ{},
			},

			// precisionConstraint: decimal type
			{
				`1.1 // {precision: 1}`,
				[]typ{},
			},

			{
				`1.0 // {precision: 1}`,
				[]typ{},
			},
			{
				`1.00 // {precision: 2}`,
				[]typ{},
			},

			{
				`0.12 // {precision: 2}`,
				[]typ{},
			},

			{
				`0.120 // {precision: 2}`,
				[]typ{},
			},
			{
				`0.1200 // {precision: 2}`,
				[]typ{},
			},

			{
				`123.45 // {type: "decimal", precision: 2}`,
				[]typ{},
			},

			// min
			{
				`123 // {min: 1}`,
				[]typ{},
			},
			{
				`123 // {"min": 1}`,
				[]typ{},
			},
			{
				`-1 // {"min": -2}`,
				[]typ{},
			},
			{
				` 0 // {"min": -2}`,
				[]typ{},
			},
			{
				` 1 // {"min": -2}`,
				[]typ{},
			},

			// max
			{
				`123 // {max: 999}`,
				[]typ{},
			},
			{
				`123 // {"max": 999}`,
				[]typ{},
			},
			{
				`-1 // {"max": 1}`,
				[]typ{},
			},
			{
				` 0 // {"max": 1}`,
				[]typ{},
			},
			{
				` 1 // {"max": 1}`,
				[]typ{},
			},

			// exclusiveMinimumConstraint
			{
				`111 // {min: 1, exclusiveMinimum: true}`,
				[]typ{},
			},
			{
				`111 // {min: 1, exclusiveMinimum: false}`,
				[]typ{},
			},

			// exclusiveMaximumConstraint
			{
				`222 // {max: 333, exclusiveMaximum: true}`,
				[]typ{},
			},
			{
				`222 // {max: 333, exclusiveMaximum: false}`,
				[]typ{},
			},

			// optionalConstraints
			{
				`{
				"key": 1 // {optional: true}
			}`,
				[]typ{},
			},

			// rule "or"
			{
				`5 // {or: [ {type: "integer"}, {type: "string"} ]}`, // "or" with two simple rule-set
				[]typ{},
			},
			{
				`5 // {or: [ {min: 0}, {type: "string"} ]}`, // "or" with two simple rule-set (the first rule-set without type specifying)
				[]typ{},
			},
			{
				`5 // {or: [ {type: "@int"}, {type: "@str"} ]}`, // "or" with type names
				[]typ{
					{`@int`, `123`},
					{`@str`, `"abc"`},
				},
			},
			{
				`5 // {or: [ "@int", "@str" ]}`, // "or" with short format type names
				[]typ{
					{`@int`, `123`},
					{`@str`, `"abc"`},
				},
			},
			{
				`"xyz" // {or: [ "@int", "@str" ]}`, // "or" with short format type names
				[]typ{
					{`@int`, `123`},
					{`@str`, `"abc"`},
				},
			},
			{
				"@int | @str", // "or" shortcut
				[]typ{
					{`@int`, `123`},
					{`@str`, `"abc"`},
				},
			},
			{
				`5 // {or: [ {type: "@int"}, "@str", {min: 0}, {type: "string"} ]}`, // "or" with different format or rule-sets
				[]typ{
					{`@int`, `123`},
					{`@str`, `"abc"`},
				},
			},

			// `{
			// 	"key": [ // {optional: true, minItems: 1}
			// 		123
			// 	]
			// }`,

			// `// text-1
			// // text-2
			//
			// { // text-3
			// 	// text-4
			// 	"aaa": 111, // {min: 1}
			// 	"bbb": 222, // {min: 2, max: 999}
			// 	"ccc": { // text-5
			// 		"ddd": 333, // {min: 3, optional: true}
			// 		"error": [] // {optional: true, minItems: 0}
			// 		// text-6
			// 	}, // text-7
			// 	"eee": 444 // {min: 4}
			// } // text-8
			//
			// // text-9
			// // text-10`,

			// some type (string)
			{
				`"abc" // {type: "@schema"}`,
				[]typ{
					{`@schema`, `"qwerty"`},
				},
			},
			{
				"@schema",
				[]typ{
					{`@schema`, `"qwerty"`},
				},
			},
			// some type (integer)
			{
				`222 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `111`},
				},
			},
			{
				"@schema",
				[]typ{
					{`@schema`, `111`},
				},
			},
			// some type (float)
			{
				`3.4 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `1.2`},
				},
			},
			{
				"@schema",
				[]typ{
					{`@schema`, `1.2`},
				},
			},

			// some type (boolean)
			{
				`false // {type: "@schema"}`,
				[]typ{
					{`@schema`, `true`},
				},
			},
			{
				"@schema",
				[]typ{
					{`@schema`, `true`},
				},
			},

			// some type (object)
			{
				"@schema",
				[]typ{
					{`@schema`, `{
					"key": "val"
				}`},
				},
			},

			// some type (array)
			{
				"@schema",
				[]typ{
					{`@schema`, `[1,2,3]`},
				},
			},

			// email and string
			{
				`"aaa@bbb.cc" // {type: "@email"}`,
				[]typ{
					{`@email`, `"ddd@eee.ff" // {type: "email"}`},
				},
			},
			{
				"@email",
				[]typ{
					{`@email`, `"ddd@eee.ff" // {type: "email"}`},
				},
			},

			// decimal and float
			{
				`3.4 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `1.2 // {precision: 1}`},
				},
			},
			{
				"@schema",
				[]typ{
					{`@schema`, `1.2 // {precision: 1}`},
				},
			},

			// or
			{
				`222 // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
			},
			{
				`"str" // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
			},
			{
				`222 // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
			},
			{
				`"str" // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
			},

			{
				`  222 // {or: [ {type:"integer"}, {type:"string"} ]}`,
				[]typ{},
			},
			{
				`"str" // {or: [ {type:"integer"}, {type:"string"} ]}`,
				[]typ{},
			},

			{
				`1 // {or: ["@int_or_str", "@obj"]}`,
				[]typ{
					{`@int_or_str`, `"abc" // {or: ["@int", "@str"]}`},
					{`@str`, `"abc"`},
					{`@int`, `123`},
					{`@obj`, `{}`},
				},
			},

			{
				`{
				"id": 1,
				"children": [
					@node
				]
			}`,
				[]typ{
					{"@node", `{
					"id": 1,
					"children": [
						@node
					]
				}`},
				},
			},

			{
				`1 // {type: "@type1"}`,
				[]typ{
					{`@type1`, `1 // {or: [ {type:"integer"}, {type:"string"} ]}`},
				},
			},

			// allowed recursions
			{
				"@type1",
				[]typ{
					{"@type1", "@type2"},
					{"@type2", "{}"},
				},
			},

			{
				"@user",
				[]typ{
					{`@user`, `{
					"name": "John",
					"best_friend": @user
				}`},
				},
			},

			// array and its required element
			{
				`[1]`,
				[]typ{},
			},
			{
				`[1,2]`,
				[]typ{},
			},
			{
				"@arr",
				[]typ{
					{`@arr`, `[1,2,3]`},
				},
			},
			{
				"@arr-1",
				[]typ{
					{"@arr-1", "@arr-2"},
					{"@arr-2", "[1,2,3]"},
				},
			},

			// Allow incorrect links-type for unused types
			{
				`1 // {type: "@used"}`,
				[]typ{
					{`@used`, `111`},
					{`@unused-1`, `222 // {type: "@unused-2"}`},
					{`@unused-2`, `333`},
				},
			},

			{
				`"abc" // {type: "enum", enum: [123, "abc"]}`,
				[]typ{},
			},
			{
				`"abc" // {type: "mixed", or: [{type:"integer"}, {type:"string"}]}`,
				[]typ{},
			},

			{
				`{
				"key": "abc" // {type: "mixed", or: [{type:"string"}, {type:"integer"}], optional: true}
			}`,
				[]typ{},
			},
			{
				`{
				"key": "abc" // {type: "enum", enum: [123, "abc"], optional: true}
			}`,
				[]typ{},
			},
			{
				`123   // {type: "any"}`,
				[]typ{},
			},
			{
				`12.3  // {type: "any"}`,
				[]typ{},
			},
			{
				`"str" // {type: "any"}`,
				[]typ{},
			},
			{
				`true  // {type: "any"}`,
				[]typ{},
			},
			{
				`false // {type: "any"}`,
				[]typ{},
			},
			{
				`null  // {type: "any"}`,
				[]typ{},
			},
			{
				`{}    // {type: "any"}`,
				[]typ{},
			},
			{
				`[]    // {type: "any"}`,
				[]typ{},
			},
			{
				`{
				"aaa": 1 // {type: "any", optional: true}
			}`,
				[]typ{},
			},

			{
				`[ // {minItems: 2}
				1,2
			]`,
				[]typ{},
			},
			{
				`[ // {maxItems: 3}
				1,2,3
			]`,
				[]typ{},
			},
			{
				`[]`,
				[]typ{},
			},
			{
				`[] // {minItems: 0}`,
				[]typ{},
			},
			{
				`[] // {maxItems: 0}`,
				[]typ{},
			},
			{
				`[] // {minItems: 0, maxItems: 0}`,
				[]typ{},
			},
			{
				`[] // {type: "array"}`,
				[]typ{},
			},
			{
				"@arr",
				[]typ{
					{"@arr", "[]"},
				},
			},

			{
				"@sub",
				[]typ{
					{"@sub", `{"_id": 123}`},
				},
			},

			{
				"@sub9",
				[]typ{
					{"@sub1", `{"_id\"": 123}`},
					{"@sub2", `{"_id\\": 123}`},
					{"@sub3", `{"_id\/": 123}`},
					{"@sub4", `{"_id\b": 123}`},
					{"@sub5", `{"_id\f": 123}`},
					{"@sub6", `{"_id\n": 123}`},
					{"@sub7", `{"_id\r": 123}`},
					{"@sub8", `{"_id\t": 123}`},
					{"@sub9", `{"_id\uAAAA": 123}`},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.schema, func(t *testing.T) {
				check(tt.schema, tt.types)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		tests := []struct {
			schema string
			types  []typ
			err    errs.Code
		}{
			// min
			{`-1 // {"min": 2}`, []typ{}, errs.ErrConstraintValidation},
			{` 0 // {"min": 2}`, []typ{}, errs.ErrConstraintValidation},
			{` 1 // {"min": 2}`, []typ{}, errs.ErrConstraintValidation},
			{`-3 // {"min": -2}`, []typ{}, errs.ErrConstraintValidation},
			{`-4 // {"min": -2}`, []typ{}, errs.ErrConstraintValidation},

			{` 3 // {"max": 2}`, []typ{}, errs.ErrConstraintValidation},
			{` 4 // {"max": 2}`, []typ{}, errs.ErrConstraintValidation},
			{`-1 // {"max": -2}`, []typ{}, errs.ErrConstraintValidation},
			{` 0 // {"max": -2}`, []typ{}, errs.ErrConstraintValidation},
			{` 1 // {"max": -2}`, []typ{}, errs.ErrConstraintValidation},

			{"2 // {unknown: 123}", []typ{}, errs.ErrUnknownRule},
			{`2 // {type: "unknown"}`, []typ{}, errs.ErrUnknownValueOfTheTypeRule},

			{"2 // {min: [1,2]}", []typ{}, errs.ErrIncorrectRuleValueType},
			{"2 // {min: {}}", []typ{}, errs.ErrIncorrectRuleValueType},

			{`222 // {min: 1, min: 1}`, []typ{}, errs.ErrDuplicateRule}, // duplicate, with the same values
			{`333 // {min: 1, min: 2}`, []typ{}, errs.ErrDuplicateRule}, // duplicate, with different values

			{`{key: 1}`, []typ{}, errs.ErrInvalidCharacter},            // incorrect json example (no quotes)
			{`{"k":1, "k":2}`, []typ{}, errs.ErrDuplicateKeysInSchema}, // duplicate keys on JSON

			{`{"key": "value"} // {min: 1}`, []typ{}, errs.ErrIncorrectRuleForSeveralNode},
			{`[
			1,2 // {min: 1}
		]`, []typ{}, errs.ErrIncorrectRuleForSeveralNode},
			{`{"key": {"bbb": // {min: 1}
			2
		}}`, []typ{}, errs.ErrIncorrectRuleForSeveralNode},
			{`[
			1
		] // {min: 1}`, []typ{}, errs.ErrAnnotationNotAllowed},

			{`{
		} // {min: 1}`, []typ{}, errs.ErrIncorrectRuleWithoutExample},

			// error on constraint validation
			{`3 // {min: 4}`, []typ{}, errs.ErrConstraintValidation},
			{`3 // {max: 2}`, []typ{}, errs.ErrConstraintValidation},
			{`"str" // {minLength: 4}`, []typ{}, errs.ErrConstraintStringLengthValidation},
			{`"str" // {maxLength: 2}`, []typ{}, errs.ErrConstraintStringLengthValidation},

			// Incompatible JSON-type
			{`[ // {min: 1}
			1
		]`, []typ{}, errs.ErrUnexpectedConstraint},
			{`1 // {minLength: 1}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`"str" // {min: 1}`, []typ{}, errs.ErrUnexpectedConstraint},

			// Unable to add constraint "email" to the node "Integer"
			{`123 // {type: "email"}`, []typ{}, errs.ErrIncompatibleTypes},

			// email
			{`"" // {type: "email"}`, []typ{}, errs.ErrEmptyEmail},
			{`"no email" // {type: "email"}`, []typ{}, errs.ErrInvalidEmail},

			// Invalid value of constraint
			{`1 // {min: 99, exclusiveMinimum: 1}`, []typ{}, errs.ErrInvalidValueOfConstraint},
			{`1 // {max: 99, exclusiveMaximum: 1}`, []typ{}, errs.ErrInvalidValueOfConstraint},
			{`1.1 // {precision: true}`, []typ{}, errs.ErrInvalidValueOfConstraint},
			{`{
					"k": 1 // {optional: 1}
		}`, []typ{}, errs.ErrInvalidValueOfConstraint},

			// typeConstraint: incorrect type conversion
			{`"abc" // {type: "integer"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`12.34 // {type: "integer"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`123 // {type: "string"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`true // {type: "string"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`null // {type: "string"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`{} // {type: "string"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`[] // {type: "string"}`, []typ{}, errs.ErrIncompatibleTypes},
			{`123 // {type: "float"}`, []typ{}, errs.ErrIncompatibleTypes},

			// precisionConstraint
			{`123.45 // {type: "decimal"}`, []typ{}, errs.ErrNotFoundRulePrecision},          // decimal without precision
			{`123 // {precision: 1}`, []typ{}, errs.ErrUnexpectedConstraint},                 // incorrect integer node type
			{`"str" // {precision: 2}`, []typ{}, errs.ErrUnexpectedConstraint},               // incorrect string node type
			{`true // {precision: 2}`, []typ{}, errs.ErrUnexpectedConstraint},                // incorrect bool node type
			{`null // {precision: 2}`, []typ{}, errs.ErrUnexpectedConstraint},                // incorrect null node type
			{`"str" // {minLength: 0, precision: 1}`, []typ{}, errs.ErrUnexpectedConstraint}, // incompatibility node type with constraint
			{`1.0 // {precision: 0}`, []typ{}, errs.ErrZeroPrecision},                        // zero precision
			{`0.12 // {precision: -2}`, []typ{}, errs.ErrInvalidValueOfConstraint},           // negative precision
			{`0.12 // {precision: 2.3}`, []typ{}, errs.ErrInvalidValueOfConstraint},          // fractional precision

			// exclusiveMinimumConstraint
			{`111 // {exclusiveMinimum: true}`, []typ{}, errs.ErrConstraintMinNotFound},
			{`111 // {min: 2, exclusiveMinimum: 1}`, []typ{}, errs.ErrInvalidValueOfConstraint}, // not bool in exclusive

			// exclusiveMaximumConstraint
			{`222 // {exclusiveMaximum: true}`, []typ{}, errs.ErrConstraintMaxNotFound},
			{`222 // {max: 2, exclusiveMaximum: "str"}`, []typ{}, errs.ErrInvalidValueOfConstraint}, // not bool in exclusive

			// optionalConstraints: Incorrect rule "optional" location. The rule "optional" applies only to object properties.
			{`"str" // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`12 // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`1.2 // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`true // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`null // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`{} // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`[] // {optional: true}`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`[
				1 // {optional: true}
			]`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},
			{`{ // {optional: true}
            }`, []typ{}, errs.ErrRuleOptionalAppliesOnlyToObjectProperties},

			// You cannot specify children node if you use a type reference.
			{`{ // {type: "@schema"}
			"key": 123
		}`, []typ{}, errs.ErrInvalidChildNodeTogetherWithTypeReference},

			// You cannot specify other rules if you use a type reference.
			{`333 // {type: "@type", min: 1}`, []typ{}, errs.ErrCannotSpecifyOtherRulesWithTypeReference},
			{`333 // {type: "@type", min: 1}`, []typ{}, errs.ErrCannotSpecifyOtherRulesWithTypeReference},
			{`333 // {type: "@type1", type: "@type2"}`, []typ{}, errs.ErrDuplicateRule},

			// rule "or"
			{
				`true // {or: [ {type: "integer"}, {type: "string"} ]}`,
				[]typ{},
				errs.ErrIncorrectUserType,
			},
			{
				`0 // {or: [ {type: "integer", min: 1}, {type: "string"} ]}`,
				[]typ{},
				errs.ErrOrRuleSetValidation,
			},
			{`2 // {or: 123}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: "some_string"}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: "@some_string"}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: true}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: null}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: {}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},
			{`2 // {or: {or: {"@type-1","@type-2"}`, []typ{}, errs.ErrArrayWasExpectedInOrRule},

			{`2 // {or: [ 1,2,3 ]}`, []typ{}, errs.ErrIncorrectArrayItemTypeInOrRule},
			{`2 // {or: [ [],[] ]}`, []typ{}, errs.ErrIncorrectArrayItemTypeInOrRule},

			{`2 // {or: [ {type: false}, {type: "string"} ]}`, []typ{}, errs.ErrUnknownValueOfTheTypeRule},
			{`2 // {or: [ {type: "unknown_json_type"}, {type: "string"} ]}`, []typ{}, errs.ErrUnknownValueOfTheTypeRule},

			{`2 // {or: [ {type: "@type", min: 0}, {type: "string"} ]}`, []typ{}, errs.ErrCannotSpecifyOtherRulesWithTypeReference},
			{`2 // {or: [ {min: 0, type: "@type"}, {type: "string"} ]}`, []typ{}, errs.ErrCannotSpecifyOtherRulesWithTypeReference},

			{`2 // {or: [ {}, {} ]}`, []typ{}, errs.ErrEmptyRuleSet},
			{`2 // {or: []}`, []typ{}, errs.ErrEmptyArrayInOrRule},

			{`2 // {or: [ {type: "integer", min: 0} ]}`, []typ{}, errs.ErrOneElementInArrayInOrRule},
			{`2 // {or: [ {min: 0} ]}`, []typ{}, errs.ErrOneElementInArrayInOrRule},
			{`2 // {or: [ {type: "@type"} ]}`, []typ{}, errs.ErrOneElementInArrayInOrRule},
			{`2 // {or: [ "@type" ]}`, []typ{}, errs.ErrOneElementInArrayInOrRule},

			// {`2 // {or: [ {type: "integer"}, {minLength:1} ]}`, []typ{}, errors.ErrIncompatibleJsonType},
			// {`2 // {or: [ {type: "integer"}, {min:1, minLength:1} ]}`, []typ{}, errors.ErrIncompatibleJsonType},

			{`2 // {or: [ "some_string" ]}`, []typ{}, errs.ErrUnknownValueOfTheTypeRule},
			{`2 // {or: [ {type: "integer", min: 0, min: 0}, {type: "string"} ]}`, []typ{}, errs.ErrDuplicateRule},
			{`2 // {or: [ {min: []}, {type: "string"} ]}`, []typ{}, errs.ErrLiteralValueExpected},

			{`2 // {min: 1, or: [ {type: "integer"}, {type: "string"} ]}`, []typ{}, errs.ErrShouldBeNoOtherRulesInSetWithOr},
			{
				`{ // {or: [ {type: "object"}, {type: "string"} ]}
				"key": 1
			}`,
				[]typ{},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},
			{
				`[ // {or: [ {type: "array"}, {type: "string"} ]}
				1,2,3
			]`,
				[]typ{},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},

			{
				``,
				[]typ{
					{`abc`, `{}`},
				},
				errs.ErrInvalidSchemaName,
			},

			{`-5 // {or: [ {min: 0}, {type: "string"} ]}`, []typ{}, errs.ErrOrRuleSetValidation},

			// duplicate type names
			{
				``,
				[]typ{
					{`@sub1`, `"some string 1"`},
					{`@sub1`, `{}`},
				},
				errs.ErrDuplicationOfNameOfTypes,
			},

			// invalid place for comment
			{
				``,
				[]typ{
					{`@sub`, `3.4
					// {precision: 1}`},
				},
				errs.ErrIncorrectRuleWithoutExample,
			},

			// schema and example type mismatch
			{
				`[] // {or: [ {type:"string"}, {type:"@arr"} ]}`,
				[]typ{
					{`@arr`, `[1,2,3]`},
				},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},
			{
				`[] // {type: "@schema"}`,
				[]typ{
					{`@schema`, `{}`},
				},
				errs.ErrInvalidChildNodeTogetherWithTypeReference,
			},
			{
				`{} // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrInvalidChildNodeTogetherWithTypeReference,
			},
			{
				`"" // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`11 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`1.2 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`true // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`null // {type: "@schema"}`,
				[]typ{
					{`@schema`, `[1]`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`111 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `1.2`},
				},
				errs.ErrIncorrectUserType,
			},
			{ // decimal and integer
				`3 // {type: "@schema"}`,
				[]typ{
					{`@schema`, `1.2 // {precision: 1}`},
				},
				errs.ErrIncorrectUserType,
			},

			// or
			{
				`1.2  // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`true // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},
			{
				`null // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType,
			},
			{
				`{}   // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},
			{
				`[]   // {or: ["@int","@str"]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},

			{
				`1.2  // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},
			{
				`true // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},
			{
				`null // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},
			{
				`{}   // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},
			{
				`[]   // {or: [ {type:"@int"}, {type:"@str"} ]}`,
				[]typ{
					{`@int`, `111`},
					{`@str`, `"abc"`},
				},
				errs.ErrInvalidChildNodeTogetherWithOrRule,
			},

			{`1.2  // {or: [ {type:"integer"}, {type:"string"} ]}`, []typ{}, errs.ErrIncorrectUserType},
			{`true // {or: [ {type:"integer"}, {type:"string"} ]}`, []typ{}, errs.ErrIncorrectUserType},
			{`null // {or: [ {type:"integer"}, {type:"string"} ]}`, []typ{}, errs.ErrIncorrectUserType},
			{`{}   // {or: [ {type:"integer"}, {type:"string"} ]}`, []typ{}, errs.ErrIncorrectUserType},
			{`[]   // {or: [ {type:"integer"}, {type:"string"} ]}`, []typ{}, errs.ErrIncorrectUserType},

			{
				`false // {or: ["@int_or_str", "@obj"]}`,
				[]typ{
					{`@int_or_str`, `"abc" // {or: ["@int", "@str"]}`},
					{`@str`, `"abc"`},
					{`@int`, `123`},
					{`@obj`, `{}`},
				},
				errs.ErrIncorrectUserType},

			{
				`1 // {type: "@type1"}`,
				[]typ{
					{`@type1`, `1 // {type: "@type1"}`},
				},
				errs.ErrImpossibleToDetermineTheJsonTypeDueToRecursion},

			{
				`1 // {type: "@type1"}`,
				[]typ{
					{`@type1`, `1 // {type: "@type2"}`},
					{`@type2`, `2 // {type: "@type1"}`},
				},
				errs.ErrImpossibleToDetermineTheJsonTypeDueToRecursion},

			{
				`1 // {type: "@type1"}`,
				[]typ{
					{`@type1`, `1 // {type: "@type2"}`},
					{`@type2`, `2 // {type: "@type3"}`},
					{`@type3`, `3 // {type: "@type1"}`},
				},
				errs.ErrImpossibleToDetermineTheJsonTypeDueToRecursion},

			{
				`1 // {type: "@type1"}`,
				[]typ{
					{`@type1`, `1 // {or: [ {type:"integer"}, "@type2" ]}`},
					{`@type2`, `2 // {or: [ {type:"integer"}, "@type1" ]}`},
				},
				errs.ErrImpossibleToDetermineTheJsonTypeDueToRecursion},

			{
				`1 // {type: "@recurring"}`,
				[]typ{
					{`@recurring`, `"abc" // {or: ["@int", "@recurring"]}`},
					{`@int`, `123`},
				},
				errs.ErrImpossibleToDetermineTheJsonTypeDueToRecursion},

			{`"abc" // {type: "enum"}`, []typ{}, errs.ErrNotFoundRuleEnum},
			{`"abc" // {type: "enum", minLength: 1}`, []typ{}, errs.ErrNotFoundRuleEnum},

			{`"abc" // {type: "mixed"}`, []typ{}, errs.ErrNotFoundRuleOr},
			{`"abc" // {type: "mixed", minLength: 1}`, []typ{}, errs.ErrNotFoundRuleOr},

			{`2.0 // {enum: [2]}`, []typ{}, errs.ErrDoesNotMatchAnyOfTheEnumValues},
			{`2 // {enum: [2.0]}`, []typ{}, errs.ErrDoesNotMatchAnyOfTheEnumValues},

			{`"abc" // {type: "string", enum: [123, "abc"]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},
			{`"abc" // {type: "integer", enum: [123, "abc"]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},
			{`"abc" // {type: "boolean", enum: [123, "abc"]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},
			{`"abc" // {type: "string", or: [{type:"integer"}, {type:"string"}]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},

			{`2 // {type: "integer", or: [ {type: "integer"}, {type: "string"} ]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},
			{`2 // {type: "@type", or: [ {type: "integer"}, {type: "string"} ]}`, []typ{}, errs.ErrInvalidValueInTheTypeRule},

			{`"abc" // {enum: [123, "abc"], minLength: 1}`, []typ{}, errs.ErrShouldBeNoOtherRulesInSetWithEnum},
			{`"abc" // {type: "enum", enum: [123, "abc"], min: 1}`, []typ{}, errs.ErrShouldBeNoOtherRulesInSetWithEnum},

			{`123 // {type: "any", min: 1}`, []typ{}, errs.ErrShouldBeNoOtherRulesInSetWithAny},
			{`123 // {min: 1, type: "any"}`, []typ{}, errs.ErrShouldBeNoOtherRulesInSetWithAny},
			{`{ // {type: "any"}
			"aaa": 1,
			"bbb": 2
		}`, []typ{}, errs.ErrInvalidNestedElementsFoundForTypeAny},
			{`[ // {type: "any"}
			1,2,3
		]`, []typ{}, errs.ErrInvalidNestedElementsFoundForTypeAny},

			{`1 // {type: "@int"}`,
				[]typ{
					{`@int`, `1 // {type: "@str"}`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},
			{`1 // {type: "@int"}`,
				[]typ{
					{`@int`, `"abc" // {type: "@str"}`},
					{`@str`, `"abc"`},
				},
				errs.ErrIncorrectUserType},

			{`"abc"`,
				[]typ{
					{`@unused`, `-1 // {min: 0}`},
				},
				errs.ErrConstraintValidation},

			{`-5 // {or: [ {min: 0}, {type: "string"}, "@used" ]}`,
				[]typ{
					{`@used`, `0 // {min: -10}`},
					{`@unused`, `-1 // {min: 0} - incorrect EXAMPLE value`},
				},
				errs.ErrConstraintValidation},

			{`-1 // {type: "@int"}`,
				[]typ{
					{`@int`, `0 // {min: 0}`},
				},
				errs.ErrConstraintValidation},

			{`-1 // {type: "@int"}`,
				[]typ{
					{`@int`, `0 // {type: "@uint"}`},
					{`@uint`, `0  // {min: 0}`},
				},
				errs.ErrConstraintValidation},

			{`-1 // {type: "@int"}`,
				[]typ{
					{`@int`, `-1 // {type: "@uint"}`},
					{`@uint`, `0  // {min: 0}`},
				},
				errs.ErrConstraintValidation},

			{`{} // {additionalProperties: "wrong"}`, []typ{}, errs.ErrUnknownJSchemaType},

			{`"abc" // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`123 // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`123.45 // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`true // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`false // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`null // {additionalProperties: "string"}`, []typ{}, errs.ErrUnexpectedConstraint},
			{`[ // {additionalProperties: "string"}
			123
		]`, []typ{}, errs.ErrUnexpectedConstraint},

			{`{} // {allOf: 123}`, []typ{}, errs.ErrUnacceptableValueInAllOfRule},
			{`{} // {allOf: true}`, []typ{}, errs.ErrUnacceptableValueInAllOfRule},
			{`{} // {allOf: false}`, []typ{}, errs.ErrUnacceptableValueInAllOfRule},
			{`{} // {allOf: null}`, []typ{}, errs.ErrUnacceptableValueInAllOfRule},
			{`{} // {allOf: {}}`, []typ{}, errs.ErrUnacceptableValueInAllOfRule},
			{`{} // {allOf: []}`, []typ{}, errs.ErrTypeNameNotFoundInAllOfRule},
			{`{} // {allOf: "not a schema name"}`, []typ{}, errs.ErrInvalidSchemaNameInAllOfRule},
			{`{} // {allOf: ["not a schema name"]}`, []typ{}, errs.ErrInvalidSchemaNameInAllOfRule},
			{
				`{ // {allOf: "@basicError"}
				"message": "Some message text"
			}`,
				[]typ{
					{`@basicError`, `{"message": "Some message text"}`},
				},
				errs.ErrDuplicateKeysInSchema,
			},
			{
				``,
				[]typ{
					{`@aaa`, `{ // {allOf: "@bbb"}
						"aaa": "aaa"
					}`},
					{`@bbb`, `{ // {allOf: "@aaa"}
						"bbb": "bbb"
					}`},
				},
				errs.ErrUnacceptableRecursionInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `{ // {allOf: "@bbb"}
						"aaa": "aaa"
					}`},
					{`@bbb`, `{ // {allOf: "@aaa"}
						"bbb": "bbb"
					}`},
				},
				errs.ErrUnacceptableRecursionInAllOfRule,
			},
			{
				``,
				[]typ{
					{`@aaa`, `{ // {allOf: "@bbb"}
						"aaa": "aaa"
					}`},
				},
				errs.ErrTypeNotFound,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{},
				errs.ErrTypeNotFound,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `[]`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `"string"`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `123`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `123.45`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `true`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{} // {allOf: "@aaa"}`,
				[]typ{
					{`@aaa`, `null`},
				},
				errs.ErrUnacceptableUserTypeInAllOfRule,
			},
			{
				`{ // {allOf: "@aaa", additionalProperties: "integer"}
				"bbb": 222
			}`,
				[]typ{
					{`@aaa`, `{ // {additionalProperties: "string"}
						"aaa": 111
					}`},
				},
				errs.ErrConflictAdditionalProperties,
			},
			{
				`{ // {allOf: "@aaa", additionalProperties: "@int"}
				"bbb": 222
			}`,
				[]typ{
					{`@aaa`, `{ // {additionalProperties: "@str"}
						"aaa": 111
					}`},
					{`@str`, `"abc"`},
					{`@int`, `123`},
				},
				errs.ErrConflictAdditionalProperties,
			},
			{
				`[] // {minItems: 1}`,
				[]typ{},
				errs.ErrIncorrectConstraintValueForEmptyArray,
			},
			{
				`[] // {maxItems: 1}`,
				[]typ{},
				errs.ErrIncorrectConstraintValueForEmptyArray,
			},
			{
				`[] // {minItems: 1, maxItems: 1}`,
				[]typ{},
				errs.ErrIncorrectConstraintValueForEmptyArray,
			},
			{
				`{} // {type: "@sub"}`,
				[]typ{
					{`@sub`, `{
					id\n: 123
				}`},
				},
				errs.ErrInvalidChildNodeTogetherWithTypeReference,
			},
			{
				`{} // {type: "@sub"}`,
				[]typ{
					{`@sub`, `{
					id": 123
				}`},
				},
				errs.ErrInvalidChildNodeTogetherWithTypeReference,
			},
		}

		for _, tt := range tests {
			t.Run(tt.schema, func(t *testing.T) {
				defer func() {
					r := recover()
					require.NotNil(t, r, "Panic expected")

					err, ok := r.(errs.CodeKeeper)
					require.Truef(t, ok, "Unexpected error type %#v", r)

					assert.Equal(t, tt.err, err.Code())
				}()
				check(tt.schema, tt.types)
			})
		}
	})
}
