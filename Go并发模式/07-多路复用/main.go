// demo07： 二类并发模式，顶配版，多路复用
// Multiplexing

package main

import "fmt"

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You`re both boring; I`m leaving")
}

// fanIn
// input1\input2是两个仅接收channel
// fanIn 返回值也是一个仅接受Channel
// 拿出 仅接受Channel 的值 的方法是  <-[仅接收channel变量名]
func fanIn(input1, input2 <-chan string) <-chan string {
	fmt.Printf("input 1:%v \nInput2: %v\n", input1, input2)
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}
