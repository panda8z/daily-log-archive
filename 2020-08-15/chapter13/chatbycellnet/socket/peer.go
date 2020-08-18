package main

import (
	"fmt"

	"github.com/davyxu/cellnet"
)

type EventFunc func(param interface{}) interface{}

type socketPeer struct {
	eventFunc cellnet.EventFunc
}

func (s *socketPeer) SetEvent(f cellnet.EventFunc) {
	s.eventFunc = f
}

func (s *socketPeer) fireEvent(ev interface{}) interface{} {
	if s.eventFunc == nil {
		return nilfunc
	}

	return s.eventFunc(ev)
}

type Peer interface {
	Start(address string) Peer
	Stop()
	Queue() EventQueue
	SetEvent(f EventFunc)
	Name() string
	SetName(string)
	SessionAccessor
}

func NewConnector() {
	queue := cellnet.NewEventQueue()
	peer := socket.NewConnector(func(raw interface{}) {
		switch ev := raw.(type) {
		case socket.ConnectedEvent:
			fmt.Println("cellnet connected")
		case socket.SessionClosedEvent:
			fmt.Println("client error:", ev.Error)
		}
	}, queue)
}
