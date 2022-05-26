package world

import (
	_ "embed"
	"github.com/Tnze/go-mc/server"
	"io"
	"unsafe"

	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server/ecs"
)

//go:embed DimensionCodec.snbt
var dimensionCodecSNBT nbt.StringifiedMessage

//go:embed Dimension.snbt
var dimensionSNBT nbt.StringifiedMessage

type Dimension struct {
	storage
	Name string
}

func NewDimension(name, path string) Dimension {
	return Dimension{
		Name:    name,
		storage: storage{regionDir: path},
	}
}

type DimensionList struct {
	Dims                  []ecs.Index
	DimNames              []string
	DimCodecSNBT, DimSNBT nbt.StringifiedMessage
}

func (d *DimensionList) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Array(*(*[]pk.Identifier)(unsafe.Pointer(&d.DimNames))).WriteTo(w)
}

func (d *DimensionList) Find(dim string) (ecs.Index, bool) {
	for i, v := range d.DimNames {
		if v == dim {
			return d.Dims[i], true
		}
	}
	return 0, false
}

func (d *DimensionList) Add(id ecs.Index, name string) {
	d.Dims = append(d.Dims, id)
	d.DimNames = append(d.DimNames, name)
}

func NewDimensionManager(g *server.Game) *DimensionList {
	return ecs.SetResource(g.World, DimensionList{
		Dims:         nil,
		DimNames:     nil,
		DimCodecSNBT: dimensionCodecSNBT,
		DimSNBT:      dimensionSNBT,
	})
}
