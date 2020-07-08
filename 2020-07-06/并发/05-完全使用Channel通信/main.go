package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Boring!")
	for i := 0; ; i++ {
		fmt.Printf("You say :%q\n", <-c)
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
