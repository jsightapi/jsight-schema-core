package openapi

import (
	"github.com/stretchr/testify/assert"

	"regexp"

	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/regex"
)

type testInfoData struct {
	jsight                  string
	userTypes               []testUserType
	expectedSchemaInfoTypes []SchemaInfoType
	expectedRootAnnotation  string
	expectedPropertiesInfos []testPropertiesInfos
}

type testPropertiesInfos struct {
	optional   bool
	key        string
	annotation string
	openapi    string
}

func (t testInfoData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

func (t testRInfoData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

type testRInfoData struct {
	jsight                  string
	expectedSchemaInfoTypes []SchemaInfoType
}

func Test_Informer_RSchema(t *testing.T) {
	rSchema := &regex.RSchema{}

	informers := Dereference(rSchema)

	require.Equal(t, 1, len(informers))
	require.Equal(t, SchemaInfoTypeRegex, informers[0].Type())

	tests := []testRInfoData{
		{
			`/OK/`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
		{
			`/ /`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
		{
			`/^[A-Z][a-z]*( [A-Z][a-z]*)*$/`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
		{
			`/^[a-zA-Z0-9.!#$%&'*+\/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
		{
			`/(?:[a-z0-9!#$%&'*+\\\/=?^_` + "`" + `{|}~-])+(?:\\.[a-z0-9!#$%&'*+\\\/=?^_` + "`" + `{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e\-\\x1f\\x21\\x23-\\x5b\\x5d\-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e\-\\x7f])*\"\)@\(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\)\\.\){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e\-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e\-\\x7f])+)\\]\)/`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
		{
			`/([^][()<>@,;:\\". \x00-\x1F\x7F]+|"(\n|(\\\r)*([^"\\\r\n]|\\[^\r]))*(\\\r)*")(\.([^][()<>@,;:\\". \x00-\x1F\x7F]+|"(\n|(\\\r)*([^"\\\r\n]|\\[^\r]))*(\\\r)*"))*@([^][()<>@,;:\\". \x00-\x1F\x7F]+|\[(\n|(\\\r)*([^][\\\r\n]|\\[^\r]))*(\\\r)*])(\.([^][()<>@,;:\\". \x00-\x1F\x7F]+|\[(\n|(\\\r)*([^][\\\r\n]|\\[^\r]))*(\\\r)*]))*/`,
			[]SchemaInfoType{SchemaInfoTypeRegex},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertRInfo(t, data)
		})
	}
}

func Test_Informer_JSchema(t *testing.T) {
	tests := []testInfoData{
		{
			`"abc" // string annotation`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar},
			"string annotation",
			[]testPropertiesInfos{},
		},
		{
			`[] // array annotation`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeArray},
			"array annotation",
			[]testPropertiesInfos{},
		},
		{
			`123 // {type: "any"} - array annotation`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeAny},
			"array annotation",
			[]testPropertiesInfos{},
		},
		{
			`{ // object annotation
				"k1": "v1", // {optional: true} - property annotation 1
				"k2": "v2", // {optional: false} - property annotation 2
				"k3": "v3" // property annotation 3
			}`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"object annotation",
			[]testPropertiesInfos{
				{
					optional:   true,
					key:        "k1",
					annotation: "property annotation 1",
					openapi: `{
						"type": "string",
						"example": "v1",
						"description": "property annotation 1"
					}`,
				},
				{
					optional:   false,
					key:        "k2",
					annotation: "property annotation 2",
					openapi: `{
						"type": "string",
						"example": "v2",
						"description": "property annotation 2"
					}`,
				},
				{
					optional:   false,
					key:        "k3",
					annotation: "property annotation 3",
					openapi: `{
						"type": "string",
						"example": "v3",
						"description": "property annotation 3"
					}`,
				},
			},
		},
		{
			`@refToObj1 // first reference annotation`,
			[]testUserType{
				{
					name:   "@refToObj1",
					jsight: `@refToObj2 // second reference annotation`,
				},
				{
					name: "@refToObj2",
					jsight: `{ // object annotation
						"k1": "v1", // property annotation 1
						"k2": "v2"  // {optional: true} - property annotation 2
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"first reference annotation",
			[]testPropertiesInfos{
				{
					optional:   false,
					key:        "k1",
					annotation: "property annotation 1",
					openapi: `{
						"type": "string",
						"example": "v1",
						"description": "property annotation 1"
					}`,
				},
				{
					optional:   true,
					key:        "k2",
					annotation: "property annotation 2",
					openapi: `{
						"type": "string",
						"example": "v2",
						"description": "property annotation 2"
					}`,
				},
				{
					optional:   false,
					annotation: "property annotation 3",
					openapi: `{
						"type": "string",
						"example": "v3",
						"description": "property annotation 3"
					}`,
				},
			},
		},
		{
			`@obj1 | @obj2 // or annotation`,
			[]testUserType{
				{
					name:   "@obj1",
					jsight: `{"k1": "v1"}`,
				},
				{
					name: "@obj2",
					jsight: `{
						"k2": "v2" // property annotation 2
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeObject},
			"or annotation",
			[]testPropertiesInfos{
				{
					optional:   false,
					key:        "k1",
					annotation: "",
					openapi: `{
						"type": "string",
						"example": "v1"
					}`,
				},
				{
					optional:   false,
					key:        "k2",
					annotation: "property annotation 2",
					openapi: `{
						"type": "string",
						"example": "v2",
						"description": "property annotation 2"
					}`,
				},
			},
		},
		{
			`@obj1 | @refToObj2 // or annotation`,
			[]testUserType{
				{
					name:   "@obj1",
					jsight: `{"k1": "v1"}`,
				},
				{
					name:   "@refToObj2",
					jsight: `@obj2`,
				},
				{
					name: "@obj2",
					jsight: `{
						"k2": "v2" // property annotation 2
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeObject},
			"or annotation",
			[]testPropertiesInfos{
				{
					optional:   false,
					key:        "k1",
					annotation: "",
					openapi: `{
						"type": "string",
						"example": "v1"
					}`,
				},
				{
					optional:   false,
					key:        "k2",
					annotation: "property annotation 2",
					openapi: `{
						"type": "string",
						"example": "v2",
						"description": "property annotation 2"
					}`,
				},
			},
		},
		{
			`@obj1 | @refToString // or annotation`,
			[]testUserType{
				{
					name:   "@obj1",
					jsight: `{"k1": "v1"}`,
				},
				{
					name:   "@refToString",
					jsight: `@str`,
				},
				{
					name:   "@str",
					jsight: `"abc" // string annotation`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeScalar},
			"or annotation",
			[]testPropertiesInfos{
				{
					optional:   false,
					key:        "k1",
					annotation: "",
					openapi: `{
						"type": "string",
						"example": "v1"
					}`,
				},
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertInfo(t, data)
		})
	}
}

func assertInfo(t *testing.T, data testInfoData) {
	jSchema := buildJSchema(t, data.jsight, data.userTypes)

	informers := Dereference(jSchema)

	require.Equal(t, data.expectedRootAnnotation, jSchema.ASTNode.Comment) // TODO need an interface?

	assertTypes(t, data.expectedSchemaInfoTypes, informers)

	expectedPropertyIndex := 0

	for _, ei := range informers {
		if ei.Type() == SchemaInfoTypeObject {
			properties := ei.(ObjectInformer).PropertiesInfos()

			for _, pi := range properties {
				require.True(t, expectedPropertyIndex < len(data.expectedPropertiesInfos))

				assertProperty(t, data.expectedPropertiesInfos[expectedPropertyIndex], pi)

				expectedPropertyIndex++
			}
		}
	}
}

func assertRInfo(t *testing.T, data testRInfoData) {
	rSchema := buildRSchema(t, data.jsight)

	informers := Dereference(rSchema)
	info := NewRSchemaInfo(rSchema)

	assert.Equal(t, SchemaInfoTypeRegex, info.Type())
	assertTypes(t, data.expectedSchemaInfoTypes, informers)
}

func assertTypes(t *testing.T, expected []SchemaInfoType, informers []SchemaInformer) {
	actual := make([]SchemaInfoType, len(informers))
	for i, ei := range informers {
		actual[i] = ei.Type()
	}
	require.Equal(t, expected, actual)
}

func assertProperty(t *testing.T, expected testPropertiesInfos, pi PropertyInformer) {
	require.Equal(t, expected.key, pi.Key())
	require.Equal(t, expected.optional, pi.Optional())
	require.Equal(t, expected.annotation, pi.Annotation())

	actual, err := pi.SchemaObject().MarshalJSON()
	require.NoError(t, err)
	require.JSONEq(t, expected.openapi, string(actual))
}
