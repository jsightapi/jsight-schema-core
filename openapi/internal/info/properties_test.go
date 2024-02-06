package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProperties_has(t *testing.T) {
	p := newProperties(2)
	p.append("a", Info{})
	p.append("b", Info{})
	p.append("c", Info{}) // over capacity

	assert.True(t, p.has())

	p.Next()
	assert.True(t, p.has())

	p.Next()
	assert.True(t, p.has())

	p.Next()
	assert.True(t, p.has())

	p.Next()
	assert.False(t, p.has())
}

func TestProperties_Rewind(t *testing.T) {
	p := newProperties(2)
	p.append("a", Info{})
	p.append("b", Info{})
	p.append("c", Info{}) // over capacity

	p.Next()
	p.Next()
	p.Next()
	assert.True(t, p.has())
	p.Next()
	assert.False(t, p.has())

	p.Rewind()
	assert.True(t, p.has())
}

func TestProperties_GetKey(t *testing.T) {
	p := newProperties(2)
	p.append("a", Info{})
	p.append("b", Info{})
	p.append("c", Info{}) // over capacity

	assert.Panics(t, func() { p.GetKey() }) // waiting for Next

	p.Next()
	assert.Equal(t, "a", p.GetKey())

	p.Next()
	assert.Equal(t, "b", p.GetKey())

	p.Next()
	assert.Equal(t, "c", p.GetKey())

	p.Next()
	assert.Panics(t, func() { p.GetKey() })
}
