// demo11
// Timeout for whole conversation using select

// Create the timer once, outside the loop, to time out the entire conversation.
// (In the previous program, we had a timeout for each message.)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1020)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("Joe")
	timeout := time.After(3 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You`re too slow...")
			return
		}
	}
}
