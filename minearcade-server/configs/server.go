package configs

import "time"

const (
	SERVER_TCP_PORT = 6000
	SERVER_UDP_PORT = 6001
)

const (
	CLIENT_PACKET_BUFSIZE = 256
	SERVER_PACKET_BUFSIZE = 256
)

const UDP_CONNECTION_TIMEOUT = 8 * time.Second
