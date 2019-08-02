package region

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

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

// OpenRegion open a .mca file and read the head.
// Close the Region after used.
func OpenRegion(name string) (r *Region, err error) {
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

func (r *Region) Close() error {
	_, err := r.f.Seek(0, 0)
	if err != nil {
		return err
	}

	// read the offsets
	err = binary.Write(r.f, binary.BigEndian, &r.offsets)
	if err != nil {
		_ = r.f.Close()
		return err
	}
	r.sectors[0] = true

	// read the timestamps
	err = binary.Write(r.f, binary.BigEndian, &r.timestamps)
	if err != nil {
		_ = r.f.Close()
		return err
	}

	return r.f.Close()
}

func sectorLoc(offset int32) (o, s int32) {
	return offset >> 8, offset & 0xFF
}

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

func (r *Region) WriteSector(x, y int, data []byte) error {
	sectorsNeeded := int32(len(data)+4)/4096 + 1
	sectorNumber, sectorsAllocated := sectorLoc(r.offsets[x][y])

	// maximum chunk size is 1MB
	if sectorsNeeded >= 256 {
		return errors.New("data too large")
	}

	if sectorNumber != 0 && sectorsAllocated == sectorsNeeded {
		// we can simply overwrite the old sectors
	} else {
		// we need to allocate new sectors

		// mark the sectors previously used for this chunk as free
		for i := int32(0); i < sectorsAllocated; i++ {
			r.sectors[sectorNumber+i] = false
		}

		// scan for a free space large enough to store this chunk
		sectorNumber = 0
		for i := int32(0); i < sectorsNeeded; i++ {
			if r.sectors[sectorNumber+i] {
				sectorNumber += i + 1
				i = -1
			}
		}

		// mark the sectors previously used for this chunk as used
		sectorsAllocated = sectorsNeeded
		for i := int32(0); i < sectorsNeeded; i++ {
			r.sectors[sectorNumber+i] = true
		}
		r.offsets[x][y] = (sectorNumber << 8) | (sectorsNeeded & 0xFF)
	}

	_, err := r.f.Seek(4096*int64(sectorNumber), 0)
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
