package CFB8

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"testing"
)

// cfb8Tests contains the test vectors from
// https://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf, section
// F.3.7. Modified for Minecraft CFB8 tests.
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
	{
		"2b7e151628aed2a6abf7158809cf4f3c",
		"000102030405060708090a0b0c0d0e0f",
		"0ecbd6d36cd12962ce671b4d96fb95aaa902096aeac366e13a6ae57c05d48673cf320c626689d05548f65fd6a108630c1d4e3aab543b006823c7a9422e97c0431587537c384f99a11488ffd9b2e9b46f49005a7e5cef64e27e2de3cf3fb87c1524766601",
		"5efb6f6b93cf5f0e135a0c932f59f9aaa2276e4b06cd4f5edca4baba735ac7708dd7c0f9e92c6b89d2245b0d9a6356b0e98529cd45e56df22e914ef9e0792facaab707af90c13162bfad06a240eb6adcbf3365fd84a003f8083f4662a7a27232c72c6c0c",
	},
}

func TestCFB8VectorsNonOverlapping(t *testing.T) {
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
		if len(plaintext) > 50 {
			cfb.XORKeyStream(ciphertext, plaintext[:len(plaintext)/2])
			cfb.XORKeyStream(ciphertext[len(plaintext)/2:], plaintext[len(plaintext)/2:])
		} else {
			cfb.XORKeyStream(ciphertext, plaintext)
		}

		if !bytes.Equal(ciphertext, expected) {
			t.Errorf("#%d: wrong output: got %x, expected %x", i, ciphertext, expected)
		}

		cfbdec := NewCFB8Decrypt(block, iv)
		plaintextCopy := make([]byte, len(ciphertext))
		if len(ciphertext) > 50 {
			cfbdec.XORKeyStream(plaintextCopy, ciphertext[:len(ciphertext)/2])
			cfbdec.XORKeyStream(plaintextCopy[len(ciphertext)/2:], ciphertext[len(ciphertext)/2:])
		} else {
			cfbdec.XORKeyStream(plaintextCopy, ciphertext)
		}

		if !bytes.Equal(plaintextCopy, plaintext) {
			t.Errorf("#%d: wrong plaintext: got %x, expected %x", i, plaintextCopy, plaintext)
		}
	}
}

func TestCFB8VectorsOverlapped(t *testing.T) {
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

		buf := make([]byte, len(plaintext))
		copy(buf, plaintext)
		cfb := NewCFB8Encrypt(block, iv)
		if len(buf) > 50 {
			cfb.XORKeyStream(buf, buf[:len(buf)/2])
			cfb.XORKeyStream(buf[len(buf)/2:], buf[len(buf)/2:])
		} else {
			cfb.XORKeyStream(buf, buf)
		}

		if !bytes.Equal(buf, expected) {
			t.Errorf("#%d: wrong output: got %x, expected %x", i, buf, expected)
		}

		cfbdec := NewCFB8Decrypt(block, iv)
		if len(buf) > 50 {
			cfbdec.XORKeyStream(buf, buf[:len(buf)/2])
			cfbdec.XORKeyStream(buf[len(buf)/2:], buf[len(buf)/2:])
		} else {
			cfbdec.XORKeyStream(buf, buf)
		}

		if !bytes.Equal(buf, plaintext) {
			t.Errorf("#%d: wrong plaintext: got %x, expected %x", i, buf, plaintext)
		}
	}
}

func benchmarkStreamOverlapped(b *testing.B, stream cipher.Stream) {
	buf := make([]byte, 1024)

	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stream.XORKeyStream(buf, buf)
	}
}

func benchmarkStreamNonOverlapping(b *testing.B, stream cipher.Stream) {
	buf := make([]byte, 1024)
	buf2 := make([]byte, 1024)

	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stream.XORKeyStream(buf2, buf)
	}
}
func BenchmarkCFB8AES1KEncryptOverlapped(b *testing.B) {
	var key [16]byte
	var iv [16]byte
	rand.Read(key[:])
	rand.Read(iv[:])
	aes, _ := aes.NewCipher(key[:])
	stream := NewCFB8Encrypt(aes, iv[:])

	benchmarkStreamOverlapped(b, stream)
}

func BenchmarkCFB8AES1KEncryptNonOverlapping(b *testing.B) {
	var key [16]byte
	var iv [16]byte
	rand.Read(key[:])
	rand.Read(iv[:])
	aes, _ := aes.NewCipher(key[:])
	stream := NewCFB8Encrypt(aes, iv[:])

	benchmarkStreamNonOverlapping(b, stream)
}

func BenchmarkCFB8AES1KDecryptOverlapped(b *testing.B) {
	var key [16]byte
	var iv [16]byte
	rand.Read(key[:])
	rand.Read(iv[:])
	aes, _ := aes.NewCipher(key[:])
	stream := NewCFB8Decrypt(aes, iv[:])

	benchmarkStreamOverlapped(b, stream)
}

func BenchmarkCFB8AES1KDecryptNonOverlapping(b *testing.B) {
	var key [16]byte
	var iv [16]byte
	rand.Read(key[:])
	rand.Read(iv[:])
	aes, _ := aes.NewCipher(key[:])
	stream := NewCFB8Decrypt(aes, iv[:])

	benchmarkStreamNonOverlapping(b, stream)
}
