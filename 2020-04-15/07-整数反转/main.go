package main

import (
	"fmt"
)

func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if y <= -(1<<31) || y >= (1<<31)-1 {
			return 0
		}
		x /= 10
	}
	return y
}

func main() {

	fmt.Println(reverse(2340))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2340))
}
