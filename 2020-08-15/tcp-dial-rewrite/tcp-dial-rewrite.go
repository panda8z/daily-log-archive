/*
优化并发同步使用等待组

src/runtime/chan.go 中,经过分析channel 也是通过常用的互斥量进行同步.
所以我们改用效率更高的传统 互斥量 等待组来实现这个简单的同步. sync.WaitGroup

*/
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	buf := make([]byte, 1024)

	for {
		a, err := conn.Read(buf)
		fmt.Println(a)
		if err != nil {
			break
		}
	}

	wg.Done()
}

func main() {

	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建等待组
	var wg sync.WaitGroup

	// 等待组 + 1
	wg.Add(1)

	go socketRecv(conn, &wg)

	time.Sleep(time.Second)

	conn.Close()

	// 等待组会等阻塞至所有 并发都 done()了之后才会继续执行下面的代码.
	wg.Wait()
	fmt.Println("recv done!")
}
