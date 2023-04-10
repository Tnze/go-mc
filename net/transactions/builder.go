package transactions

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type TransactionBuilder struct {
	WindowID pk.UnsignedByte
	StateID  pk.VarInt
	Actions  []*SlotAction
}

type Transaction struct {
	Packets []*pk.Packet
}

func (t *TransactionBuilder) Build() *Transaction {
	packets := make([]*pk.Packet, 0, len(t.Actions))
	for _, action := range t.Actions {
		p := pk.Marshal(packetid.SPacketClickWindow, t.WindowID, t.StateID, action)
		packets = append(packets, &p)
	}
	return &Transaction{Packets: packets}
}
