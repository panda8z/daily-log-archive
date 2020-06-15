package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("%#b %#v\n",9>>1, 9>>1)
	fmt.Println(maxDepth(100))
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}