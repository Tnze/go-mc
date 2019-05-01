//From https://play.golang.org/p/LTbId4b6M2

package CFB8

import "crypto/cipher"

type CFB8 struct {
	c         cipher.Block
	blockSize int
	iv, tmp   []byte
	de        bool
}

func NewCFB8Decrypt(c cipher.Block, iv []byte) *CFB8 {
	cp := make([]byte, len(iv))
	copy(cp, iv)
	return &CFB8{
		c:         c,
		blockSize: c.BlockSize(),
		iv:        cp,
		tmp:       make([]byte, c.BlockSize()),
		de:        true,
	}
}

func NewCFB8Encrypt(c cipher.Block, iv []byte) *CFB8 {
	cp := make([]byte, len(iv))
	copy(cp, iv)
	return &CFB8{
		c:         c,
		blockSize: c.BlockSize(),
		iv:        cp,
		tmp:       make([]byte, c.BlockSize()),
		de:        false,
	}
}

func (cf *CFB8) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		val := src[i]
		copy(cf.tmp, cf.iv)
		cf.c.Encrypt(cf.iv, cf.iv)
		val = val ^ cf.iv[0]

		copy(cf.iv, cf.tmp[1:])
		if cf.de {
			cf.iv[15] = src[i]
		} else {
			cf.iv[15] = val
		}

		dst[i] = val
	}
}
