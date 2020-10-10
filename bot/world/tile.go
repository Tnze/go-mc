package world

import (
  "fmt"
)

// TilePosition describes the location of a tile/block entity within a chunk.
type TilePosition uint32

func (p TilePosition) Pos() (x, y, z int) {
  return int((p>>8) & 0xff), int((p>>16) & 0xff), int(p&0xff)
}

func (p TilePosition) String() string {
  x, y, z := p.Pos()
  return fmt.Sprintf("(%d, %d, %d)", x, y, z)
}

func ToTilePos(x, y, z int) TilePosition {
  return TilePosition((y&0xff) << 16 | (x&15) << 8 | (z&15))
}
