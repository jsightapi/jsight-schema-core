package info

type elementInfoList struct {
	lst []ElementInfo
}

func newElementInfoList() *elementInfoList {
	return &elementInfoList{
		lst: make([]ElementInfo, 0, 5),
	}
}

func (l *elementInfoList) append(r ElementInfo) {
	l.lst = append(l.lst, r)
}

func (l elementInfoList) list() []ElementInfo {
	return l.lst
}
