//go:build !go1.20

package CFB8

import (
	_ "crypto/cipher"
	_ "unsafe"
)

//go:linkname xorBytes crypto/cipher.xorBytes
func xorBytes(dst, x, y []byte) int
