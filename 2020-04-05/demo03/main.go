package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.Goexit() // goroutine exits here
		fmt.Println("this line code never executed")
	}()
	wg.Wait()
}

