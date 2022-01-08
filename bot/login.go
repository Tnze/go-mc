package bot

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/login"
	pk "github.com/Tnze/go-mc/net/packet"
)

// Auth includes an account
type Auth struct {
	Name string
	UUID string
	AsTk string
}

func handleEncryptionRequest(c *Client, p pk.Packet) error {
	// 创建AES对称加密密钥
	key, encryptStream, decryptStream := login.NewSymmetricEncryption()

	// Read EncryptionRequest
	var er login.EncryptionRequest
	if err := p.Scan(&er); err != nil {
		return err
	}

	err := loginAuth(c.Auth, key, er) //向Mojang验证
	if err != nil {
		return fmt.Errorf("login fail: %v", err)
	}

	// 响应加密请求
	// Write Encryption Key Response
	p, err = er.GenResponsePacket(key)
	if err != nil {
		return fmt.Errorf("gen encryption key response fail: %v", err)
	}

	err = c.Conn.WritePacket(p)
	if err != nil {
		return err
	}

	// 设置连接加密
	c.Conn.SetCipher(encryptStream, decryptStream)
	return nil
}

func loginAuth(auth Auth, shareSecret []byte, er login.EncryptionRequest) error {
	digest := login.AuthDigest(er.ServerID, shareSecret, er.PublicKey)
	return login.JoinServer(auth.AsTk, auth.UUID, digest)
}
