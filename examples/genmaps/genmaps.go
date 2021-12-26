package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/Tnze/go-mc/data/block"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
)

var colors []color.RGBA64
var sectionWorkerNum = 1

var (
	regionWorkerNum = flag.Int("workers", runtime.NumCPU(), "worker numbers")
	regionsFold     = flag.String("region", filepath.Join(os.Getenv("AppData"), ".minecraft", "saves", "World", "region"), "region directory path")
	drawBigPicture  = flag.Bool("bigmap", true, "draw the big map")
)

func main() {
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

	if *drawBigPicture {
		for _, dir := range de {
			name := dir.Name()
			var pos [2]int // {x, z}
			if _, err := fmt.Sscanf(name, "r.%d.%d.mca", &pos[0], &pos[1]); err != nil {
				log.Printf("Error parsing file name of %s: %v, ignoring", name, err)
				continue
			}
			updateMinMax(pos)
		}
	}

	type regions struct {
		pos [2]int
		*region.Region
	}
	// Open mca files
	var rs = make(chan regions, *regionWorkerNum)
	go func() {
		for _, dir := range de {
			name := dir.Name()
			path := filepath.Join(*regionsFold, name)
			var pos [2]int // {x, z}
			if _, err := fmt.Sscanf(name, "r.%d.%d.mca", &pos[0], &pos[1]); err != nil {
				log.Printf("Error parsing file name of %s: %v, ignoring", name, err)
				continue
			}

			r, err := region.Open(path)
			if err != nil {
				log.Printf("Error when opening %s: %v, ignoring", name, err)
				continue
			}
			rs <- regions{pos: pos, Region: r}
		}
		close(rs)
	}()
	var bigPicture *image.RGBA
	if *drawBigPicture {
		bigPicture = image.NewRGBA(image.Rect(min[0]*512, min[1]*512, max[0]*512+512, max[1]*512+512))
	}
	var bigWg sync.WaitGroup
	// draw columns
	for r := range rs {
		img := image.NewRGBA(image.Rect(0, 0, 32*16, 32*16))
		type task struct {
			data []byte
			pos  [2]int
		}
		c := make(chan task)
		var wg sync.WaitGroup
		for i := 0; i < *regionWorkerNum; i++ {
			go func() {
				var column save.Chunk
				for task := range c {
					if err := column.Load(task.data); err != nil {
						log.Printf("Decode column (%d.%d) error: %v", task.pos[0], task.pos[1], err)
					}
					//pos := [2]int{int(column.Level.PosX), int(column.Level.PosZ)}
					//if pos != task.pos {
					//	fmt.Printf("chunk position not match: want %v, get %v\n", task.pos, pos)
					//}
					draw.Draw(
						img, image.Rect(task.pos[0]*16, task.pos[1]*16, task.pos[0]*16+16, task.pos[1]*16+16),
						drawColumn(&column), image.Pt(0, 0),
						draw.Over,
					)
					wg.Done()
				}
			}()
		}

		for x := 0; x < 32; x++ {
			for z := 0; z < 32; z++ {
				if !r.ExistSector(x, z) {
					continue
				}
				data, err := r.ReadSector(x, z)
				if err != nil {
					log.Printf("Read sector (%d.%d) error: %v", x, z, err)
				}
				wg.Add(1)
				c <- task{data: data, pos: [2]int{x, z}}
			}
		}
		close(c)
		wg.Wait()
		// Save pictures
		bigWg.Add(1)
		log.Print("Draw: ", r.pos)
		go func(img image.Image, pos [2]int) {
			savePng(img, fmt.Sprintf("r.%d.%d.png", pos[0], pos[1]))
			if *drawBigPicture {
				draw.Draw(
					bigPicture, image.Rect(pos[0]*512, pos[1]*512, pos[0]*512+512, pos[1]*512+512),
					img, image.Pt(0, 0), draw.Src,
				)
			}
			bigWg.Done()
		}(img, r.pos)
		// To close mca files
		if err := r.Close(); err != nil {
			log.Printf("Close r.%d.%d.mca error: %v", r.pos[0], r.pos[1], err)
		}
	}
	bigWg.Wait()
	if *drawBigPicture {
		savePng(bigPicture, "maps.png")
	}
}

func drawColumn(column *save.Chunk) (img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, 16, 16))
	s := column.Sections
	c := make(chan *save.Section)
	var wg sync.WaitGroup
	for i := 0; i < sectionWorkerNum; i++ {
		go func() {
			for s := range c {
				drawSection(s, img)
				wg.Done()
			}
		}()
	}
	defer close(c)

	wg.Add(len(s))
	for i := range s {
		c <- &s[i]
	}
	wg.Wait()

	return
}

func drawSection(s *save.Section, img *image.RGBA) {
	data := *(*[]uint64)((unsafe.Pointer)(&s.BlockStates.Data))
	palette := s.BlockStates.Palette
	rawPalette := make([]int, len(palette))
	for i, v := range palette {
		// TODO: Consider the properties of block, not only index the block name
		rawPalette[i] = int(stateIDs[strings.TrimPrefix(v.Name, "minecraft:")])
	}
	c := level.NewStatesPaletteContainerWithData(16*16*16, data, rawPalette)
	for y := 0; y < 16; y++ {
		layerImg := image.NewRGBA(image.Rect(0, 0, 16, 16))
		for i := 16*16 - 1; i >= 0; i-- {
			var bid block.ID
			bid = block.ID(c.Get(i))
			c := colors[block.ByID[bid].ID]
			layerImg.Set(i%16, i/16, c)
		}
		draw.Draw(
			img, image.Rect(0, 0, 16, 16),
			layerImg, image.Pt(0, 0),
			draw.Over,
		)
	}
	return
}

// TODO: This map should be moved to data/block.
var stateIDs map[string]uint32

func init() {
	for i, v := range block.StateID {
		name := block.ByID[v].Name
		if _, ok := stateIDs[name]; !ok {
			stateIDs[name] = i
		}
	}
}
