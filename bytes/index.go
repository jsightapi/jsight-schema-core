package bytes

import (
	"fmt"
	"strconv"
)

type Index uint

func (i Index) Int() int {
	return int(i)
}

func (i Index) Uint() uint {
	return uint(i)
}

func (i Index) String() string {
	return strconv.Itoa(i.Int())
}

type Indexer interface {
	int | uint | Index
}

func Int(i any) int {
	switch ii := i.(type) {
	case Index:
		return int(ii)
	case uint:
		return int(ii)
	case int:
		return ii
	}
	// This might happen only when we extend `Indexer` interface and forget
	// to add new case to the type switch above this point.
	panic(fmt.Sprintf("Unhandled content type %T", i))
}

// func NewIndex[T Indexer](i T) Index {
// 	switch ii := any(i).(type) {
// 	case Index:
// 		return ii
// 	case uint:
// 		return Index(ii)
// 	case int:
// 		return Index(ii)
// 	}
// 	// This might happen only when we extend `Indexer` interface and forget
// 	// to add new case to the type switch above this point.
// 	panic(fmt.Sprintf("Unhandled content type %T", i))
// }
