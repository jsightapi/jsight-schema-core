// Code generated by mockery v2.12.3. DO NOT EDIT.

package ischema

import (
	jschema "github.com/jsightapi/jsight-schema-core"
	bytes "github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/json"
	"github.com/jsightapi/jsight-schema-core/lexeme"
	constraint "github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"

	mock "github.com/stretchr/testify/mock"
)

// MockNode is an autogenerated mock type for the Node type
type MockNode struct {
	mock.Mock
}

// ASTNode provides a mock function with given fields:
func (_m *MockNode) ASTNode() (jschema.ASTNode, error) {
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
func (_m *MockNode) AddConstraint(_a0 constraint.Constraint) {
	_m.Called(_a0)
}

// BasisLexEventOfSchemaForNode provides a mock function with given fields:
func (_m *MockNode) BasisLexEventOfSchemaForNode() lexeme.LexEvent {
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
func (_m *MockNode) Comment() string {
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
func (_m *MockNode) Constraint(_a0 constraint.Type) constraint.Constraint {
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
func (_m *MockNode) ConstraintMap() *Constraints {
	ret := _m.Called()

	var r0 *Constraints
	if rf, ok := ret.Get(0).(func() *Constraints); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Constraints)
		}
	}

	return r0
}

// Copy provides a mock function with given fields:
func (_m *MockNode) Copy() Node {
	ret := _m.Called()

	var r0 Node
	if rf, ok := ret.Get(0).(func() Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Node)
		}
	}

	return r0
}

// DeleteConstraint provides a mock function with given fields: _a0
func (_m *MockNode) DeleteConstraint(_a0 constraint.Type) {
	_m.Called(_a0)
}

// Grow provides a mock function with given fields: _a0
func (_m *MockNode) Grow(_a0 lexeme.LexEvent) (Node, bool) {
	ret := _m.Called(_a0)

	var r0 Node
	if rf, ok := ret.Get(0).(func(lexeme.LexEvent) Node); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Node)
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
func (_m *MockNode) InheritedFrom() string {
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
func (_m *MockNode) NumberOfConstraints() int {
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
func (_m *MockNode) Parent() Node {
	ret := _m.Called()

	var r0 Node
	if rf, ok := ret.Get(0).(func() Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Node)
		}
	}

	return r0
}

// RealType provides a mock function with given fields:
func (_m *MockNode) RealType() string {
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
func (_m *MockNode) SetComment(_a0 string) {
	_m.Called(_a0)
}

// SetInheritedFrom provides a mock function with given fields: _a0
func (_m *MockNode) SetInheritedFrom(_a0 string) {
	_m.Called(_a0)
}

// SetParent provides a mock function with given fields: _a0
func (_m *MockNode) SetParent(_a0 Node) {
	_m.Called(_a0)
}

// SetRealType provides a mock function with given fields: _a0
func (_m *MockNode) SetRealType(_a0 string) bool {
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
func (_m *MockNode) Type() json.Type {
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
func (_m *MockNode) Value() bytes.Bytes {
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

type NewMockNodeT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockNode creates a new instance of MockNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockNode(t NewMockNodeT) *MockNode {
	mock := &MockNode{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
