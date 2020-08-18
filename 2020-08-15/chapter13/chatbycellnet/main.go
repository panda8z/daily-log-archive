package main

import "net"

func main() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		c.fireEvent(ConnectErrorEvent{ses, err})
		return
	}
}
