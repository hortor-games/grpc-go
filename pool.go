package grpc

import "sync"

var bytesPools = []sync.Pool{
	{
		New: func() interface{} {
			return make([]byte, 0, 1024)
		},
	},
	{
		New: func() interface{} {
			return make([]byte, 0, 10240)
		},
	},
	{
		New: func() interface{} {
			return make([]byte, 0, 102400)
		},
	},
	{
		New: func() interface{} {
			return []byte(nil)
		},
	},
}

func newBytes(length int) []byte {
	if length <= 1024 {
		return bytesPools[0].Get().([]byte)[:length]
	}
	if length <= 10240 {
		return bytesPools[1].Get().([]byte)[:length]
	}
	if length <= 102400 {
		return bytesPools[2].Get().([]byte)[:length]
	}
	bytes := bytesPools[3].Get().([]byte)
	if cap(bytes) < length {
		bytesPools[3].Put(bytes)
		bytes = make([]byte, length)
	}
	return bytes
}

func freeBytes(bytes []byte) {
	c := cap(bytes)
	if c < 1024 || c > 1024*1024*4 {
		return
	}
	if c > 102400 {
		bytesPools[3].Put(bytes[:0])
		return
	}
	if c == 102400 {
		bytesPools[2].Put(bytes[:0])
		return
	}
	if c >= 10240 {
		bytesPools[1].Put(bytes[:0])
		return
	}
	if c >= 1024 {
		bytesPools[0].Put(bytes[:0])
		return
	}
}
