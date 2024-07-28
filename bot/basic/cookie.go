package basic

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (p *Player) handleCookieRequestPacket(packet pk.Packet) error {
	var key pk.Identifier
	err := packet.Scan(&key)
	if err != nil {
		return Error{err}
	}
	cookieContent := p.c.Cookies[string(key)]
	err = p.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundCookieResponse,
		key, pk.OptionEncoder[pk.ByteArray]{
			Has: cookieContent != nil,
			Val: pk.ByteArray(cookieContent),
		},
	))
	if err != nil {
		return Error{err}
	}
	return nil
}

func (p *Player) handleStoreCookiePacket(packet pk.Packet) error {
	var key pk.Identifier
	var payload pk.ByteArray
	err := packet.Scan(&key, &payload)
	if err != nil {
		return Error{err}
	}
	p.c.Cookies[string(key)] = []byte(payload)
	return nil
}
