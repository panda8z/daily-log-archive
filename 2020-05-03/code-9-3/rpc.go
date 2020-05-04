// p258 模拟RPC
package main

import (
	"fmt"
	"time"
	"errors"
	"math/rand"
)

// RPCClient ，
func RPCClient(ch chan string, req string) (string, error) {
	ch <- req 
	select {
	case ack := <- ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

// RPCServer ，
func RPCServer(ch chan string)  {
	for {
		data := <-ch
		fmt.Println("server recevied:", data)
		time.Sleep(time.Duration(rand.Intn(3))* time.Millisecond * 1000)
		ch <- "roger"
	}
}

func main()  {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received:", recv)
	}
}