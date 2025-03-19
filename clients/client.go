package clients

import (
	"MineArcade-backend/clients/accounts"
	"MineArcade-backend/clients/player_store"
	"MineArcade-backend/protocol"
	"MineArcade-backend/protocol/decoder"
	"MineArcade-backend/protocol/packets"
	"fmt"
	"net"
	"time"

	"github.com/pterm/pterm"
)

type NetClient struct {
	Conn      net.Conn
	IPString  string
	AuthInfo  *accounts.UserAuthInfo
	StoreInfo *player_store.PlayerStore
	pkReader  protocol.Reader
}

func (c *NetClient) WritePacket(p packets.ServerPacket) error {
	w := protocol.NewWriter()
	w.Int32(int32(p.ID()))
	p.Marshal(&w)
	_, err := c.Conn.Write(w.GetFullBytes())
	return err
}

func (c *NetClient) ReadNextPacket() (packets.ClientPacket, error) {
	if !c.pkReader.End() {
		pk, err := decoder.DecodeClientPacket(&c.pkReader)
		if err != nil {
			return nil, err
		}
		return pk, nil
	} else {
		buf := make([]byte, 524288) // 512KB
		n, err := c.Conn.Read(buf)
		// todo
		if err != nil {
			return nil, err
		}
		if n == 0 {
			return nil, fmt.Errorf("EOF")
		}
		c.pkReader.SetFullBytes(buf, n)
		return c.ReadNextPacket()
	}
}

func (c *NetClient) InitAuthInfo(info *accounts.UserAuthInfo) {
	c.AuthInfo = info
}

func (c *NetClient) InitStoreInfo(info *player_store.PlayerStore) {
	c.StoreInfo = info
}

func (c *NetClient) Kick(kick_msg string) {
	c.WritePacket(&packets.KickClient{
		Message:    kick_msg,
		StatusCode: 0,
	})
	pterm.Warning.Printfln("踢出客户端 %s: %s", c.IPString, kick_msg)
	time.Sleep(time.Second)
	c.Conn.Close()
}
