package filesystem

import "io/fs"

type FS interface {
	Open(name string) (fs.File, error)
	OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error)
	fs.StatFS
	fs.ReadDirFS
	WriteFile(name string, data []byte, perm fs.FileMode) error
	Mkdir(name string, perm fs.FileMode) error
}
