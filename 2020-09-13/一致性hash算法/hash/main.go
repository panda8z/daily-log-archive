package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type ServiceNode struct {
	Ip   string
	Port string
}

func NewServiceNode(ip, port string) *ServiceNode {
	return &ServiceNode{
		Ip:   ip,
		Port: port,
	}
}

func main() {
	m := md5.New()
	m.Write([]byte("Hello"))

	md5Str := hex.EncodeToString(m.Sum(nil))
	fmt.Println(md5Str)
	fmt.Println("Hello")

	node1 := NewServiceNode("127.0.0.1", "3306")
	node2 := NewServiceNode("127.0.0.1", "3307")
	node3 := NewServiceNode("127.0.0.1", "3308")

	virtualServiceNode.addVirtualNode(node1, 3)
	virtualServiceNode.addVirtualNode(node2, 3)
	virtualServiceNode.addVirtualNode(node3, 3)
	PrintRoutList()
	for i := 0; i < 20; i++ {

		virtualServiceNode.removeVirtualNode(node2, 3)

		serviceNdoe := virtualServiceNode.getVirtualNode("get_cc" + strconv.Itoa(i))

		fmt.Println(i, "   ", serviceNdoe.Ip, ":", serviceNdoe.Port)
	}
	PrintRoutList()
}

func PrintRoutList() {
	for key, node := range virtualServiceNode.VirtualNodes {
		fmt.Println("key=", strconv.Itoa(int(key)), "ip: ", node.Ip, " : ", node.Port)
	}
}
