# 怎样书写 Go 语言程序？（[How to Write Go Code ](https://golang.google.cn/doc/code.html)）



## 1. 简述


- 代码的组织方式。
  - 使用 package 组织代码，中文叫做 包。
  - 一个仓库包含一个或多个 module。
    - 仓库是指类似github之类的仓库
    - module 是多个有联系的 go package 的集合。
- 第一个Go语言程序。
  - 新建一个 hello 文件夹。
  - 使用 `go mod init example.com/user/hello`命令创建一个 go.mod 文件。
  - 查看一下 go.mod 文件，使用 `cat go.mod`命令。
  - go源码文件的第一行有效代码必须是 `package name`, 这叫包声明。
- 在同一个 module 里倒入另一个 package
  - hello 文件夹下新建 reverse.go 文件
  - 在 main.go 中 使用 `import "example.com/user/hello/morestrings"` 语句导入这个文件就能使用这个包的方法了。

## 2. 一步一步开始书写 Go 程序

新建文件夹 `ch01`。 然后在 `ch01` 内新建文件夹 `demo01`。最后在 `demo01` 内文件 `main.go`。
最后在 `main.go` 里书写一下内容（注意所有标点都是英文半角）。

**Code·1-1**
```go
package main

import (
  "time"
  "fmt"
)

func main() {
  fmt.Println("Hello Go!")
  fmt.Println("现在时间是: ",time.Now())
}
```


## 3. 代码解释

如果你已经完成了上一步，恭喜你🎉！ 你已经完成了第一个 Go语言程序的源代码书写。

接下来我们来逐行解释，每一行代码的意义和执行后的表现。

**Code·1-1**
```go
// 包声明： 我们的程序用 main 包 组织起来。（同一个项目下，即使有多个 .go 文件只要源代码的包声明是 “main” 就是被 main包组织起来的。 ）
package main

// 导入声明： 我们的程序代码一部分细节已经被标准库实现，我们只需要导入相应的实现就可以使用了。
import (
  "time" // 导入了 标准库 的 time 包，它主要实现了 时间 的格式化输出和打印。
  "fmt"  // 导入了 标准库 的 fmt 包， 它主要实现了标准输出的格式化。
)

// main 函数： go程序的入口
func main() {
  fmt.Println("Hello Go!")  // 在 控制台 打印 “Hello Go!”，然后换行。
  fmt.Println("现在时间是: ",time.Now()) //在控制台打印：“现在时间是: 【当前时间】” 
  //例如： “现在时间是: Mon Jan 2 15:04:05 -0700 MST 2006” 
  //(这里的时间是每次运行程序的时间哦，因为代码里我们使用了 time.Now() 函数获取的是程序执行时的时间，这个是动态的。)
}
```

## 4. 运行程序 

打开命令行，在 demo01 路径下 执行命令 `go run main.go`

这是 `go` 命令工具会编译并运行 `main.go`，结果打印如下。

```bash
Hello Go!
现在时间是:  2020-05-19 00:12:02.865433 +0800 CST m=+0.000257320

```

## 5. 关键词释疑

- 标准库
- 标准输出
- main 包组织起来
- 函数
- 导入
- go 程序入口