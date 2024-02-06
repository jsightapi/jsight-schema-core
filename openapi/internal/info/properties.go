package info

import "github.com/jsightapi/jsight-schema-core/errs"

type Properties struct {
	i     int
	keys  []string
	value []Info
}

func newProperties(capacity int) *Properties {
	return &Properties{
		i:     0,
		keys:  make([]string, 0, capacity),
		value: make([]Info, 0, capacity),
	}
}

func (p *Properties) append(k string, v Info) {
	p.keys = append(p.keys, k)
	p.value = append(p.value, v)
}

func (p Properties) has() bool {
	return p.i >= 0 && p.i <= len(p.keys)
}

func (p *Properties) Rewind() {
	p.i = 0
}

func (p *Properties) Next() bool {
	p.i++
	return p.has()
}

func (p Properties) GetKey() string {
	if !p.has() {
		panic(errs.ErrRuntimeFailure.F())
	}

	return p.keys[p.i-1]
}

func (p Properties) GetInfo() Info {
	if !p.has() {
		panic(errs.ErrRuntimeFailure.F())
	}

	return p.value[p.i-1]
}
