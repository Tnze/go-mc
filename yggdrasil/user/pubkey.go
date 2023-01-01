package user

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"io"
	"time"

	pk "github.com/Tnze/go-mc/net/packet"
)

type PublicKey struct {
	ExpiresAt time.Time
	PubKey    *rsa.PublicKey
	Signature []byte
}

func (p PublicKey) WriteTo(w io.Writer) (n int64, err error) {
	pubKeyEncoded, err := x509.MarshalPKIXPublicKey(p.PubKey)
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		pk.Long(p.ExpiresAt.UnixMilli()),
		pk.ByteArray(pubKeyEncoded),
		pk.ByteArray(p.Signature),
	}.WriteTo(w)
}

func (p *PublicKey) ReadFrom(r io.Reader) (n int64, err error) {
	var (
		ExpiresAt pk.Long
		PubKey    pk.ByteArray
		Signature pk.ByteArray
	)
	n, err = pk.Tuple{
		&ExpiresAt,
		&PubKey,
		&Signature,
	}.ReadFrom(r)
	if err != nil {
		return n, err
	}
	p.ExpiresAt = time.UnixMilli(int64(ExpiresAt))
	pubKey, err := x509.ParsePKIXPublicKey(PubKey)
	if err != nil {
		return n, err
	}
	if key, ok := pubKey.(*rsa.PublicKey); !ok {
		return n, errors.New("expect RSA public key")
	} else {
		p.PubKey = key
	}

	p.Signature = Signature
	return n, nil
}

func (p *PublicKey) Verify() bool {
	if p.ExpiresAt.Before(time.Now()) {
		return false
	}
	encoded, err := x509.MarshalPKIXPublicKey(p.PubKey)
	if err != nil {
		return false
	}
	return VerifySignature(encoded, p.Signature)
}

func (p *PublicKey) VerifyMessage(hash, signature []byte) error {
	return rsa.VerifyPKCS1v15(p.PubKey, crypto.SHA256, hash, signature)
}
