// This is an example of how to use the go-mc/save/region package to read and write .mca files.
// It can unpack .mca files to .mcc files, or repack .mcc files to .mca files, controlled by the -p flags.
package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Tnze/go-mc/save/region"
)

var (
	decomp = flag.Bool("x", false, "decompress each chunk to NBT format")
	repack = flag.Bool("p", false, "repack .mcc file to .mca")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	var o string
	o = "." // output dir
	if len(args) < 2 {
		usage()
	}
	for _, f := range args[1:] {
		fs := must(filepath.Glob(f))

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
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: %s [-x] [-r] r.<X>.<Z>.mc{a,c}\n", os.Args[0])
	os.Exit(1)
}

func unpack(f, o string) {
	var x, z int
	rn := filepath.Base(f)
	must(fmt.Sscanf(rn, "r.%d.%d.mca", &x, &z))

	r := must(region.Open(f))
	defer r.Close()

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if !r.ExistSector(i, j) {
				continue
			}

			data := must(r.ReadSector(i, j))
			var r io.Reader = bytes.NewReader(data[1:])

			fn := fmt.Sprintf("c.%d.%d.mcc", x<<5+i, z<<5+j)
			if *decomp {
				var err error
				fn += ".nbt" // 解压后就是一个标准的NBT文件，可以加个.nbt后缀
				switch data[0] {
				default:
					err = fmt.Errorf("unknown compression type 0x%02x", data[0])
				case 1:
					r, err = gzip.NewReader(r)
				case 2:
					r, err = zlib.NewReader(r)
				}
				must(0, err)
			}

			cf := must(os.OpenFile(filepath.Join(o, fn), os.O_CREATE|os.O_RDWR|os.O_EXCL, 0o666))

			must(io.Copy(cf, r))
		}
	}
}

func pack(f, o string) {
	var x, z int
	rn := filepath.Base(f)
	must(fmt.Sscanf(rn, "c.%d.%d.mcc", &x, &z))

	fn := fmt.Sprintf("r.%d.%d.mca", x>>5, z>>5)
	r, err := region.Open(fn)
	if err != nil && os.IsNotExist(err) {
		r = must(region.Create(filepath.Join(o, fn)))
	} else {
		must(0, err)
	}
	defer r.Close()

	mcc := must(os.ReadFile(f))

	rx, rz := region.In(x, z)
	must(0, r.WriteSector(rx, rz, mcc))
}

func must[T any](v T, err error) T {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return v
}
