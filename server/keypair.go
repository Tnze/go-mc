package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"sync/atomic"

	"github.com/Tnze/go-mc/net"
	"github.com/google/uuid"
)

// KeyManager is the interface to get the RSA key for Encryption Key Request.
type KeyManager interface {
	GetServerKey(conn *net.Conn, name string, id uuid.UUID, profilePubKey *rsa.PublicKey) (privKey *rsa.PrivateKey, x509Pub []byte, ok bool)
}

type KeyPairCache struct {
	key    *rsa.PrivateKey
	pubKey []byte
}

type DefaultKeyManager struct {
	key atomic.Value // *KeyPairCache
}

func (km *DefaultKeyManager) GenerateKey(bits int) error {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	return km.SetKey(key)
}

func (km *DefaultKeyManager) SetKey(privKey *rsa.PrivateKey) error {
	publicKey, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return err
	}

	kp := &KeyPairCache{
		key:    privKey,
		pubKey: publicKey,
	}
	km.key.Store(kp)
	return nil
}

func (km *DefaultKeyManager) GetKey() (privKey *rsa.PrivateKey, x509Pub []byte, ok bool) {
	kp, ok := km.key.Load().(*KeyPairCache)
	if !ok {
		return nil, nil, false
	}
	return kp.key, kp.pubKey, true
}

func (km *DefaultKeyManager) GetServerKey(conn *net.Conn, name string, id uuid.UUID, profilePubKey *rsa.PublicKey) (privKey *rsa.PrivateKey, x509Pub []byte, ok bool) {
	return km.GetKey() // just return
}

func NewDefaultKeyManager(bits int) (*DefaultKeyManager, error) {
	km := &DefaultKeyManager{}
	err := km.GenerateKey(bits)
	if err != nil {
		return nil, err
	}
	return km, nil
}
