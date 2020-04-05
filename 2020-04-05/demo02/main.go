package main

import (
	"sync"
	"fmt"
)
func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		fmt.Println("goroutine is running...")
		wg.Done()
	}()
	fmt.Println("main is running...")
	wg.Wait()
}