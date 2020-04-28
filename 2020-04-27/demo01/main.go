package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		select {}
	}()

	fmt.Println("boring!  I`m leaving.")
}
