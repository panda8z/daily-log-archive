// demo01: 	

package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	boring("boring!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
