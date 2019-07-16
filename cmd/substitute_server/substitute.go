//创建一个服务器代理，自带正版验证登录
//Create a server proxy with authenticated login
//非正版客户端连接此程序时可以以正版身份进入指定服务器
//Offline-clients can access the designated server as online when connecting to this program

//That is, it implements a simple server and client
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Tnze/go-mc/authenticate"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/net/CFB8"
	pk "github.com/Tnze/go-mc/net/packet"
	uuid "github.com/satori/go.uuid"
)

const verifyTokenLen = 16

//Proxy login info
const (
	Email = "YourEmail"
	Paswd = "YourPassword"
)

//Proxy server settings
const (
	ServerIP   = "localhost"
	ServerPort = 25565
	onlineMode = true
)

//Target server address
const (
	RemoteServerIP   = "YourTargetServerIP"
	RemoteServerPort = 25565
)

func main() {
	listener, err := net.ListenMC(fmt.Sprintf("%s:%d", ServerIP, ServerPort))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go Handle(conn)
	}
}

//Threshold 指定了数据传输时最小压缩包大小
const Threshold = 256

//Client 封装了与客户端之间的底层的网络交互
type Client struct {
	net.Conn

	ProtocolVersion int32
	Name            string
	ID              uuid.UUID
	skin            string
}

//OnPlayer 在正版玩家连接时被调用，其返回值将被作为断开连接的原因被发送给客户端
//返回值应当为一个JSON Chat值，例如`"msg"`
var OnPlayer func(name string, UUID uuid.UUID, protocol int32) string

