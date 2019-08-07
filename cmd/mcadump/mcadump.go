package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Tnze/go-mc/save/region"
)

var decomp = flag.Bool("x", false, "decompress each chunk to NBT format")
var repack = flag.Bool("p", false, "repack .mcc file to .mca")

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	var f, o string
	switch len(args) {
	default:
		usage()
	case 1:
		f, o = args[0], "."
	case 2:
		f, o = args[0], args[1]
	}

	fs, err := filepath.Glob(f)
	checkerr(err)

	if *repack {
		for _, f := range fs {
			pack(f, o)
		}
	} else {
		for _, f := range fs {
			unpack(f, o)
		}
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: %s [-x] [-r] r.<X>.<Z>.mc{a,c} [outdir]\n", flag.Arg(0))
	os.Exit(1)
}

// we use this function to check for laziness. Don't scold me >_<
func checkerr(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func unpack(f, o string) {
	var x, z int
	rn := filepath.Base(f)
	_, err := fmt.Sscanf(rn, "r.%d.%d.mca", &x, &z)
	if err != nil {
		checkerr(fmt.Errorf("cannot use %s as mca file name: %v", rn, err))
	}

	r, err := region.Open(f)
	checkerr(err)
	defer r.Close()

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if !r.ExistSector(i, j) {
				continue
			}

			data, err := r.ReadSector(i, j)
			checkerr(err)
			var r io.Reader = bytes.NewReader(data[1:])

			fn := fmt.Sprintf("c.%d.%d.mcc", x*32+i, z*32+j)
			if *decomp {
				fn += ".nbt" //解压后就是一个标准的NBT文件，可以加个.nbt后缀
				switch data[0] {
				default:
					err = fmt.Errorf("unknown compression type 0x%02x", data[0])
				case 1:
					r, err = gzip.NewReader(r)
				case 2:
					r, err = zlib.NewReader(r)
				}
				checkerr(err)
			}

			cf, err := os.OpenFile(filepath.Join(o, fn), os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
			checkerr(err)

			_, err = io.Copy(cf, r)
			checkerr(err)
		}
	}
}

func pack(f, o string) {
	var x, z int
	rn := filepath.Base(f)
	_, err := fmt.Sscanf(rn, "c.%d.%d.mcc", &x, &z)
	if err != nil {
		checkerr(fmt.Errorf("cannot use %s as mcc file name: %v", rn, err))
	}

	fn := fmt.Sprintf("r.%d.%d.mca", x>>5, z>>5)
	r, err := region.Open(fn)
	if err != nil && os.IsNotExist(err) {
		r, err = region.Create(filepath.Join(o, fn))
	}
	checkerr(err)
	defer r.Close()

	mcc, err := ioutil.ReadFile(f)
	checkerr(err)

	rx, rz := region.In(x, z)
	err = r.WriteSector(rx, rz, mcc)
	checkerr(err)
}
