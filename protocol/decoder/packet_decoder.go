package decoder

import (
	"MineArcade-backend/protocol"
	"MineArcade-backend/protocol/packets"
	"fmt"
)

func DecodeClientPacket(r *protocol.Reader) (cli_pk packets.ClientPacket, err error) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		cli_pk = nil
		err = err.(error)
	}()
	cli_pk = unsafeDecodeClientPacket(r)
	return
}

func unsafeDecodeClientPacket(r *protocol.Reader) packets.ClientPacket {
	var id uint32
	r.UInt32(&id)
	packet_fn := packets.ClientPool[id]
	if packet_fn == nil {
		panic(fmt.Sprintf("Packet not found: %d", id))
	}
	packet := packet_fn()
	packet.Unmarshal(r)
	return packet
}
