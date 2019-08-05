package region

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"time"
)

// Region contain 32*32 chunks in one .mca file
type Region struct {
	f          *os.File
	offsets    [32][32]int32
	timestamps [32][32]int32

	// sectors record if a sector is in used.
	// contrary to mojang's, because false is the default value in Go.
	sectors map[int32]bool
}

// In calculate chunk's coordinates relative to region
func In(cx, cy int) (int, int) {
	// c & (32-1)
	// is equal to:
	// (c %= 32) > 0 ? c : -c; //C language
	return cx & 31, cy & 31
}

// Open open a .mca file and read the head.
// Close the Region after used.
func Open(name string) (r *Region, err error) {
	r = new(Region)
	r.sectors = make(map[int32]bool)

	r.f, err = os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	// read the offsets
	err = binary.Read(r.f, binary.BigEndian, &r.offsets)
	if err != nil {
		_ = r.f.Close()
		return nil, err
	}
	r.sectors[0] = true

	// read the timestamps
	err = binary.Read(r.f, binary.BigEndian, &r.timestamps)
	if err != nil {
		_ = r.f.Close()
		return nil, err
	}
	r.sectors[1] = true

	// generate sectorFree table
	for _, v := range r.offsets {
		for _, v := range v {
			if o, s := sectorLoc(v); o != 0 {
				for i := int32(0); i < s; i++ {
					r.sectors[o+i] = true
				}
			}
		}
	}

	return r, nil
}

// Create open .mca file with os.O_CREATE|os. O_EXCL, and init the region
func Create(name string) (r *Region, err error) {
	r = new(Region)
	r.sectors = make(map[int32]bool)

	r.f, err = os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
	if err != nil {
		return nil, err
	}

	// read the offsets
	err = binary.Write(r.f, binary.BigEndian, &r.offsets)
	if err != nil {
		_ = r.f.Close()
		return nil, err
	}
	r.sectors[0] = true

	// read the timestamps
	err = binary.Write(r.f, binary.BigEndian, &r.timestamps)
	if err != nil {
		_ = r.f.Close()
		return nil, err
	}
	r.sectors[1] = true

	return r, nil
}

// Close close the region file
func (r *Region) Close() error {
	return r.f.Close()
}

func sectorLoc(offset int32) (o, s int32) {
	return offset >> 8, offset & 0xFF
}

// ReadSector find and read the Chunk data from region
func (r *Region) ReadSector(x, y int) (data []byte, err error) {
	offset, _ := sectorLoc(r.offsets[x][y])

	if offset == 0 {
		return nil, errors.New("sector not exist")
	}

	_, err = r.f.Seek(4096*int64(offset), 0)
	if err != nil {
		return
	}

	var length int32
	err = binary.Read(r.f, binary.BigEndian, &length)
	if err != nil {
		return
	}

	data = make([]byte, length)
	_, err = io.ReadFull(r.f, data)

	return
}

// WriteSector write Chunk data into region file
func (r *Region) WriteSector(x, y int, data []byte) error {
	need := int32(len(data)+4)/4096 + 1
	n, now := sectorLoc(r.offsets[x][y])

	// maximum chunk size is 1MB
	if need >= 256 {
		return errors.New("data too large")
	}

	if n != 0 && now == need {
		// we can simply overwrite the old sectors
	} else {
		// we need to allocate new sectors

		// mark the sectors previously used for this chunk as free
		for i := int32(0); i < now; i++ {
			r.sectors[n+i] = false
		}

		// scan for a free space large enough to store this chunk
		n = r.findSpace(need)

		// mark the sectors previously used for this chunk as used
		now = need
		for i := int32(0); i < need; i++ {
			r.sectors[n+i] = true
		}

		r.offsets[x][y] = (n << 8) | (need & 0xFF)

		// update file head
		err := r.setHead(x, y, uint32(r.offsets[x][y]), uint32(time.Now().Unix()))
		if err != nil {
			return err
		}
	}

	_, err := r.f.Seek(4096*int64(n), 0)
	if err != nil {
		return err
	}
	//data length
	err = binary.Write(r.f, binary.BigEndian, int32(len(data)))
	if err != nil {
		return err
	}

	//data
	_, err = r.f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// ExistSector return if a sector is exist
func (r *Region) ExistSector(x, y int) bool {
	return r.offsets[x][y] != 0
}

func (r *Region) findSpace(need int32) (n int32) {
	for i := int32(0); i < need; i++ {
		if r.sectors[n+i] {
			n += i + 1
			i = -1
		}
	}
	return
}

func (r *Region) setHead(x, y int, offset, timestamp uint32) (err error) {
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:], offset)
	_, err = r.f.WriteAt(buf[:], 4*(int64(x)*32+int64(y)))
	if err != nil {
		return
	}

	binary.BigEndian.PutUint32(buf[:], timestamp)
	_, err = r.f.WriteAt(buf[:], 4096+4*(int64(x)*32+int64(y)))
	if err != nil {
		return
	}

	return
}
