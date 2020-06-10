**阅读Go语言官方博客系列**

# Go语言并发模式

这是**阅读Go语言官方博客系列**的第一篇文章。Go语言官方的内容特别多，并且很有指向性的解决了我们遇到的大多数问题。
本文的内容是上个月在油管看了一个2012年的 Go 团队在 Google I/O 的演讲视频，然后在B站找到了这个视频的熟肉。有趣的是我截屏整理后才发现官方公布了 **幻灯片**。

另外，本文仅做整理，无教程性内容。

好了，让我们进入正题。

## 资料搜集

 原视频是英文演讲。

 B站的视频是国内 Up 主 [Capricornwqh](https://space.bilibili.com/296855068)上传的熟肉。

 幻灯片是 Go 语言官方，在 go-blog 发布的 slide。

- 原视频 ： [Google I/O 2012 - Go Concurrency Patterns - YouTube](https://www.youtube.com/watch?v=f6kdp27TYZs)（需科学上网）

- 国内翻译视频：[Go Concurrency Patterns 中文字幕 bilibili](https://www.bilibili.com/video/BV1UJ411m7U1)
- 幻灯片：[Go Concurrency Patterns-concurrency.slide](https://talks.golang.org/2012/concurrency.slide#1)（需科学上网）

## 幻灯片截屏整理

（多图预警）

![Screen Shot 2020-04-26 at 09.57.44](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf9q97n5l2j31hl0u0wma.jpg)

![Screen Shot 2020-04-26 at 10.01.38](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e1ciabj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 10.03.14](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf9q9ftvd2j31hl0u0wma.jpg)

![Screen Shot 2020-04-26 at 10.04.55](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9hj1pjgj31b70u0nbr.jpg)

![Screen Shot 2020-04-26 at 10.05.16](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9idro6mj31ed0u01kx.jpg)

![Screen Shot 2020-04-26 at 10.05.58](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9imq732j31c00u0gq3.jpg)

![Screen Shot 2020-04-26 at 10.06.26](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9fxqklej31c00u07i5.jpg)

![Screen Shot 2020-04-26 at 10.08.07](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e7llsjj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 10.09.40](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9ja2k4gj31b80u0b29.jpg)

![Screen Shot 2020-04-26 at 10.09.52](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9jwbrhzj31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 10.24.33](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e5iw7nj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 10.26.57](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9fqolaaj31c00u07i5.jpg)

![Screen Shot 2020-04-26 at 10.28.42](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9kqyi7xj31b70u07ls.jpg)

![Screen Shot 2020-04-26 at 10.29.27](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9l6uth9j31b70u07ls.jpg)

![Screen Shot 2020-04-26 at 10.31.04](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9m2tqi4j31c00u07ib.jpg)

![Screen Shot 2020-04-26 at 10.35.10](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9mf4g2zj31c00u0wqw.jpg)

![Screen Shot 2020-04-26 at 10.36.41](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf9q9eda86j31hl0u0wma.jpg)

![Screen Shot 2020-04-26 at 10.37.00](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9r7ssjgj31b70u0wrb.jpg)

![Screen Shot 2020-04-26 at 10.38.40](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9q2x2ivj31bw0u07hz.jpg)

![Screen Shot 2020-04-26 at 10.39.29](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e43vyrj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 10.39.57](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9wuomv6j31c00u07i5.jpg)

![Screen Shot 2020-04-26 at 10.42.44](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9xatfckj31b70u0gyf.jpg)

![Screen Shot 2020-04-26 at 10.43.17](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9k15eyfj31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 10.45.04](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9x2tmnoj31b70u0wtz.jpg)

![Screen Shot 2020-04-26 at 10.46.35](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9jwf0frj31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 10.48.37](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e7uwszj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 10.48.46](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9kqnj5pj31b70u07ls.jpg)

![Screen Shot 2020-04-26 at 10.51.44](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9g2wj95j31c00u018e.jpg)

![Screen Shot 2020-04-26 at 10.52.14](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9wbmthwj31b70u0aq6.jpg)

![Screen Shot 2020-04-26 at 10.54.26](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9jxyl13j31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 10.55.05](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9m7z915j31cz0u0jy8.jpg)

![Screen Shot 2020-04-26 at 11.01.17](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9q2h9wjj31bw0u0ncx.jpg)

![Screen Shot 2020-04-26 at 11.02.28](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9r8ut18j31b70u0wrb.jpg)

![Screen Shot 2020-04-26 at 11.02.57](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9x6x9duj31c00u07i5.jpg)

![Screen Shot 2020-04-26 at 11.03.28](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9wckyxgj31b70u0aq6.jpg)

![Screen Shot 2020-04-26 at 11.04.10](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9q30553j31bw0u07hz.jpg)

![Screen Shot 2020-04-26 at 11.05.25](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9kyqfqzj31c00u0tke.jpg)

![Screen Shot 2020-04-26 at 11.07.34](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9g3bgjdj31c00u018e.jpg)

![Screen Shot 2020-04-26 at 11.08.35](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9e91kszj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 11.09.57](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9g03csvj31cl0u048m.jpg)

![Screen Shot 2020-04-26 at 11.11.23](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9l7xxbgj31b70u017d.jpg)

![Screen Shot 2020-04-26 at 11.12.12](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9mln10sj31b70u017d.jpg)

![Screen Shot 2020-04-26 at 11.20.06](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf9q9fhxjaj31hl0u0wma.jpg)

![Screen Shot 2020-04-26 at 11.20.14](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9ee46mrj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 11.21.05](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9k0msxjj31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 11.22.03](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9f6s2ggj31b70u0arr.jpg)

![Screen Shot 2020-04-26 at 11.23.23](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9mk8zpoj31c00u0gwn.jpg)

![Screen Shot 2020-04-26 at 11.25.04](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9k0gu0nj31cp0u010u.jpg)

![Screen Shot 2020-04-26 at 11.25.42](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf9q9g9xqej31hl0u0wma.jpg)

![Screen Shot 2020-04-26 at 11.26.06](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9ef5rbnj31tl0u0gsw.jpg)

![Screen Shot 2020-04-26 at 11.29.36](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfa9mmu7zgj31b70u017d.jpg)
## 代码整理

幻灯片结束之后，我们来整理代码

#### demo01

```go
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

```

#### demo02

```go
// demo02
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	go boring("boring!")
}

func boring(msg string)  {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
```


#### demo03

```go
// demo03
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("boring!")

	fmt.Println("I`m Listening...")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring! I`m leaving!")

}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

```

#### demo04

```go
// demo04
// Using channels

// An aside about buffered channels

// Note for exports: Go channels can also be created with a buffer.
// Buffering removes synchronization.
// Buffering makes them more like Erlang`s mailboxes.
// Buffered channels can be important for some problems 
// but they are more subtle to reason about
// We won`t need them today.

// The Go approach
// Don`t communicate by sharing memory, sharing memory by communicating.
// 不要使用内存共享的方式通信，使用通信来完成内存共享。

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	c := make(chan, string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprint("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
```

#### demo05

```go
// demo05
// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

```

#### demo06

```go
// demo06
// Channels as a handle on service

// Our boring function returns a channel that let us communicate with the boring service it provides.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You are boring! I`m leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

```

#### demo07

```go
// demo 07
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

#### demo08

```go
// demo08
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

#### demo09

```go
// demo09
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

#### demo10

```go
// demo10
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

#### demo11

```go
// demo11
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

#### demo12

```go
// demo12
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

#### demo13

```go
// demo13
// Recive on quit channel

// How do we Konw it`s finished?
// Wait for it to tell us it`s done: recive onthe quit channel

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				// fmt.Println("boring: Oh! You Stop Me!")
				quit <- "See you!"
				return
			}

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("clean somthings")
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

```

#### demo14

```go
// demo14
// Daisy-chain

package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost 
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left,right)
		left = right
	}
	go func (c chan int)  {
		c <- 2
	}(right)
	fmt.Println(<-leftmost)
}
```

#### google-search-00

```go
// demo - Google Search 00
// Google Search: A fake framework
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

