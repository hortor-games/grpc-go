package transport

import (
	"runtime"
	"sync"
)

type AutoPool struct {
	sync.Pool
	Free func(data interface{})
}

func NewAutoPool(new func() interface{}, free func(data interface{})) *AutoPool {
	return &AutoPool{
		Pool: sync.Pool{New: new},
		Free: free,
	}
}

func (p *AutoPool) Get() interface{} {
	v := p.Pool.Get()
	runtime.SetFinalizer(v, func(data interface{}) {
		p.Free(data)
		p.Put(data)
	})
	return v
}
