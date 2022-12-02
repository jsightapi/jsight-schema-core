package test

import (
	"fmt"
	"runtime/debug"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/errs"
)

func didPanic(f assert.PanicTestFunc) (didPanic bool, message interface{}, stack string) {
	didPanic = true

	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}

func PanicsWithErr(t assert.TestingT, expected *errs.Err, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	if expected == nil {
		return assert.Fail(t, "expected should not be nil")
	}

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return assert.Fail(t, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", f, panicValue), msgAndArgs...)
	}

	actual, ok := panicValue.(*errs.Err)
	if !ok {
		return assert.Fail(t, "actual should be *errs.Err")
	}
	if actual == nil {
		return assert.Fail(t, "actual should not be nil")
	}
	if !expected.Equal(*actual) {
		return assert.Fail(
			t,
			fmt.Sprintf("func %#v should panic with value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s",
				f,
				expected,
				actual,
				panickedStack),
			msgAndArgs...,
		)
	}

	return true
}
