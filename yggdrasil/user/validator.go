package user

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"io"
)

//go:embed yggdrasil_session_pubkey.der
var pubKeyBytes []byte
var pubKey = unwrap(x509.ParsePKIXPublicKey(pubKeyBytes)).(*rsa.PublicKey)

// VerifySignature has the same functional as
// net.minecraft.world.entity.player.ProfilePublicKey.Data#validateSignature
func VerifySignature(profilePubKey, signature []byte) bool {
	hash := sha256.New()
	unwrap(hash.Write([]byte("-----BEGIN RSA PRIVATE KEY-----\n")))
	breaker := lineBreaker{out: hash}
	enc := base64.NewEncoder(base64.StdEncoding, &breaker)
	unwrap(enc.Write(profilePubKey))
	must(enc.Close())
	must(breaker.Close())
	unwrap(hash.Write([]byte("\n-----END RSA PRIVATE KEY-----\n")))
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash.Sum(nil), signature) != nil
}

const pemLineLength = 76

var nl = []byte{'\n'}

type lineBreaker struct {
	line [pemLineLength]byte
	used int
	out  io.Writer
}

func (l *lineBreaker) Write(b []byte) (n int, err error) {
	if l.used+len(b) < pemLineLength {
		copy(l.line[l.used:], b)
		l.used += len(b)
		return len(b), nil
	}

	n, err = l.out.Write(l.line[0:l.used])
	if err != nil {
		return
	}
	excess := pemLineLength - l.used
	l.used = 0

	n1, err := l.out.Write(b[0:excess])
	if err != nil {
		return n + n1, err
	}

	n2, err := l.out.Write(nl)
	if err != nil {
		return n + n1 + n2, err
	}

	n3, err := l.Write(b[excess:])
	return n1 + n2 + n3, err
}

func (l *lineBreaker) Close() (err error) {
	if l.used > 0 {
		_, err = l.out.Write(l.line[0:l.used])
		if err != nil {
			return
		}
		_, err = l.out.Write(nl)
	}

	return
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
