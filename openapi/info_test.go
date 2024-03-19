package openapi

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	"github.com/jsightapi/jsight-schema-core/openapi/info"
)

type testInfoData struct {
	jsight                  string
	userTypes               []testUserType
	expectedElementTypes    []info.ElementType
	expectedRootAnnotation  string
	expectedPropertiesInfos []testPropertiesInfos
}

type testUserType struct {
	name   string
	jsight string
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

func Test_Info_RSchema(t *testing.T) {
	rSchema := &regex.RSchema{}

	elements := Dereference(rSchema)

	require.Equal(t, 1, len(elements))
	require.Equal(t, info.ElementTypeRegex, elements[0].Type())
}

func Test_Info_JSchema(t *testing.T) {
	tests := []testInfoData{
		{
			`"abc" // string annotation`,
			[]testUserType{},
			[]info.ElementType{info.ElementTypeScalar},
			"string annotation",
			[]testPropertiesInfos{},
		},
		{
			`[] // array annotation`,
			[]testUserType{},
			[]info.ElementType{info.ElementTypeArray},
			"array annotation",
			[]testPropertiesInfos{},
		},
		{
			`123 // {type: "any"} - array annotation`,
			[]testUserType{},
			[]info.ElementType{info.ElementTypeAny},
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
			[]info.ElementType{info.ElementTypeObject},
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
			[]info.ElementType{info.ElementTypeObject},
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
			[]info.ElementType{info.ElementTypeObject, info.ElementTypeObject},
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
			[]info.ElementType{info.ElementTypeObject, info.ElementTypeObject},
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
			[]info.ElementType{info.ElementTypeObject, info.ElementTypeScalar},
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

func buildJSchema(t *testing.T, jsight string, userTypes []testUserType) *jschema.JSchema {
	jSchema := jschema.New("root", jsight)

	for _, ut := range userTypes {
		err := jSchema.AddType(ut.name, jschema.New(ut.name, ut.jsight))
		require.NoError(t, err)
	}

	err := jSchema.Check()
	require.NoError(t, err)

	return jSchema
}

func assertInfo(t *testing.T, data testInfoData) {
	jSchema := buildJSchema(t, data.jsight, data.userTypes)

	elements := Dereference(jSchema)

	require.Equal(t, data.expectedRootAnnotation, jSchema.ASTNode.Comment) // TODO need an interface?

	assertTypes(t, data.expectedElementTypes, elements)

	expectedPropertyIndex := 0

	for _, ei := range elements {
		if ei.Type() == info.ElementTypeObject {
			oi := newObjectInfoImpl(ei)
			properties := oi.PropertiesInfos()

			for _, pi := range properties {
				require.True(t, expectedPropertyIndex < len(data.expectedPropertiesInfos))

				assertProperty(t, data.expectedPropertiesInfos[expectedPropertyIndex], pi)

				expectedPropertyIndex++
			}
		}
	}
}

func assertTypes(t *testing.T, expected []info.ElementType, elements []ElementInfo) {
	actual := make([]info.ElementType, len(elements))
	for i, ei := range elements {
		actual[i] = ei.Type()
	}
	require.Equal(t, expected, actual)
}

func assertProperty(t *testing.T, expected testPropertiesInfos, pi PropertyInfo) {
	require.Equal(t, expected.key, pi.Key())
	require.Equal(t, expected.optional, pi.Optional())
	require.Equal(t, expected.annotation, pi.Annotation())

	actual, err := pi.SchemaObject().MarshalJSON()
	require.NoError(t, err)
	require.JSONEq(t, expected.openapi, string(actual))
}
