// Autogenerated code!
// DO NOT EDIT!
//
// Generated by OrderedMap generator from the internal/cmd/generator command.

package schema

import (
	"bytes"
	"encoding/json"
)

// Set sets a value with specified key.
func (m *RuleASTNodes) Set(k string, v RuleASTNode) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.data == nil {
		m.data = map[string]RuleASTNode{}
	}
	if !m.has(k) {
		m.order = append(m.order, k)
	}
	m.data[k] = v
}

// Update updates a value with specified key.
func (m *RuleASTNodes) Update(k string, fn func(v RuleASTNode) RuleASTNode) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if !m.has(k) {
		// Prevent from possible nil pointer dereference if map value type is a
		// pointer.
		return
	}

	m.data[k] = fn(m.data[k])
}

// GetValue gets a value by key.
func (m *RuleASTNodes) GetValue(k string) RuleASTNode {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.data[k]
}

// Get gets a value by key.
func (m *RuleASTNodes) Get(k string) (RuleASTNode, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	v, ok := m.data[k]
	return v, ok
}

// Has checks that specified key is set.
func (m *RuleASTNodes) Has(k string) bool {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.has(k)
}

func (m *RuleASTNodes) has(k string) bool {
	_, ok := m.data[k]
	return ok
}

// Len returns count of values.
func (m *RuleASTNodes) Len() int {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return len(m.data)
}

func (m *RuleASTNodes) Delete(k string) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.delete(k)
}

func (m *RuleASTNodes) delete(k string) {
	var kk string
	i := -1

	for i, kk = range m.order {
		if kk == k {
			break
		}
	}

	delete(m.data, k)
	if i != -1 {
		m.order = append(m.order[:i], m.order[i+1:]...)
	}
}

// Filter iterates and changes values in the map.
func (m *RuleASTNodes) Filter(fn filterRuleASTNodesFunc) {
	m.mx.Lock()
	defer m.mx.Unlock()

	for _, k := range m.order {
		if !fn(k, m.data[k]) {
			m.delete(k)
		}
	}
}

type filterRuleASTNodesFunc = func(k string, v RuleASTNode) bool

// Find finds first matched item from the map.
func (m *RuleASTNodes) Find(fn findRuleASTNodesFunc) (RuleASTNodesItem, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	for _, k := range m.order {
		if fn(k, m.data[k]) {
			return RuleASTNodesItem{
				Key:   k,
				Value: m.data[k],
			}, true
		}
	}
	return RuleASTNodesItem{}, false
}

type findRuleASTNodesFunc = func(k string, v RuleASTNode) bool

func (m *RuleASTNodes) Each(fn eachRuleASTNodesFunc) error {
	m.mx.RLock()
	defer m.mx.RUnlock()

	for _, k := range m.order {
		if err := fn(k, m.data[k]); err != nil {
			return err
		}
	}
	return nil
}

type eachRuleASTNodesFunc = func(k string, v RuleASTNode) error

func (m *RuleASTNodes) EachSafe(fn eachSafeRuleASTNodesFunc) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	for _, k := range m.order {
		fn(k, m.data[k])
	}
}

type eachSafeRuleASTNodesFunc = func(k string, v RuleASTNode)

// Map iterates and changes values in the map.
func (m *RuleASTNodes) Map(fn mapRuleASTNodesFunc) error {
	m.mx.Lock()
	defer m.mx.Unlock()

	for _, k := range m.order {
		v, err := fn(k, m.data[k])
		if err != nil {
			return err
		}
		m.data[k] = v
	}
	return nil
}

type mapRuleASTNodesFunc = func(k string, v RuleASTNode) (RuleASTNode, error)

// RuleASTNodesItem represent single data from the RuleASTNodes.
type RuleASTNodesItem struct {
	Key   string
	Value RuleASTNode
}

var _ json.Marshaler = &RuleASTNodes{}

func (m *RuleASTNodes) MarshalJSON() ([]byte, error) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	var buf bytes.Buffer
	buf.WriteRune('{')

	for i, k := range m.order {
		if i != 0 {
			buf.WriteRune(',')
		}

		// marshal key
		key, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteRune(':')

		// marshal value
		val, err := json.Marshal(m.data[k])
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteRune('}')
	return buf.Bytes(), nil
}
