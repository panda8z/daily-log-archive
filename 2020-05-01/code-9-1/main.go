// p251 使用for从通道中循环接收

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. create the channel.
	ch := make(chan int)

	// 2. start a goroutine.
	go func() {

		// 3. send four times channel, sleep for every times.
		for i := 3; i >= 0; i-- {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	// 4. use for-range statement to takeout the data in channel.
	for data := range ch {
		fmt.Println(data)
		// 5. the  important break
		// 接收和发送都是阻塞的 接收不到也会 ：	all goroutines are asleep - deadlock!
		if data == 0 {
			break
		}
	}
}
