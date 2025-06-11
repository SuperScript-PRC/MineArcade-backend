package server

import (
	"MineArcade-backend/minearcade-server/configs"
	"fmt"
	"log/slog"
	"net"
)

type MineArcadeServer struct {
	TCPListener            net.Listener
	UDPListener            *net.UDPConn
	ClientConnHandler      func(tcp_conn net.Conn, udp_conn *net.UDPConn)
	ClientUDPPacketHandler func(packet_data []byte, udp_addr *net.UDPAddr)
}

func NewServer() *MineArcadeServer {
	return &MineArcadeServer{}
}

func (s *MineArcadeServer) SetClientUDPPacketHandler(hdl func(packet_data []byte, udp_addr *net.UDPAddr)) {
	s.ClientUDPPacketHandler = hdl
}

func NewServerWithConnHandler(hdl func(tcp_conn net.Conn, udp_conn *net.UDPConn)) *MineArcadeServer {
	s := &MineArcadeServer{}
	s.SetConnHandler(hdl)
	return s
}

func (s *MineArcadeServer) SetConnHandler(handler func(tcp_conn net.Conn, udp_conn *net.UDPConn)) {
	s.ClientConnHandler = handler
}

func (s *MineArcadeServer) StartServer() {
	if s.ClientConnHandler == nil {
		slog.Error("未使用 SetConnHandler() 设置连接处理方法")
		return
	}
	tcp_listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.SERVER_TCP_PORT))
	if err != nil {
		slog.Error(fmt.Sprintf("tcp_server open error: %v", err))
		return
	}
	s.TCPListener = tcp_listener
	udp_listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: configs.SERVER_UDP_PORT})
	if err != nil {
		slog.Error(fmt.Sprintf("udp_server open error: %v", err))
		return
	}
	s.UDPListener = udp_listener
	slog.Info("MineArcade-backend 已启动")
	go s.tcpServerEntry()
	go s.udpServerEntry()
}

func (s *MineArcadeServer) tcpServerEntry() {
	for {
		tcp_conn, err := s.TCPListener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				slog.Error(fmt.Sprintf("TCPServerEntry Accept() error: %v", err))
				break
			} else {
				slog.Error(fmt.Sprintf("TCPServerEntry Accept() error: %v", err))
			}
		}

		slog.Info(fmt.Sprintf("新连接: %v", tcp_conn.RemoteAddr().String()))
		go s.ClientConnHandler(tcp_conn, s.UDPListener)
	}
}

func (s *MineArcadeServer) udpServerEntry() {
	for {
		bs := make([]byte, 65536)
		bs_len, udp_addr, err := s.UDPListener.ReadFromUDP(bs)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				slog.Error(fmt.Sprintf("UDPServerEntry ReadFromUDP() error: %v", err))
				break
			} else {
				slog.Error(fmt.Sprintf("UDPServerEntry ReadFromUDP() error: %v", err))
			}
		} else {
			bs := bs[:bs_len]
			s.ClientUDPPacketHandler(bs, udp_addr)
		}
	}

}
