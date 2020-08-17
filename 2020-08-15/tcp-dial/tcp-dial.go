/*

通道 Channel 在Go源码内实现,为了保证goroutine的并发安全,也使用了一些锁操作.
因此 channel 其实并不比锁高效.

下面我们实现一个用 channel 做 接收同步的tcp链接.

*/

package main

import (
	"fmt"
	"net"
	"time"
)

func socketRecv(conn net.Conn, exitChan chan string) {
	buff := make([]byte, 1024)

	for {
		content, err := conn.Read(buff)
		fmt.Println(content)
		if err != nil {
			break
		}
	}

	exitChan <- "recv exit!"
}


func main()  {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
		return 
	}

	exit := make(chan string )

	go socketRecv(conn, exit)

	time.Sleep(time.Second  * 3)

	conn.Close()

	fmt.Println(<-exit)
}