// Result type
type Result string

// Search fake
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google ,
func Google(query string) Result {
	return Result(fmt.Sprintf("%s%s%s", web(query), image(query), video(query)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result, elapsed)

}

```

#### google-search-01

```go
// demo - Google Search 01
// Google Search 1.0
// Synchronous
// The Google function takes a query and returns a slice of Results(which are just strings)
// Google invokes Web, Image and Video searches serially, appending them to the results slice.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

// Result type
type Result string

// Search fake
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google ,
func Google(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}

```

#### google-search-02

```go
// demo - Google Search 02
// Google Search 2.0
// Run the Web, Image and Video searches concurrently, and wait for all results.
// No locks, No condition variables, No callbacks
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

// Result type
type Result string

// Search fake
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google ,
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- web(query)
	}()
	go func() {
		c <- image(query)
	}()
	go func() {
		c <- video(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}

```

#### google-search-03

```go
// demo - Google Search 03
// Google Search 2.1
// Don`t wait for slow servers.
// No locks, No condition variables, No callbacks
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

// Result type
type Result string

// Search fake
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google ,
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- web(query)
	}()
	go func() {
		c <- image(query)
	}()
	go func() {
		c <- video(query)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {

		select{
			case result := <-c
			results = append(results, result)
			case <-timeout
			fmt.Println("time out")
			return 
		}
		
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}

```

#### google-search-04

```go
// demo - Google Search 04
// Google Search 3.0

// Q: How do we avoid discarding results from slow servers?
// A: Replicate the servers. Send requests to multiple replicas, and use the first response.

// Reduce tail latency using replicated search servers.(使用复制的搜索服务器减少尾部延迟。)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = fakeSearch("web")
	web2   = fakeSearch("web")
	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
)

// Result 定义返回值
type Result string

// Search 搜索服务
type Search func(query string) Result

// fakeSearch 虚拟的搜索服务
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// First  query: 查询字串  replicas: 服务器的复制
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// Google 聚合的Google搜索服务
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- First(query, web1, web2)
	}()
	go func() {
		c <- First(query, image1, image2)
	}()
	go func() {
		c <- First(query, video1, video2)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

```

[Panda张向北](https://github.com/panda8z)© 2018-2020 版权所有， 采用[知识共享署名-非商业性使用-禁止演绎 4.0 国际许可协议许可](http://creativecommons.org/licenses/by-nc-nd/4.0/)，代码使用 [MIT](https://opensource.org/licenses/MIT) 协议开源。