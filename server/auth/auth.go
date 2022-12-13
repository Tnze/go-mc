package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/net/CFB8"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

const verifyTokenLen = 16

// Encrypt a connection, with authentication
func Encrypt(conn *net.Conn, name string, serverKey *rsa.PrivateKey) (*Resp, error) {
	publicKey, err := x509.MarshalPKIXPublicKey(&serverKey.PublicKey)
	if err != nil {
		return nil, err
	}

	verifyToken := make([]byte, verifyTokenLen)
	_, err = rand.Read(verifyToken)
	if err != nil {
		return nil, err
	}

	// encryption request
	err = encryptionRequest(conn, publicKey, verifyToken)
	if err != nil {
		return nil, err
	}

	// encryption response
	SharedSecret, err := encryptionResponse(conn, serverKey, verifyToken)
	if err != nil {
		return nil, err
	}

	// encryption the connection
	block, err := aes.NewCipher(SharedSecret)
	if err != nil {
		return nil, errors.New("load aes encryption key fail")
	}

	conn.SetCipher( // 启用加密
		CFB8.NewCFB8Encrypt(block, SharedSecret),
		CFB8.NewCFB8Decrypt(block, SharedSecret),
	)
	hash := authDigest("", SharedSecret, publicKey)
	resp, err := authentication(name, hash) // auth
	if err != nil {
		return nil, errors.New("auth servers down")
	}

	return resp, nil
}

func encryptionRequest(conn *net.Conn, publicKey, verifyToken []byte) error {
	return conn.WritePacket(pk.Marshal(
		packetid.LoginEncryptionRequest,
		pk.String(""),
		pk.ByteArray(publicKey),
		pk.ByteArray(verifyToken),
	))
}

func encryptionResponse(conn *net.Conn, serverKey *rsa.PrivateKey, verifyToken []byte) ([]byte, error) {
	var p pk.Packet
	err := conn.ReadPacket(&p)
	if err != nil {
		return nil, err
	}
	if p.ID != packetid.LoginEncryptionResponse {
		return nil, fmt.Errorf("0x%02X is not Encryption Response", p.ID)
	}

	var keyBytes pk.ByteArray
	var encryptedVerifyToken pk.ByteArray

	err = p.Scan(&keyBytes, &encryptedVerifyToken)
	if err != nil {
		return nil, err
	}

	// confirm to verify token
	decryptedVerifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, serverKey, encryptedVerifyToken)
	if err != nil {
		return nil, err
	} else if !bytes.Equal(verifyToken, decryptedVerifyToken) {
		return nil, errors.New("verifyToken not match")
	}

	// get sharedSecret
	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, serverKey, keyBytes)
	if err != nil {
		return nil, err
	}

	return sharedSecret, nil
}

func authentication(name, hash string) (*Resp, error) {
	resp, err := http.Get("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=" + name + "&serverId=" + hash)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var Resp Resp
	err = json.Unmarshal(body, &Resp)

	return &Resp, err
}

// authDigest computes a special SHA-1 digest required for Minecraft web
// authentication on Premium servers (online-mode=true).
// Source: http://wiki.vg/Protocol_Encryption#Server
//
// Also many, many thanks to SirCmpwn and his wonderful gist (C#):
// https://gist.github.com/SirCmpwn/404223052379e82f91e6
func authDigest(serverID string, sharedSecret, publicKey []byte) string {
	h := sha1.New()
	h.Write([]byte(serverID))
	h.Write(sharedSecret)
	h.Write(publicKey)
	hash := h.Sum(nil)

	// Check for negative hashes
	negative := (hash[0] & 0x80) == 0x80
	if negative {
		hash = twosComplement(hash)
	}

	// Trim away zeroes
	res := strings.TrimLeft(fmt.Sprintf("%x", hash), "0")
	if negative {
		res = "-" + res
	}

	return res
}

// little endian
func twosComplement(p []byte) []byte {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = byte(^p[i])
		if carry {
			carry = p[i] == 0xff
			p[i]++
		}
	}
	return p
}

// Resp is the response of authentication
type Resp struct {
	Name       string
	ID         uuid.UUID
	Properties []Property
}

type Property struct {
	Name, Value, Signature string
}

func (p Property) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.String(p.Name),
		pk.String(p.Value),
		pk.Option[pk.String, *pk.String]{
			Has: p.Signature != "",
			Val: pk.String(p.Signature),
		},
	}.WriteTo(w)
}

// Texture includes player's skin and cape
type Texture struct {
	TimeStamp int64     `json:"timestamp"`
	ID        uuid.UUID `json:"profileId"`
	Name      string    `json:"profileName"`
	Textures  struct {
		SKIN, CAPE struct {
			URL string `json:"url"`
		}
	} `json:"textures"`
}

// Texture unmarshal the base64 encoded texture of Resp
func (r *Resp) Texture() (t Texture, err error) {
	var texture []byte
	texture, err = base64.StdEncoding.DecodeString(r.Properties[0].Value)
	if err != nil {
		return
	}

	err = json.Unmarshal(texture, &t)
	return
}
