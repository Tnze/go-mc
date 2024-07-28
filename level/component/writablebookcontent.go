package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*WritableBookContent)(nil)

type WritableBookContent struct {
	Pages []Page
}

// ID implements DataComponent.
func (w *WritableBookContent) ID() string {
	return "minecraft:writable_book_content"
}

// ReadFrom implements DataComponent.
func (w *WritableBookContent) ReadFrom(reader io.Reader) (n int64, err error) {
	return pk.Array(&w.Pages).ReadFrom(reader)
}

// WriteTo implements DataComponent.
func (w *WritableBookContent) WriteTo(writer io.Writer) (n int64, err error) {
	return pk.Array(&w.Pages).WriteTo(writer)
}

type Page struct {
	Raw      pk.String
	Filtered pk.Option[pk.String, *pk.String]
}