//Handle 接受客户端连接
func Handle(conn net.Conn) {
	defer conn.Close()
	c := Client{Conn: conn}

	nextState, err := c.handshake()
	if err != nil {
		// log.Println(err)
		return
	}

	const (
		CheckState  = 1
		PlayerLogin = 2
	)
	switch nextState {
	case CheckState:
		c.status()
	case PlayerLogin:
		signal := make(chan int)
		client := bot.NewClient()
		auth, err := authenticate.Authenticate(Email, Paswd)
		if err != nil {
			panic(err)
		}
		client.Name, client.Auth.UUID, client.AsTk = auth.SelectedProfile.Name, auth.SelectedProfile.ID, auth.AccessToken
		go func() {
			err := client.JoinServer("localhost", 25565)
			if err != nil {
				log.Fatal(err)
			}
			signal <- 1
		}()

		log.Println(c.Conn.Socket.RemoteAddr(), "协议号", c.ProtocolVersion)

		err = c.login()
		if err != nil {
			msg := chat.Message{Translate: ("multiplayer.disconnect." + err.Error())}
			jmsg, err := json.Marshal(msg)
			if err != nil {
				return
			}
			packet := pk.Packet{ID: 0x00, Data: pk.String(string(jmsg)).Encode()}
			c.WritePacket(packet)
			return
		}
		<-signal
		fmt.Println("start proxy")
		conn := client.Conn()
		go func() {
			for {
				p, err := c.ReadPacket()
				if err != nil {
					fmt.Println(err)
					return
				}
				if err := conn.WritePacket(p); err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
		for {
			p, err := conn.ReadPacket()
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := c.WritePacket(p); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (c *Client) handshake() (nextState int32, err error) {
	p, err := c.ReadPacket()
	if err != nil {
		return -1, err
	}
	if p.ID != 0 {
		return -1, fmt.Errorf("packet ID 0x%X is not handshake", p.ID)
	}

	var (
		sid pk.String
		spt pk.Short
	)
	if err := p.Scan(
		(*pk.VarInt)(&c.ProtocolVersion),
		&sid, &spt,
		(*pk.VarInt)(&nextState)); err != nil {
		return -1, err
	}

	//检查服务器ID和端口是否匹配
	// if sid != ServerID || uint16(spt) != ServerPort {
	// 	return -1, fmt.Errorf("server address rejected")
	// }

	return nextState, nil
}

func (c *Client) status() {
	for i := 0; i < 2; i++ {
		p, err := c.ReadPacket()
		if err != nil {
			break
		}

		switch p.ID {
		case 0x00:
			respPack := getStatus()
			c.WritePacket(respPack)
		case 0x01:
			c.WritePacket(p)
		}
	}
}

func getStatus() pk.Packet {
	return pk.Packet{
		ID: 0x00,
		Data: pk.String(`
		{
			"version": {
				"name": "1.14.1",
				"protocol": 480
			},
			"players": {
				"max": 1,
				"online": 0,
				"sample": []
			},	
			"description": {
				"text": "军刀破服"
			}
		}
		`).Encode(),
	}
}

func disconnectID(protocol int32) byte {
	switch protocol {
	case 404:
		return 0x1B
	case 477, 480:
		return 0x1A
	default:
		return 0x1A
	}
}

func (c *Client) login() (err error) {
	c.Name, err = c.loginStart()
	if err != nil {
		return fmt.Errorf("unexpected_query_response")
	}

	if Threshold >= 0 {
		err = c.setCompression(Threshold)
		if err != nil {
			return fmt.Errorf("unexpected_query_response")
		}
	}

	if onlineMode {
		key, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			return fmt.Errorf("unexpected_query_response")
		}

		publicKey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
		if err != nil {
			return fmt.Errorf("generic")
		}

		VT1, err := c.encryptionRequest(publicKey)
		if err != nil {
			return fmt.Errorf("generic")
		}

		ESharedSecret, EVerifyToken, err := c.encryptionResponse()
		if err != nil {
			return fmt.Errorf("generic")
		}

		SharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, key, ESharedSecret)
		if err != nil {
			return fmt.Errorf("generic")
		}
		VT2, err := rsa.DecryptPKCS1v15(rand.Reader, key, EVerifyToken)
		if err != nil {
			return fmt.Errorf("generic")
		}

		if !bytes.Equal(VT1, VT2) {
			return fmt.Errorf("generic")
		}

		b, err := aes.NewCipher(SharedSecret)
		if err != nil {
			return fmt.Errorf("generic")
		}
		//启用加密
		c.SetCipher(
			CFB8.NewCFB8Encrypt(b, SharedSecret),
			CFB8.NewCFB8Decrypt(b, SharedSecret),
		)
		hash := authDigest("", SharedSecret, publicKey)
		resp, err := c.authentication(hash)
		if err != nil {
			return fmt.Errorf("authservers_down")
		}

		c.ID, err = uuid.FromString(resp.ID)
		if err != nil {
			return fmt.Errorf("authservers_down")
		}

		if c.Name != resp.Name {
			return fmt.Errorf("unverified_username")
		}

		c.skin = resp.Properties[0].Value

	}

	err = c.loginSuccess()
	if err != nil {
		return fmt.Errorf("generic")
	}
	return
}

func (c *Client) loginStart() (string, error) {
	loginStart, err := c.ReadPacket()
	if err != nil {
		return "", err
	}
	if loginStart.ID != 0x00 {
		return "", fmt.Errorf("0x%02X is not LoginStart packet's ID", loginStart.ID)
	}
	var name pk.String
	err = loginStart.Scan(&name)
	return string(name), err
}

func (c *Client) setCompression(threshold int) error {
	sc := pk.Packet{
		ID:   0x03,
		Data: pk.VarInt(threshold).Encode(),
	}
	err := c.WritePacket(sc)
	c.SetThreshold(threshold)
	return err
}

func (c *Client) loginSuccess() error {
	ls := pk.Packet{ID: 0x02}
	ls.Data = append(ls.Data, pk.String(c.ID.String()).Encode()...)
	ls.Data = append(ls.Data, pk.String(c.Name).Encode()...)
	err := c.WritePacket(ls)
	return err
}

func (c *Client) encryptionRequest(publicKey []byte) ([]byte, error) {
	var verifyToken [verifyTokenLen]byte
	_, err := rand.Read(verifyToken[:])
	if err != nil {
		return nil, err
	}

	er := pk.Packet{ID: 0x01}
	er.Data = append(er.Data, pk.String("").Encode()...)
	er.Data = append(er.Data, pk.VarInt(len(publicKey)).Encode()...)
	er.Data = append(er.Data, publicKey...)
	er.Data = append(er.Data, pk.VarInt(verifyTokenLen).Encode()...)
	er.Data = append(er.Data, verifyToken[:]...)

	err = c.WritePacket(er)
	return verifyToken[:], err
}

func (c *Client) encryptionResponse() ([]byte, []byte, error) {
	p, err := c.ReadPacket()
	if err != nil {
		return nil, nil, err
	}
	if p.ID != 0x01 {
		return nil, nil, fmt.Errorf("0x%02X is not Encryption Response", p.ID)
	}

	var (
		SharedSecret pk.ByteArray
		VerifyToken  pk.ByteArray
	)
	if err := p.Scan(&SharedSecret, &VerifyToken); err != nil {
		return nil, nil, err
	}
	return SharedSecret, VerifyToken, nil
}

type authResp struct {
	ID, Name   string
	Properties [1]struct {
		Name, Value, Signature string
	}
}

func (c *Client) authentication(hash string) (*authResp, error) {
	resp, err := http.Get(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%s&serverId=%s",
		c.Name, hash))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var Resp authResp

	err = json.Unmarshal(body, &Resp)
	if err != nil {
		return nil, err
	}

	return &Resp, nil
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
