package main

import (
	"github.com/davyxu/cellnet"
)

type RecvEvent struct {
	Ses cellnet.Session
}

type SendEvent struct {
	Ses cellnet.Session
	Msg interface{}
}

type RecvErrorEvent struct {
	Ses   cellnet.Session
	Error error
}

type SendErrorEvent struct {
	Ses   cellnet.Session
	Error error
	Msg   interface{}
}

type SessionStartEvent struct {
	Ses cellnet.Session
}

type ConnectedEvent = SessionStartEvent

type ConnectErrorEvent struct {
	Ses   cellnet.Session
	Error error
}

type AcceptedEvent = SessionStartEvent

type SessionClosedEvent struct {
	Ses   cellnet.Session
	Error error
}
