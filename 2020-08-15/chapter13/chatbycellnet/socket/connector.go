package main

import (
	"net"

	"github.com/davyxu/cellnet"
)

type socketConnector struct {
	socketPeer
	internal.SessionManager
	ses cellnet.Session
}

func (s *socketConnector) Start(address string) cellnet.Peer {
	a.address = address
	go c.connect(address)
	return c
}

func (s *socketConnector) connect(address string) {
	conn, err := net.Dial("tcp", address)

	ses := newSession(conn, &c.socketPeer)
}
