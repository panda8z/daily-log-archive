// Fan-in using select

// Rewrite our original fanIn function.
// Only one goroutine is needed.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message msg wait
type Message struct {
	msg  string
	wait chan bool
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	fmt.Printf("input 1:%v \nInput2: %v\n", input1, input2)
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitFor := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s,%d", msg, i), waitFor}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitFor
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.msg)
		msg2 := <-c
		fmt.Println(msg2.msg)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You`re both boring; I`m leaving")

}
