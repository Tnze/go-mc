package CFB8

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"testing"
)

// cfb8Tests contains the test vectors from
// https://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf, section
// F.3.13. Modified for CFB8 tests.
var cfb8Tests = []struct {
	key, iv, plaintext, ciphertext string
}{
	{
		"2b7e151628aed2a6abf7158809cf4f3c",
		"000102030405060708090a0b0c0d0e0f",
		"6bc1bee22e409f96e93d7e117393172a",
		"3b79424c9c0dd436bace9e0ed4586a4f",
	},
	{
		"2b7e151628aed2a6abf7158809cf4f3c",
		"3B3FD92EB72DAD20333449F8E83CFB4A",
		"ae2d8a571e03ac9c9eb76fac45af8e51",
		"c8b0723943d71f61a2e5b0e8cedf87c8",
	},
	{
		"2b7e151628aed2a6abf7158809cf4f3c",
		"C8A64537A0B3A93FCDE3CDAD9F1CE58B",
		"30c81c46a35ce411e5fbc1191a0a52ef",
		"260d20e9395d3501067286d3a2a7002f",
	},
	{
		"2b7e151628aed2a6abf7158809cf4f3c",
		"26751F67A3CBB140B1808CF187A4F4DF",
		"f69f2445df4f9b17ad2b417be66c3710",
		"c0af633cd9c599309f924802af599ee6",
	},
}

func TestCFB8Vectors(t *testing.T) {
	for i, test := range cfb8Tests {
		key, err := hex.DecodeString(test.key)
		if err != nil {
			t.Fatal(err)
		}
		iv, err := hex.DecodeString(test.iv)
		if err != nil {
			t.Fatal(err)
		}
		plaintext, err := hex.DecodeString(test.plaintext)
		if err != nil {
			t.Fatal(err)
		}
		expected, err := hex.DecodeString(test.ciphertext)
		if err != nil {
			t.Fatal(err)
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			t.Fatal(err)
		}

		ciphertext := make([]byte, len(plaintext))
		cfb := NewCFB8Encrypt(block, iv)
		cfb.XORKeyStream(ciphertext, plaintext)

		if !bytes.Equal(ciphertext, expected) {
			t.Errorf("#%d: wrong output: got %x, expected %x", i, ciphertext, expected)
		}

		cfbdec := NewCFB8Decrypt(block, iv)
		plaintextCopy := make([]byte, len(ciphertext))
		cfbdec.XORKeyStream(plaintextCopy, ciphertext)

		if !bytes.Equal(plaintextCopy, plaintext) {
			t.Errorf("#%d: wrong plaintext: got %x, expected %x", i, plaintextCopy, plaintext)
		}
	}
}

func BenchmarkCFB8AES1K(b *testing.B) {
	var key [16]byte
	var iv [16]byte
	buf := make([]byte, 1024)
	rand.Read(key[:])
	rand.Read(iv[:])
	aes, _ := aes.NewCipher(key[:])
	stream := NewCFB8Encrypt(aes, iv[:])

	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stream.XORKeyStream(buf, buf)
	}
}
