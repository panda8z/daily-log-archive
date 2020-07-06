# Go并发模式

> 资料搜集：
>
> 1. 

### 循环打印



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

### 使用并发模式

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



### 能看到输出的并发模式



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

