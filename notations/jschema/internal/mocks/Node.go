// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	jschema "github.com/jsightapi/jsight-schema-core"
	bytes "github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/json"
	"github.com/jsightapi/jsight-schema-core/lexeme"
	constraint "github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"

	mock "github.com/stretchr/testify/mock"

	schema "github.com/jsightapi/jsight-schema-core/notations/jschema/ischema"
)

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// ASTNode provides a mock function with given fields:
func (_m *Node) ASTNode() (jschema.ASTNode, error) {
	ret := _m.Called()

	var r0 jschema.ASTNode
	if rf, ok := ret.Get(0).(func() jschema.ASTNode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(jschema.ASTNode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddConstraint provides a mock function with given fields: _a0
func (_m *Node) AddConstraint(_a0 constraint.Constraint) {
	_m.Called(_a0)
}

// BasisLexEventOfSchemaForNode provides a mock function with given fields:
func (_m *Node) BasisLexEventOfSchemaForNode() lexeme.LexEvent {
	ret := _m.Called()

	var r0 lexeme.LexEvent
	if rf, ok := ret.Get(0).(func() lexeme.LexEvent); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(lexeme.LexEvent)
	}

	return r0
}

// Comment provides a mock function with given fields:
func (_m *Node) Comment() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Constraint provides a mock function with given fields: _a0
func (_m *Node) Constraint(_a0 constraint.Type) constraint.Constraint {
	ret := _m.Called(_a0)

	var r0 constraint.Constraint
	if rf, ok := ret.Get(0).(func(constraint.Type) constraint.Constraint); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(constraint.Constraint)
		}
	}

	return r0
}

// ConstraintMap provides a mock function with given fields:
func (_m *Node) ConstraintMap() *schema.Constraints {
	ret := _m.Called()

	var r0 *schema.Constraints
	if rf, ok := ret.Get(0).(func() *schema.Constraints); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Constraints)
		}
	}

	return r0
}

// Copy provides a mock function with given fields:
func (_m *Node) Copy() schema.Node {
	ret := _m.Called()

	var r0 schema.Node
	if rf, ok := ret.Get(0).(func() schema.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.Node)
		}
	}

	return r0
}

// DeleteConstraint provides a mock function with given fields: _a0
func (_m *Node) DeleteConstraint(_a0 constraint.Type) {
	_m.Called(_a0)
}

// Grow provides a mock function with given fields: _a0
func (_m *Node) Grow(_a0 lexeme.LexEvent) (schema.Node, bool) {
	ret := _m.Called(_a0)

	var r0 schema.Node
	if rf, ok := ret.Get(0).(func(lexeme.LexEvent) schema.Node); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.Node)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(lexeme.LexEvent) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// InheritedFrom provides a mock function with given fields:
func (_m *Node) InheritedFrom() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NumberOfConstraints provides a mock function with given fields:
func (_m *Node) NumberOfConstraints() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Parent provides a mock function with given fields:
func (_m *Node) Parent() schema.Node {
	ret := _m.Called()

	var r0 schema.Node
	if rf, ok := ret.Get(0).(func() schema.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.Node)
		}
	}

	return r0
}

// RealType provides a mock function with given fields:
func (_m *Node) RealType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetComment provides a mock function with given fields: _a0
func (_m *Node) SetComment(_a0 string) {
	_m.Called(_a0)
}

// SetInheritedFrom provides a mock function with given fields: _a0
func (_m *Node) SetInheritedFrom(_a0 string) {
	_m.Called(_a0)
}

// SetParent provides a mock function with given fields: _a0
func (_m *Node) SetParent(_a0 schema.Node) {
	_m.Called(_a0)
}

// SetRealType provides a mock function with given fields: _a0
func (_m *Node) SetRealType(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *Node) Type() json.Type {
	ret := _m.Called()

	var r0 json.Type
	if rf, ok := ret.Get(0).(func() json.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(json.Type)
	}

	return r0
}

// Value provides a mock function with given fields:
func (_m *Node) Value() bytes.Bytes {
	ret := _m.Called()

	var r0 bytes.Bytes
	if rf, ok := ret.Get(0).(func() bytes.Bytes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.Bytes)
		}
	}

	return r0
}

type NewNodeT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNode creates a new instance of Node. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNode(t NewNodeT) *Node {
	mock := &Node{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
