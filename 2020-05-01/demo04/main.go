// p245 channel 发送必须有接受  发送将持续阻塞，知道数据被接收。
package main

func main() {
	ch := make(chan int)

	ch <- 0
}
