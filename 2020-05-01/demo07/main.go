// p264 给被关闭的channel发送数据将会触发 panic

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	close(ch)

	fmt.Printf("ptr:%p, cap:%d, len:%d\n", ch, cap(ch), len(ch))

	ch <- 1
}
