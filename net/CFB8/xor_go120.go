//go:build go1.20

package CFB8

import (
	_ "crypto/subtle"
	_ "unsafe"
)

//go:linkname xorBytes crypto/subtle.XORBytes
func xorBytes(dst, x, y []byte) int
