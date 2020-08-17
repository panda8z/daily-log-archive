package main

import (
	"fmt"
	"net"
	"strconv"
)

// Connector 连接器,传入连接地址和发送封包的次数
func Connector(address string, sendTimes int) {

	// 尝试用socket 连接地址
	// 这里得到了 conn 直接调用 conn的Write方法就能把数据发送出去了.
	conn, err := net.Dial("tcp", address)

	// 错误发生时退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 循环指定次数
	for i := 0; i < sendTimes; i++ {
		// 将循环序号转为 字符串
		str := strconv.Itoa(i)

		// 发送字符串封包
		if err := writePacket(conn, []byte(str)); err != nil {
			fmt.Println(err)
			break
		}
	}
}
