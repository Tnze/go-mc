package utils

import (
	"os"
)

// GetCacheDirectory returns the cache directory of go-mc, If the directory does not exist, it will be created.
func GetCacheDirectory() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dir = dir + "/go-mc"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	return dir
}
