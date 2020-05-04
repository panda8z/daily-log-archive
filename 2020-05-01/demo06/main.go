// p256 带缓冲通道的channel

package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 0
	ch <- 1
	ch <- 2

	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
