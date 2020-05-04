//  并发打印

package main

import (
	"fmt"
)

func printer(ch chan int) {
	for {
		data := <-ch
		if data == 10 {
			break
		}
		fmt.Println(data)
	}

	ch <- 0
}

func main() {
	ch := make(chan int)

	go printer(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	ch <- 10
	if 0 == <-ch {
		fmt.Println("End!")
	}
}
