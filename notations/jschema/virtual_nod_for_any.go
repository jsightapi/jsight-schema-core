package jschema

import (
	"sync"

	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema"
)

var once sync.Once
var virtualAnyNode ischema.Node

func VirtualNodeForAny() ischema.Node {
	once.Do(func() {
		virtualAnyNode = makeVirtualNodeForAny()
	})
	return virtualAnyNode
}

func makeVirtualNodeForAny() ischema.Node {
	js := New("virtual", `"" // {type: "any"}`)

	if err := js.Check(); err != nil {
		panic(err)
	}

	return js.Inner.RootNode()
}
