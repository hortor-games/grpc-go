package transport

import (
	"golang.org/x/net/http2/hpack"
	"sync"
)

var headerFieldsPool = sync.Pool{
	New: func() interface{} {
		return make([]hpack.HeaderField, 0, 20)
	},
}

func newHeaderFields() []hpack.HeaderField {
	return headerFieldsPool.Get().([]hpack.HeaderField)
}

func freeHeaderField(data []hpack.HeaderField) {
	data = data[:0]
	headerFieldsPool.Put(data)
}
