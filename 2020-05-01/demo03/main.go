// p244  使用匿名函数 创建goroutine

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)

			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
