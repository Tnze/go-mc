package packet

import (
	"bytes"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type BuffReaderPool sync.Pool

func (pl *BuffReaderPool) Get(b []byte) *bytes.Reader {
	br := ((*sync.Pool)(pl)).Get().(*bytes.Reader)
	br.Reset(b)
	return br
}

func (pl *BuffReaderPool) Return(b *bytes.Reader) {
	((*sync.Pool)(pl)).Put(b)
}

func NewBuffReaderPool() *BuffReaderPool {
	return (*BuffReaderPool)(&sync.Pool{
		New: func() interface{} {
			return new(bytes.Reader)
		},
	})
}

var buffReaderPool = NewBuffReaderPool()
