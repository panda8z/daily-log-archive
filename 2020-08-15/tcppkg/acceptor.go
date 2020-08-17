package main

import (
	"fmt"
	"net"
	"sync"
)

// Acceptor 接收器
type Acceptor struct {
	// 保存侦听器
	l net.Listener
	// 侦听器的停止同步
	wg sync.WaitGroup
	// 连接的数据回调
	OnSessionData func(net.Conn, []byte) bool
}

// Listen .侦听
func (a *Acceptor) Listen(address string) {
	// 等待组 +1
	a.wg.Add(1)

	// 侦听结束 完成该等待项
	defer a.wg.Done()

	// 准备一个err
	var err error

	// 使用socket开启侦听
	a.l, err = net.Listen("tcp", address)

	// 侦听发生错误,打印错误并退出服务
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 循环侦听
	for {

		// 新连接没有到来时, Accept 是阻塞的.
		conn, err := a.l.Accept()

		// 如发生任何侦听错误,打印错误并退出服务
		if err != nil {
			break
		}
		// 根据连接开启绘画,这个过程需要并行执行.
		go handleSession(conn, a.OnSessionData)
	}
}

// Stop .停止侦听
func (a *Acceptor) Stop() {
	a.l.Close()
}

// Wait .等待侦听完全停止.
func (a *Acceptor) Wait() {
	a.wg.Wait()
}

// Start .开始异步侦听
func (a *Acceptor) Start(address string) {
	go a.Listen(address)
}

// NewAcceptor .实例化接收器
func NewAcceptor() *Acceptor {
	return &Acceptor{}
}
