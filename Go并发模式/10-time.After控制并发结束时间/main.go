// demo10：三类并发模式，时间控制的并发，使用 select + time.After 控制并发时间
// Timeout using select

// The time.After function returns a channel that blocks
// for the specified duration.
// After the interval, the channel delivers the current time, once.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
 c := make(chan string)
 go func() {
  for i := 0; ; i++ {
   c <- fmt.Sprintf("%s: %d", msg, i)
   time.Sleep(time.Duration(rand.Intn(1020)) * time.Millisecond)
  }
 }()
 return c
}

func main() {
 c := boring("Joe")
 for {
  select {
  case s := <-c:
   fmt.Println(s)
  case <-time.After(1 * time.Second):
   fmt.Println("You`re too slow...")
   return
  }
 }
}
