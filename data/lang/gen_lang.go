// This program can automatically download language.json file and convert into .go
package main

import (
	"errors"
	"net/http"
	stdos "os"
	"path/filepath"

	"github.com/hack-pad/hackpadfs/os"
)

//go:generate go run $GOFILE
//go:generate go fmt ./...
func main() {
	fsys, err := resolveFS(string(filepath.Separator))
	if err != nil {
		panic(err)
	}
	run(fsys, http.Get, stdos.Args)
}

func resolveFS(base string) (*os.FS, error) {
	fs := os.NewFS()

	baseDirectory, err := fs.FromOSPath(base) // Convert to an FS path
	if err != nil {
		return nil, err
	}

	baseDirFS, err := fs.Sub(baseDirectory) // Run all file system operations rooted at the current working directory
	if err != nil {
		return nil, err
	}

	ofs, ok := baseDirFS.(*os.FS)
	if !ok {
		return nil, errors.New("sub FS not an OS instance FS")
	}

	return ofs, nil
}
