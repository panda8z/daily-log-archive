// p250  接收任意数据忽略接收的数据格式

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("start gorutine")
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	<-ch
	fmt.Println("all done")
}
