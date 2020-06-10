// demo04
// Using channels

// An aside about buffered channels

// Note for exports: Go channels can also be created with a buffer.
// Buffering removes synchronization.
// Buffering makes them more like Erlang`s mailboxes.
// Buffered channels can be important for some problems 
// but they are more subtle to reason about
// We won`t need them today.

// The Go approach
// Don`t communicate by sharing memory, sharing memory by communicating.
// 不要使用内存共享的方式通信，使用通信来完成内存共享。

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	c := make(chan, string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprint("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}