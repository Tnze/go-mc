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

		// If pos has been moved after block size in the slice,
		// we also store the encrypted bytes to the head
		// for next round.
		if cf.ivPos > cf.blockSize {
			// insert the encrypted byte to the head
			if cf.de {
				cf.iv[cf.ivPos-cf.blockSize-1] = src[i]
			} else {
				cf.iv[cf.ivPos-cf.blockSize-1] = val
			}
		}

		if cf.ivPos == cf.blockSize<<1 {
			// bound reached; move to next round for next operation
			// Since we have stored the same things both in the start
			// and the end of the slice (see above and below),
			// the next operation can be spliced without any overhead
			// by simply resetting the pos.
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
