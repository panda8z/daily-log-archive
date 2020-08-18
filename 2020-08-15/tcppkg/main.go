package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {

	const TestCount = 100000

	const address = "127.0.0.1:8010"

	var recvCounter int

	acceptor := NewAcceptor()

	acceptor.Start(address)

	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool {
		str := string(data)
		fmt.Println("recive:", str)
		n, err := strconv.Atoi(str)
		if err != nil || recvCounter != n {
			panic("failed")
		}
		recvCounter++

		if recvCounter >= TestCount {
			acceptor.Stop()
			return false
		}
		return true
	}

	Connector(address, TestCount)

	acceptor.Wait()
}
