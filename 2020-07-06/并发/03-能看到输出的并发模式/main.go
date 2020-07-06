package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("Boring!")
	fmt.Println("I`m Listenning...")
	time.Sleep(2 * time.Second)
	fmt.Println("Too boring! I`m Leaving..")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
