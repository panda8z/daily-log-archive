// p242

package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++

		fmt.Println("tick", times)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	go running()

	var input string

	fmt.Scanln(&input)
	// fmt.Println(input)
}
