// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	jschema "github.com/jsightapi/jsight-schema-core"
	mock "github.com/stretchr/testify/mock"
)

// Rule is an autogenerated mock type for the Rule type
type Rule struct {
	mock.Mock
}

// Check provides a mock function with given fields:
func (_m *Rule) Check() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAST provides a mock function with given fields:
func (_m *Rule) GetAST() (jschema.ASTNode, error) {
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

// Len provides a mock function with given fields:
func (_m *Rule) Len() (uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRuleT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRule creates a new instance of Rule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRule(t NewRuleT) *Rule {
	mock := &Rule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
