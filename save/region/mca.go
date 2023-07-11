package region

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"time"
)

var (
	ErrNoSector             = errors.New("sector does not exist")
	ErrNoData               = errors.New("data is missing")
	ErrSectorNegativeLength = errors.New("declared length of data is negative")
	ErrTooLarge             = errors.New("data too large")
)

// Region contain 32*32 chunks in one .mca file
// Not MT-Safe!
type Region struct {
	f          io.ReadWriteSeeker
	offsets    [32][32]int32
	Timestamps [32][32]int32

	// sectors record if a sector is in used.
	// contrary to mojang's, because false is the default value in Go.
	sectors map[int32]bool
}

// In calculate chunk's coordinates relative to region
// 计算chunk在region中的相对坐标。即，除以32并取余。
func In(cx, cz int) (int, int) {
	// c & (32-1)
	// is equal to:
	// (c %= 32) > 0 ? c : -c; //C language
	return cx & 31, cz & 31
}

// At calculate the region's coordinates where the chunk in
// 计算chunk在哪一个region中
func At(cx, cz int) (int, int) {
	return cx >> 5, cz >> 5
}

// Open a .mca file and read the head.
// Close the Region after used.
func Open(name string) (r *Region, err error) {
	f, err := os.OpenFile(name, os.O_RDWR, 0o666)
	if err != nil {
		return nil, err
	}
	r, err = Load(f)
	if err != nil {
		_ = f.Close()
	}
	return
}

// Load works like Open but read from an io.ReadWriteSeeker.
func Load(f io.ReadWriteSeeker) (r *Region, err error) {
	r = &Region{
		f:       f,
		sectors: make(map[int32]bool),
	}

	// read the offsets
	err = binary.Read(r.f, binary.BigEndian, &r.offsets)
	if err != nil {
		return nil, err
	}
	r.sectors[0] = true

	// read the timestamps
	err = binary.Read(r.f, binary.BigEndian, &r.Timestamps)
	if err != nil {
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
func Create(name string) (*Region, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0o666)
	if err != nil {
		return nil, err
	}
	return CreateWriter(f)
}

// CreateWriter create Region by an io.ReadWriteSeeker
func CreateWriter(f io.ReadWriteSeeker) (r *Region, err error) {
	r = new(Region)
	r.sectors = make(map[int32]bool)
	r.f = f

	// write the offsets
	err = binary.Write(r.f, binary.BigEndian, &r.offsets)
	if err != nil {
		_ = r.Close()
		return nil, err
	}
	r.sectors[0] = true

	// write the timestamps
	err = binary.Write(r.f, binary.BigEndian, &r.Timestamps)
	if err != nil {
		_ = r.Close()
		return nil, err
	}
	r.sectors[1] = true

	return r, nil
}

// Close the region file if possible.
// The responsibility for Close the file with who Open it.
// If you made the Region with Load(),
// this method close the file only if your io.ReadWriteSeeker implement io.Close
func (r *Region) Close() error {
	if closer, ok := r.f.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func sectorLoc(offset int32) (sec, num int32) {
	return (offset >> 8) & 0xFFFFFF, offset & 0xFF
}

// ReadSector find and read the Chunk data from region
func (r *Region) ReadSector(x, z int) (data []byte, err error) {
	sec, num := sectorLoc(r.offsets[z][x])
	if sec == 0 {
		return nil, ErrNoSector
	}
	_, err = r.f.Seek(4096*int64(sec), 0)
	if err != nil {
		return
	}
	reader := io.LimitReader(r.f, 4096*int64(num))

	var length int32
	err = binary.Read(reader, binary.BigEndian, &length)
	if err != nil {
		return
	}
	if length == 0 {
		return nil, ErrNoData
	}
	if length < 0 {
		return nil, ErrSectorNegativeLength
	}
	if length > 4096*num {
		return nil, ErrTooLarge
	}
	data = make([]byte, length)
	_, err = io.ReadFull(reader, data)

	return
}

// WriteSector write Chunk data into region file
func (r *Region) WriteSector(x, z int, data []byte) error {
	need := int32((len(data) + 4 + 4096 - 1) / 4096)
	n, now := sectorLoc(r.offsets[z][x])

	// maximum chunk size is 1MB
	if need >= 256 {
		return ErrTooLarge
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

		r.offsets[z][x] = (n << 8) | (need & 0xFF)

		// update file head
		err := r.setHead(x, z, uint32(r.offsets[z][x]), uint32(time.Now().Unix()))
		if err != nil {
			return err
		}
		r.Timestamps[x][z] = int32(time.Now().Unix())
	}

	_, err := r.f.Seek(4096*int64(n), 0)
	if err != nil {
		return err
	}
	// data length
	err = binary.Write(r.f, binary.BigEndian, int32(len(data)))
	if err != nil {
		return err
	}

	// data
	_, err = r.f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// ExistSector return if a sector is existed
func (r *Region) ExistSector(x, z int) bool {
	return r.offsets[z][x] != 0
}

// PadToFullSector writes zeros to the end of the file to make size a multiple of 4096.
// Legacy versions of Minecraft require this.
// Need to be called right before Close.
func (r *Region) PadToFullSector() error {
	size, err := r.f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	if size%4096 != 0 {
		_, err = r.f.Write(make([]byte, 4096-size%4096))
		if err != nil {
			return err
		}
	}
	return nil
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

func (r *Region) setHead(x, z int, offset, timestamp uint32) (err error) {
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:], offset)
	_, err = r.writeAt(buf[:], 4*(int64(z)*32+int64(x)))
	if err != nil {
		return
	}

	binary.BigEndian.PutUint32(buf[:], timestamp)
	_, err = r.writeAt(buf[:], 4096+4*(int64(z)*32+int64(x)))
	if err != nil {
		return
	}

	return
}

func (r *Region) writeAt(p []byte, off int64) (n int, err error) {
	if f, ok := r.f.(io.WriterAt); ok {
		return f.WriteAt(p, off)
	}
	_, err = r.f.Seek(off, 0)
	if err != nil {
		return 0, err
	}
	return r.f.Write(p)
}
