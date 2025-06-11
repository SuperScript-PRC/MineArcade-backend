package clients

import (
	"MineArcade-backend/minearcade-server/clients/accounts"
	"MineArcade-backend/minearcade-server/clients/player_store"
	"MineArcade-backend/minearcade-server/protocol/handler"
	"MineArcade-backend/minearcade-server/protocol/packets"
	packets_general "MineArcade-backend/minearcade-server/protocol/packets/general"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"time"
)

type ArcadeClient struct {
	TCPConn      net.Conn
	UDPConn      *net.UDPConn
	TCPAddr      net.Addr
	UDPAddr      *net.UDPAddr
	AuthInfo     *accounts.UserAuthInfo
	StoreInfo    *player_store.PlayerStore
	PacketReader *handler.PacketReader
	PacketWriter *handler.PacketWriter
	Online       bool
}

// 初始化客户端，开启数据包接收和发送通道。

func NewArcadeClient(tcp_conn net.Conn, udp_conn *net.UDPConn) *ArcadeClient {
	return &ArcadeClient{
		TCPConn: tcp_conn,
		UDPConn: udp_conn,
		TCPAddr: tcp_conn.RemoteAddr(),
		Online:  true,
	}
}

func MakeClient(tcp_conn net.Conn, udp_conn *net.UDPConn) *ArcadeClient {
	cli := NewArcadeClient(tcp_conn, udp_conn)
	cli.Init()
	return cli
}

func (c *ArcadeClient) Init() {
	c.PacketReader = handler.NewPacketReader(c.TCPConn, c.UDPConn, c.UDPAddr)
	c.PacketWriter = handler.NewPacketWriter(c.TCPConn, c.UDPConn, c.UDPAddr)
	c.PacketReader.Active()
	c.PacketWriter.Active()
}

func (c *ArcadeClient) SetUDPAddr(udp_addr *net.UDPAddr) {
	c.UDPAddr = udp_addr
	c.PacketWriter.UDPConnAddr = udp_addr
}

// 向此客户端发送数据包。
// 数据包不会立即发送，而是先缓存到发送通道中，等待发送。
func (c *ArcadeClient) WritePacket(p packets.ServerPacket) error {
	return c.PacketWriter.WritePacket(p)
}

// 接收客户端发送的下一个数据包。
func (c *ArcadeClient) NextPacket() (packets.ClientPacket, error) {
	if !c.Online {
		return nil, errors.New("client not online")
	}
	return c.PacketReader.NextPacket()
}

// 接收客户端发送的下一个数据包，但是允许中断
func (c *ArcadeClient) NextPacketWithInterrupt(ch chan bool) (packets.ClientPacket, error, bool) {
	if !c.Online {
		return nil, errors.New("client not online"), false
	}
	return c.PacketReader.NextPacketWithInterrupt(ch)
}

func (c *ArcadeClient) WaitForPacket(pkID int, timeout time.Duration) (pk packets.ClientPacket, getted bool) {
	return c.PacketReader.WaitForPacket(pkID, timeout)
}

// 存入并初始化登录信息。
func (c *ArcadeClient) InitAuthInfo(info *accounts.UserAuthInfo) {
	c.AuthInfo = info
}

// 存入并初始化玩家背包信息。
func (c *ArcadeClient) InitStoreInfo(info *player_store.PlayerStore) {
	c.StoreInfo = info
}

// 踢出客户端。
func (c *ArcadeClient) Kick(kick_msg string) {
	c.WritePacket(&packets_general.KickClient{
		Message:    kick_msg,
		StatusCode: 0,
	})
	slog.Info(fmt.Sprintf("踢出客户端 %s: %s", c.TCPAddr.String(), kick_msg))
	time.Sleep(time.Second)
	c.TCPConn.Close()
	c.Online = false
}

func (c *ArcadeClient) Username() string {
	return c.AuthInfo.AccountName
}

func (c *ArcadeClient) Nickname() string {
	return c.AuthInfo.Nickname
}

func (c *ArcadeClient) UID() string {
	return c.AuthInfo.UIDStr
}

func (c *ArcadeClient) FormatNameWithUUID() string {
	return fmt.Sprintf("%s(%s)", c.AuthInfo.Nickname, c.AuthInfo.UIDStr)
}
