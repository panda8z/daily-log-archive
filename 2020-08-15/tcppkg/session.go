package main

import (
	"bufio"
	"net"
)

// 连接的会话逻辑
func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	// 创建一个 socket 的 读取器
	dataReader := bufio.NewReader(conn)

	// 循环接受数据
	for {
		// 从连接中读取封包
		pkt, err := readPacket(dataReader)
		// 回调到外部
		if err != nil || !callback(conn, pkt.Body) {
			// 回调要求退出
			conn.Close()
			break
		}
	}
}
