package arraylist

import "github.com/454270186/GoDazzle/list"

type Option struct {
	ThreadSafe bool
}

func New(opts ...Option) list.List {
	for _, opt := range opts {
		if opt.ThreadSafe {
			return newTFArrList()
		}
	}

	return newArrayList()
}