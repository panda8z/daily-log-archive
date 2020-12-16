# Go并发模式

这是**阅读Go语言官方博客系列**的第一篇文章。Go语言官方的内容特别多，并且很有指向性的解决了我们遇到的大多数问题。 本文的内容是上个月在油管看了一个2012年的 Go 团队在 Google I/O 的演讲视频，然后在B站找到了这个视频的熟肉。有趣的是我截屏整理后才发现官方公布了 **幻灯片**。

另外，本文仅做整理，无教程性内容。

好了，让我们进入正题。

### 资料搜集

原视频是英文演讲。



幻灯片是 Go 语言官方，在 go-blog 发布的 slide。

- 原视频 ： [Google I/O 2012 - Go Concurrency Patterns - YouTube](https://www.youtube.com/watch?v=f6kdp27TYZs)（需科学上网）

- 国内翻译视频：[Go Concurrency Patterns 中文字幕 bilibili](https://www.bilibili.com/video/BV1UJ411m7U1) 

  B站的视频是国内 Up 主 [Capricornwqh](https://space.bilibili.com/296855068)上传的熟肉。

- 幻灯片：[Go Concurrency Patterns-concurrency.slide](https://talks.golang.org/2012/concurrency.slide#1)（需科学上网）



### 01-循环打印

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	 boring("Boring!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond
	}
}
```

这是一个使用 for循环实现的循环打印，因为没有设置退出条件，将一直在for循环执行打印操作，i会越来越大。

### 02-使用并发模式

```go
package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	go boring("Boring!")
}

func boring(msg string) {
	for i:=0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
```



这个例子我们看不到任何输出的情况下程序就结束了。

这里我们使用了go关键字，开始使用go并发模式。

main函数执行完毕后，整个程序也就执行完毕了，并发也跟着结束了。



### 03-能看到输出的并发模式



```go
package main 

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	go boring("Boring!")
	fmt.Println("I`m Listenning...")
	time.Sleep(2 * time.Second)
	fmt.Println("Too boring! I`m Leaving..")
}

func boring(msg string) {
	for i:=0; ; i++ {
		fmt.Println(msg,i)
		timg.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
```

输出如下：

```
I`m Listenning...
Boring! 0
Boring! 1
Boring! 2
Boring! 3
Boring! 4
Boring! 5
Too boring! I`m Leaving..
```

并发开始后两秒内，打印了6次，最后main函数结束执行，整个程序结束，并发也就结束了。



### 使用Channel的并发模式



```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!",c)
	for i:=0; i<5; i++ {
			fmt.Printf("You say :%s",<-c)
	}
	fmt.Println("Yes! You are boring! I`m leaving")
}

func boring(msg string, c chan string) {
	for i:=0; ; i++ {
		c <- fmt.Sprintf("%s %d",msg,i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
```

本例中使用channel 这个数据类型做了并发过程中的数据传递。

Don`t communicate by sharing memory, sharing memory by communicating.

### 05-完全改用channel通信

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := borring("Boring!")
	for i := 0; i<5; i++ {
		fmt.Printf("You say :%q\n",<-c)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func(){
		for i := 0; ;i++ {
			c <- fmt.Sprintf("%s %d", msg,i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}
```

### 06-两路并发



```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe!")
	yanny := boring("Yanny!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-yanny)
	}
	fmt.Println("You are boring. I`m leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func(){
		for i:=0;;i++ {
			c <- fmt.Sprintf("%s %d",msg,i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
  return c
}
```

### 07-多路复用

```go
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

```



### 08-回环序列


```go
// demo08： 二类并发模式，回环序列
// Reastoring sequencing
package main

import (
 "fmt"
 "math/rand"
 "time"
)

// Message msg wait
type Message struct {
 msg  string
 wait chan bool
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
 fmt.Printf("input 1:%v \nInput2: %v\n", input1, input2)
 c := make(chan Message)
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

func boring(msg string) <-chan Message {
 c := make(chan Message)
 waitFor := make(chan bool)
 go func() {
  for i := 0; ; i++ {
   c <- Message{fmt.Sprintf("%s,%d", msg, i), waitFor}
   time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
   <-waitFor
  }
 }()
 return c
}

func main() {
 c := fanIn(boring("Joe"), boring("Ann"))
 for i := 0; i < 5; i++ {
  msg1 := <-c
  fmt.Println(msg1.msg)
  msg2 := <-c
  fmt.Println(msg2.msg)
  msg1.wait <- true
  msg2.wait <- true
 }
 fmt.Println("You`re both boring; I`m leaving")

}
```


### 09- // demo09：二类并发，使用 select 语句优化的多路复用

```go
// demo09：二类并发，使用 select 语句优化的多路复用
// Fan-in using select

// Rewrite our original fanIn function.
// Only one goroutine is needed.

package main

import (
 "fmt"
 "math/rand"
 "time"
)

// Message msg wait
type Message struct {
 msg  string
 wait chan bool
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
 fmt.Printf("input 1:%v \nInput2: %v\n", input1, input2)
 c := make(chan Message)
 go func() {
  for {
   select {
   case s := <-input1:
    c <- s
   case s := <-input2:
    c <- s
   }
  }
 }()
 return c
}

func boring(msg string) <-chan Message {
 c := make(chan Message)
 waitFor := make(chan bool)
 go func() {
  for i := 0; ; i++ {
   c <- Message{fmt.Sprintf("%s,%d", msg, i), waitFor}
   time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
   <-waitFor
  }
 }()
 return c
}

func main() {
 c := fanIn(boring("Joe"), boring("Ann"))
 for i := 0; i < 5; i++ {
  msg1 := <-c
  fmt.Println(msg1.msg)
  msg2 := <-c
  fmt.Println(msg2.msg)
  msg1.wait <- true
  msg2.wait <- true
 }
 fmt.Println("You`re both boring; I`m leaving")

}
```


### 10-// demo10：三类并发模式，时间控制的并发，使用 select + time.After 控制并发时间

```go

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

```

### 11-// demo11：三类并发模式，时间控制的并发，控制整个并发会话

```go
// demo11：三类并发模式，时间控制的并发，控制整个并发会话
// Timeout for whole conversation using select

// Create the timer once, outside the loop, to time out the entire conversation.
// (In the previous program, we had a timeout for each message.)
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
 timeout := time.After(3 * time.Second)
 for {
  select {
  case s := <-c:
   fmt.Println(s)
  case <-timeout:
   fmt.Println("You`re too slow...")
   return
  }
 }
}
```

### 12-// demo12: 三类并发模式，使用 select + channel 控制并发


```go
// demo12: 三类并发模式，使用 select + channel 控制并发
// Qiut Channel

package main

import (
 "fmt"
 "math/rand"
 "time"
)

func boring(msg string, quit chan bool) <-chan string {
 c := make(chan string)
 go func() {
  for i := 0; ; i++ {
   select {
   case c <- fmt.Sprintf("%s: %d", msg, i):
    // do nothing
   case <-quit:
    fmt.Println("Oh! You Stop Me!")
    return
   }

   time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
  }
 }()
 return c
}

func main() {
 quit := make(chan bool)
 c := boring("Joe", quit)
 for i := rand.Intn(10); i >= 0; i-- {
  fmt.Println(<-c)
 }
 quit <- true
}

```