package decoder

import (
	"MineArcade-backend/minearcade-server/protocol"
	"MineArcade-backend/minearcade-server/protocol/packets"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"fmt"
)

func DecodeClientPacket(r *protocol.Reader) (cli_pk packet_define.ClientPacket, err error) {
	defer func() {
		err_orig := recover()
		if err_orig == nil {
			return
		}
		cli_pk = nil
		err = fmt.Errorf(err_orig.(string))
	}()
	cli_pk = unsafeDecodeClientPacket(r)
	return
}

// will panic for errors such as EOF
func unsafeDecodeClientPacket(r *protocol.Reader) packet_define.ClientPacket {
	var id_orig int32
	r.Int32(&id_orig)
	id := uint32(id_orig)
	packet_fn := packets.ClientPool[id]
	if packet_fn == nil {
		panic(fmt.Sprintf("Packet not found: %d", id))
	}
	packet := packet_fn()
	packet.Unmarshal(r)
	return packet
}
