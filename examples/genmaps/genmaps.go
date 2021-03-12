package main

import (
	"flag"
	"fmt"
	"github.com/Tnze/go-mc/data/block"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

var colors []color.RGBA64

var (
	regionsFold = flag.String("region", filepath.Join(os.Getenv("AppData"), ".minecraft", "saves", "World", "region"), "region directory path")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	de, err := os.ReadDir(*regionsFold)
	if err != nil {
		log.Fatal(err)
	}

	var min, max [2]int
	updateMinMax := func(pos [2]int) {
		mkmax(&max[0], &pos[0])
		mkmax(&max[1], &pos[1])
		mkmin(&min[0], &pos[0])
		mkmin(&min[1], &pos[1])
	}

	// Open mca files
	var rs = make(map[[2]int]*region.Region, len(de))
	for _, dir := range de {
		name := dir.Name()
		path := filepath.Join(*regionsFold, name)
		var pos [2]int // {x, z}
		if _, err := fmt.Sscanf(name, "r.%d.%d.mca", &pos[0], &pos[1]); err != nil {
			log.Printf("Error parsing file name of %s: %v, ignoring", name, err)
			continue
		}
		updateMinMax(pos)

		r, err := region.Open(path)
		if err != nil {
			log.Printf("Error when opening %s: %v, ignoring", name, err)
			continue
		}
		rs[pos] = r
	}
	// To close mca files
	defer func() {
		for pos, r := range rs {
			if err := r.Close(); err != nil {
				log.Printf("Close r.%d.%d.mca error: %v", pos[0], pos[1], err)
			}
		}
	}()

	for pos, r := range rs {
		var column save.Column
		img := image.NewRGBA(image.Rect(0, 0, 32*16, 32*16))
		for x := 0; x < 32; x++ {
			for z := 0; z < 32; z++ {
				if !r.ExistSector(x, z) {
					continue
				}
				data, err := r.ReadSector(x, z)
				if err != nil {
					log.Printf("Read sector (%d.%d) error: %v", x, z, err)
				}
				if err := column.Load(data); err != nil {
					log.Printf("Decode column (%d.%d) error: %v", x, z, err)
				}

				draw.Draw(
					img, image.Rect(x*16, z*16, x*16+16, z*16+16),
					drawColumn(&column), image.Pt(0, 0),
					draw.Over,
				)
			}
		}
		savePng(img, fmt.Sprintf("r.%d.%d.png", pos[0], pos[1]))
		log.Print(pos, r)
	}
}

func drawColumn(column *save.Column) (img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, 16, 16))

	s := column.Level.Sections
	for i := range s {
		s := s[i]
		// calculate bits per block
		bpb := len(s.BlockStates) * 64 / (16 * 16 * 16)
		// skip empty
		if len(s.BlockStates) == 0 {
			continue
		}
		// decode section
		//n := int(math.Max(4, math.Ceil(math.Log2(float64(len(s.Palette))))))

		// decode status
		data := *(*[]uint64)(unsafe.Pointer(&s.BlockStates)) // convert []int64 into []uint64
		bs := save.NewBitStorage(bpb, 4096, data)
		for y := 0; y < 16; y++ {
			layerImg := image.NewRGBA(image.Rect(0, 0, 16, 16))
			for i := 16*16 - 1; i >= 0; i-- {
				b := block.ByID[block.StateID[uint32(bs.Get(y*16*16+i))]]
				c := colors[b.ID]
				layerImg.Set(i/16, i%16, c)
			}
			draw.Draw(
				img, image.Rect(0, 0, 16, 16),
				layerImg, image.Pt(0, 0),
				draw.Over,
			)
		}
	}
	return
}
