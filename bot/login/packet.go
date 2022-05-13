package login

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func (e *EncryptionRequest) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		(*pk.String)(&e.ServerID),
		(*pk.ByteArray)(&e.PublicKey),
		(*pk.ByteArray)(&e.VerifyToken),
	}.ReadFrom(r)
}

// GenResponsePacket generates a ready-to-send Encryption Response packet for Encryption Request.
func (e EncryptionRequest) GenResponsePacket(sharedSecret []byte) (pk.Packet, error) {
	iPK, err := x509.ParsePKIXPublicKey(e.PublicKey) // Decode Public Key
	if err != nil {
		err = fmt.Errorf("decode public key fail: %v", err)
		return pk.Packet{}, err
	}
	rsaKey := iPK.(*rsa.PublicKey)
	cryptPK, err := rsa.EncryptPKCS1v15(rand.Reader, rsaKey, sharedSecret)
	if err != nil {
		err = fmt.Errorf("encryption share secret fail: %v", err)
		return pk.Packet{}, err
	}
	verifyT, err := rsa.EncryptPKCS1v15(rand.Reader, rsaKey, e.VerifyToken)
	if err != nil {
		err = fmt.Errorf("encryption verfy tokenfail: %v", err)
		return pk.Packet{}, err
	}
	return pk.Marshal(
		packetid.LoginEncryptionResponse,
		pk.ByteArray(cryptPK),
		pk.ByteArray(verifyT),
	), nil
}
