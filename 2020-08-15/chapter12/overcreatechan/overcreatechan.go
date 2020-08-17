package main

import (
	"fmt"
	"runtime"
)

func consumer(ch chan int) {
	for {
		data := <-ch
		fmt.Println(data)
	}
}

func main() {
	ch := make(chan int)

	for {
		var dummy string

		fmt.Scan(&dummy)

		go consumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
