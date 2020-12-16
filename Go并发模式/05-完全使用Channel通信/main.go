// demo05：二类并发模式，函数返回 channel
// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
