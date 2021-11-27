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
	"github.com/Tnze/go-mc/data/packetid"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/net/CFB8"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

const verifyTokenLen = 16

//Encrypt a connection, with authentication
func Encrypt(conn *net.Conn, name string) (*Resp, error) {
	//generate keys
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}

	//encryption request
	VT1, err := encryptionRequest(conn, publicKey)
	if err != nil {
		return nil, err
	}

	//encryption response
	ESharedSecret, EVerifyToken, err := encryptionResponse(conn)
	if err != nil {
		return nil, err
	}

	//encryption the connection
	SharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, key, ESharedSecret)
	if err != nil {
		return nil, err
	}
	VT2, err := rsa.DecryptPKCS1v15(rand.Reader, key, EVerifyToken)
	if err != nil {
		return nil, err
	}

	//confirm the verify token
	if !bytes.Equal(VT1, VT2) {
		return nil, errors.New("verify token not match")
	}

	block, err := aes.NewCipher(SharedSecret)
	if err != nil {
		return nil, errors.New("load aes encryption key fail")
	}

	conn.SetCipher( //启用加密
		CFB8.NewCFB8Encrypt(block, SharedSecret),
		CFB8.NewCFB8Decrypt(block, SharedSecret))

	hash := authDigest("", SharedSecret, publicKey)
	resp, err := authentication(name, hash) //auth
	if err != nil {
		return nil, errors.New("auth servers down")
	}

	return resp, nil
}

func encryptionRequest(conn *net.Conn, publicKey []byte) ([]byte, error) {
	var verifyToken [verifyTokenLen]byte
	_, err := rand.Read(verifyToken[:])
	if err != nil {
		return nil, err
	}
	err = conn.WritePacket(pk.Marshal(
		packetid.EncryptionBeginClientbound,
		pk.String(""),
		pk.ByteArray(publicKey),
		pk.ByteArray(verifyToken[:]),
	))
	return verifyToken[:], err
}

func encryptionResponse(conn *net.Conn) ([]byte, []byte, error) {
	var p pk.Packet
	err := conn.ReadPacket(&p)
	if err != nil {
		return nil, nil, err
	}
	if p.ID != packetid.EncryptionBeginServerbound {
		return nil, nil, fmt.Errorf("0x%02X is not Encryption Response", p.ID)
	}

	var SharedSecret, VerifyToken pk.ByteArray
	if err = p.Scan(&SharedSecret, &VerifyToken); err != nil {
		return nil, nil, err
	}

	return SharedSecret, VerifyToken, nil
}

func authentication(name, hash string) (*Resp, error) {
	resp, err := http.Get("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=" + name + "&serverId=" + hash)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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

//Resp is the response of authentication
type Resp struct {
	Name       string
	ID         uuid.UUID
	Properties [1]struct {
		Name, Value, Signature string
	}
}

//Texture includes player's skin and cape
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

//Texture unmarshal the base64 encoded texture of Resp
func (r *Resp) Texture() (t Texture, err error) {
	var texture []byte
	texture, err = base64.StdEncoding.DecodeString(r.Properties[0].Value)
	if err != nil {
		return
	}

	err = json.Unmarshal(texture, &t)
	return
}
