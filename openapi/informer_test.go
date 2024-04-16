package openapi

import (
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

func Test_Informer_RSchema(t *testing.T) {
	rSchema := &regex.RSchema{}

	informers := Dereference(rSchema)

	require.Equal(t, 1, len(informers))
	require.Equal(t, SchemaInfoTypeRegex, informers[0].Type())

	tests := []testInfoData{
		{
			`/OK/`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeRegex},
			"",
			[]testPropertiesInfos{},
		},
		// Actually JSight API Core doesn't allow annotation after regex. I left this test as an example of unexpected work.
		{
			`/OK/ // string annotation`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeRegex},
			"",
			[]testPropertiesInfos{},
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
		{
			`{ // {allOf: "@base23"} - object annotation
				"k1": "v1" // property annotation 1
			}`,
			[]testUserType{
				{
					name: "@base23",
					jsight: `{
						"k2": "v2", // {optional: true} - property annotation 2
						"k3": "v3" // property annotation 3
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"object annotation",
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
			`{ // { allOf: ["@base2","@base3"] } - object annotation
				"k1": "v1" // property annotation 1
			}`,
			[]testUserType{
				{
					name: "@base2",
					jsight: `{
						"k2": "v2" // {optional: true} - property annotation 2
					}`,
				},
				{
					name: "@base3",
					jsight: `{
						"k3": "v3" // property annotation 3
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"object annotation",
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
			`@base23 // root annotation`,
			[]testUserType{
				{
					name: "@base23",
					jsight: `{
						"k2": "v2", // {optional: true} - property annotation 2
						"k3": "v3" // property annotation 3
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"root annotation",
			[]testPropertiesInfos{
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
			`@base2 // root annotation`,
			[]testUserType{
				{
					name: "@base2",
					jsight: `{ // {allOf: "@base3"}
						"k2": "v2" // {optional: true} - property annotation 2
					}`,
				},
				{
					name: "@base3",
					jsight: `{
						"k3": "v3" // property annotation 3
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"root annotation",
			[]testPropertiesInfos{
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
			`@base1 // root annotation`,
			[]testUserType{
				{
					name: "@base1",
					jsight: `{ // {allOf: ["@base2", "@base3"]}
						"k1": "v1" // property annotation 1
					}`,
				},
				{
					name: "@base2",
					jsight: `{
						"k2": "v2" // {optional: true} - property annotation 2
					}`,
				},
				{
					name: "@base3",
					jsight: `{
						"k3": "v3" // property annotation 3
					}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
			"root annotation",
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

	require.Equal(t, len(data.expectedPropertiesInfos), expectedPropertyIndex)
}

func assertRInfo(t *testing.T, data testInfoData) {
	rSchema := buildRSchema(t, data.jsight)

	informers := Dereference(rSchema)

	node, err := rSchema.GetAST()
	require.NoError(t, err)

	assertTypes(t, data.expectedSchemaInfoTypes, informers)
	require.Equal(t, data.expectedRootAnnotation, node.Comment)
	require.Equal(t, data.expectedPropertiesInfos, []testPropertiesInfos{})
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
