package openapi

type elementInfoList struct {
	lst []ElementInformer
}

func newElementInfoList() *elementInfoList {
	return &elementInfoList{
		lst: make([]ElementInformer, 0, 5),
	}
}

func (l *elementInfoList) append(r ElementInformer) {
	l.lst = append(l.lst, r)
}

func (l elementInfoList) list() []ElementInformer {
	return l.lst
}
