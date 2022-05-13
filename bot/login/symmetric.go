package login

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/Tnze/go-mc/net/CFB8"
)

// AES/CFB8 with random key
func NewSymmetricEncryption() (key []byte, encryptStream, decryptStream cipher.Stream) {
	key = make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	b, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	decryptStream = CFB8.NewCFB8Decrypt(b, key)
	encryptStream = CFB8.NewCFB8Encrypt(b, key)
	return
}
