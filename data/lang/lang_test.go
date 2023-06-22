package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/matryer/is"
)

type mapFSWithMkdir struct {
	fs fstest.MapFS
}

func (m mapFSWithMkdir) Open(name string) (fs.File, error) {
	return m.fs.Open(name)
}

func (m mapFSWithMkdir) OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error) {
	createMode := flag & os.O_CREATE
	truncMode := flag & os.O_TRUNC
	if createMode != 0 {
		m.fs[name] = &fstest.MapFile{
			Mode: perm,
		}
	}

	if truncMode != 0 {
		m.fs[name].Data = nil
	}

	return m.fs.Open(name)
}

func (m mapFSWithMkdir) Mkdir(name string, perm fs.FileMode) error {
	m.fs[name] = &fstest.MapFile{}
	return nil
}

func (m mapFSWithMkdir) ReadDir(name string) ([]fs.DirEntry, error) {
	return m.fs.ReadDir(name)
}

func (m mapFSWithMkdir) Stat(name string) (fs.FileInfo, error) {
	return m.fs.Stat(name)
}
func (m mapFSWithMkdir) WriteFile(name string, data []byte, perm fs.FileMode) error {
	if _, ok := m.fs[name]; !ok {
		return os.ErrNotExist
	}
	m.fs[name].Data = data
	return nil
}

func buildMockFS() mapFSWithMkdir {
	return mapFSWithMkdir{
		fs: fstest.MapFS{},
	}
}

func TestRunWithNoArgs(t *testing.T) {
	is := is.New(t)

	fs := buildMockFS()

	is.NoErr(run(fs, []string{}))

	for k, v := range fs.fs {
		fmt.Printf("%s: %s\n", k, string(v.Data))
	}

	is.Equal(len(fs.fs), 1)
}
