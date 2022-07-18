package packet

import (
	"bytes"
	"compress/zlib"
	"io"
	"sync"
	// "github.com/klauspost/compress/zlib"
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

type ZlibWriterPool sync.Pool

func (pl *ZlibWriterPool) Get(w io.Writer) *zlib.Writer {
	zw := ((*sync.Pool)(pl)).Get().(*zlib.Writer)
	zw.Reset(w)
	return zw
}

func (pl *ZlibWriterPool) Return(zw *zlib.Writer) {
	((*sync.Pool)(pl)).Put(zw)
}

func NewZlibWriterPool() *ZlibWriterPool {
	return (*ZlibWriterPool)(&sync.Pool{
		New: func() interface{} {
			// return new(zlib.Writer)
			zw, _ := zlib.NewWriterLevel(nil, zlib.BestSpeed)
			return zw
		},
	})
}

var zlibWriterPool = NewZlibWriterPool()
