// Package CFB8 is copied from https://play.golang.org/p/LTbId4b6M2
package CFB8

import "crypto/cipher"

type CFB8 struct {
	c         cipher.Block
	blockSize int
	ivPos     int
	iv        []byte
	de        bool
}

func NewCFB8Decrypt(c cipher.Block, iv []byte) *CFB8 {
	cp := make([]byte, len(iv)*3)
	copy(cp, iv)
	return &CFB8{
		c:         c,
		blockSize: c.BlockSize(),
		iv:        cp,
		de:        true,
	}
}

func NewCFB8Encrypt(c cipher.Block, iv []byte) *CFB8 {
	cp := make([]byte, len(iv)*3)
	copy(cp, iv)
	return &CFB8{
		c:         c,
		blockSize: c.BlockSize(),
		iv:        cp,
		de:        false,
	}
}

func (cf *CFB8) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		posPlusBlockSize := cf.ivPos + cf.blockSize
		val := src[i]
		// fast mod; 2*blockSize must be a non-negative integer power of 2
		tempPos := posPlusBlockSize & (cf.blockSize<<1 - 1)
		cf.c.Encrypt(cf.iv[tempPos:], cf.iv[cf.ivPos:])
		val ^= cf.iv[tempPos]

		if cf.ivPos > cf.blockSize {
			if cf.de {
				cf.iv[cf.ivPos-cf.blockSize-1] = src[i]
			} else {
				cf.iv[cf.ivPos-cf.blockSize-1] = val
			}
		}

		if cf.ivPos == cf.blockSize<<1 {
			cf.ivPos = 0
		} else {
			if cf.de {
				cf.iv[posPlusBlockSize] = src[i]
			} else {
				cf.iv[posPlusBlockSize] = val
			}
			cf.ivPos += 1
		}

		dst[i] = val
	}
}
