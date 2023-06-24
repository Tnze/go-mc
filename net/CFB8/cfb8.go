// Package CFB8 implements CFB8 block cipher mode of operation used by Minecraft protocol.
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
	return newCFB8(c, iv, true)
}

func NewCFB8Encrypt(c cipher.Block, iv []byte) *CFB8 {
	return newCFB8(c, iv, false)
}

func newCFB8(c cipher.Block, iv []byte, de bool) *CFB8 {
	cp := make([]byte, len(iv)*3)
	copy(cp, iv)
	return &CFB8{
		c:         c,
		blockSize: c.BlockSize(),
		iv:        cp,
		de:        de,
	}
}

func (cf *CFB8) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		posPlusBlockSize := cf.ivPos + cf.blockSize
		val := src[i]
		// fast mod; 2*blockSize must be a non-negative integer power of 2
		tempPos := posPlusBlockSize & (cf.blockSize<<1 - 1)
		// reuse space to store encrypted block
		cf.c.Encrypt(cf.iv[tempPos:], cf.iv[cf.ivPos:])
		// Only the first byte of the encrypted block is used
		// for encryption/decryption, other bytes are ignored.
		val ^= cf.iv[tempPos]

		if cf.ivPos == cf.blockSize<<1 {
			// bound reached; move to next round for next operation
			// copy next block to the start of tbe ring buffer
			copy(cf.iv, cf.iv[cf.ivPos+1:])
			// insert the encrypted byte to the end of IV
			if cf.de {
				cf.iv[cf.blockSize-1] = src[i]
			} else {
				cf.iv[cf.blockSize-1] = val
			}
			cf.ivPos = 0
		} else {
			// insert the encrypted byte to the end of IV
			if cf.de {
				cf.iv[posPlusBlockSize] = src[i]
			} else {
				cf.iv[posPlusBlockSize] = val
			}
			// move to next block
			cf.ivPos += 1
		}

		dst[i] = val
	}
}
