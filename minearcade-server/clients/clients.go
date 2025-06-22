package clients

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
)

type ArcadeClients struct {
	mu          *sync.Mutex
	tcp_clients map[net.Addr]*ArcadeClient
	udp_clients map[*net.UDPAddr]*ArcadeClient
}

func NewClients() *ArcadeClients {
	clis := &ArcadeClients{
		mu:          &sync.Mutex{},
		tcp_clients: map[net.Addr]*ArcadeClient{},
		udp_clients: map[*net.UDPAddr]*ArcadeClient{},
	}
	return clis
}

func (c *ArcadeClients) AddClient(cli *ArcadeClient) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.tcp_clients[cli.TCPConn.RemoteAddr()] = cli
	c.udp_clients[cli.UDPAddr] = cli
}

func (c *ArcadeClients) RemoveClient(cli *ArcadeClient) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.tcp_clients, cli.TCPConn.RemoteAddr())
	delete(c.udp_clients, cli.UDPAddr)
}

func (c *ArcadeClients) HandoutUDPBytePacket(bs []byte, udp_addr *net.UDPAddr) {
	cli, ok := c.udp_clients[udp_addr]
	if !ok {
		ip := udp_addr.IP
		// todo: 假设同一设备上仅有 1 个客户端连接到服务器.
		// future: 更改逻辑为接受带有 uniqueID 的 UDP 包以确认此 UDP 包来源于哪个客户端
		cli, ok = c.getClientByIP(ip)
		if !ok {
			slog.Error(fmt.Sprintf("No such udp client: %v", udp_addr))
			return
		}
	}
	cli.SetUDPAddr(udp_addr)
	go cli.PacketReader.ReceiveUDPBytePacket(bs)
}

func (c *ArcadeClients) getClientByIP(ip net.IP) (*ArcadeClient, bool) {
	for _, cli := range c.tcp_clients {
		if cli.TCPConn.RemoteAddr().(*net.TCPAddr).IP.Equal(ip) {
			return cli, true
		}
	}
	return nil, false
}
