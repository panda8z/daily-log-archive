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

		var skip int
		for {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			fmt.Printf("%s%d\n", file, line)
			skip++
		}
		wg.Done()
	}()
	wg.Wait()
}

/* output:

/Users/zcj/panda/git4me/daily-log-archive/2020-04-05/demo01/main.go16
/usr/local/go/src/runtime/asm_amd64.s1373

*/
