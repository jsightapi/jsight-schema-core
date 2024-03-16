package info

import "sync"

type ReferenceTargetInfoList struct {
	mu  sync.Mutex
	lst []ReferenceTargetInfo
}

func newReferenceTargetInfoList() ReferenceTargetInfoList {
	return ReferenceTargetInfoList{
		lst: make([]ReferenceTargetInfo, 0, 3),
	}
}

func (l *ReferenceTargetInfoList) append(r ReferenceTargetInfo) *ReferenceTargetInfo {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.lst = append(l.lst, r)
	return &(l.lst[len(l.lst)-1])
}

func (l *ReferenceTargetInfoList) list() []ReferenceTargetInfo {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.lst
}
