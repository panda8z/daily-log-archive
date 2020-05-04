// p261  延迟回调

package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan int)

	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		exit <- 0
	})

	<-exit
}
