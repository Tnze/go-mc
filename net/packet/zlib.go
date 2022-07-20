package packet

import (
	// "compress/zlib" // should use BestSpeed(level 1)
	"io"

	"github.com/klauspost/compress/zlib"
)

type ZlibWriter zlib.Writer

func (zw *ZlibWriter) Reset(w io.Writer) {
	(*zlib.Writer)(zw).Reset(w)
}
func (zw *ZlibWriter) Write(p []byte) (n int, err error) {
	return (*zlib.Writer)(zw).Write(p)
}
func (zw *ZlibWriter) Flush() error {
	return (*zlib.Writer)(zw).Flush()
}
func (zw *ZlibWriter) Close() error {
	return (*zlib.Writer)(zw).Close()
}

func NewZlibWriterLevel(level int) *ZlibWriter {
	zw, _ := zlib.NewWriterLevel(nil, level)
	return (*ZlibWriter)(zw)
}

func NewZlibReader(r io.Reader) (io.ReadCloser, error) {
	return zlib.NewReader(r)
}
