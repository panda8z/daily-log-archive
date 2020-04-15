## Golang 之轻松化解 defer 的温柔陷阱

[CSDN](javascript:void(0);) *2019-03-19*

以下文章来源于码农桃花源 ，作者饶全成

<img src="Go中defer/0-20200407091714408.jpeg" alt="码农桃花源" style="zoom:25%;" />



[**码农桃花源**一线互联网工程师死磕技术的心路历程、经验教训。欢迎围观！](https://mp.weixin.qq.com/s/4JZoF5mZ4E1_WQw2ek7NYQ#)

作者 | 饶全成

责编 | 胡巍巍

defer是Go语言提供的一种用于注册延迟调用的机制：让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行。

深受Go开发者的欢迎，但一不小心就会掉进它的温柔陷阱，只有深入理解它的原理，我们才能轻松避开，写出漂亮稳健的代码。

为了更好的阅读体验，按惯例我手动贴上文章目录：

![img](Go中defer/640.jpeg)





#### **1、什么是defer？**



defer是Go语言提供的一种用于注册延迟调用的机制：让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行。

defer语句通常用于一些成对操作的场景：打开连接/关闭连接；加锁/释放锁；打开文件/关闭文件等。

defer在一些需要回收资源的场景非常有用，可以很方便地在函数结束前做一些清理操作。在打开资源语句的下一行，直接一句defer就可以在函数返回前关闭资源，可谓相当优雅。

```
f, _ := os.Open("defer.txt")
defer f.Close()
```

注意：以上代码，忽略了err, 实际上应该先判断是否出错，如果出错了，直接return. 接着再判断 f是否为空，如果 f为空，就不能调用 f.Close()函数了，会直接panic的。



#### **2、为什么需要defer？**



程序员在编程的时候，经常需要打开一些资源，比如数据库连接、文件、锁等，这些资源需要在用完之后释放掉，否则会造成内存泄漏。

但是程序员都是人，是人就会犯错。

因此经常有程序员忘记关闭这些资源。Golang直接在语言层面提供 defer关键字，在打开资源语句的下一行，就可以直接用 defer语句来注册函数结束后执行关闭资源的操作。因为这样一颗“小小”的语法糖，程序员忘写关闭资源语句的情况就大大地减少了。





#### **3、怎样合理使用defer?**



defer的使用其实非常简单：

```
f,err := os.Open(filename)
if err != nil {
    panic(err)
}
if f != nil {
    defer f.Close()
}
```

在打开文件的语句附近，用defer语句关闭文件。这样，在函数结束之前，会自动执行defer后面的语句来关闭文件。

当然，defer会有小小地延迟，对时间要求特别特别特别高的程序，可以避免使用它，其他一般忽略它带来的延迟。

#### 4、 defer的底层原理是什么？

我们先看一下官方对 defer的解释：

> Each time a “defer” statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. If a deferred function value evaluates to nil, execution panics when the function is invoked, not when the “defer” statement is executed.

翻译一下：每次defer语句执行的时候，会把函数“压栈”，函数参数会被拷贝下来；当外层函数（非代码块，如一个for循环）退出时，defer函数按照定义的逆序执行；如果defer执行的函数为nil, 那么会在最终调用函数的产生panic.

defer语句并不会马上执行，而是会进入一个栈，函数return前，会按先进后出的顺序执行。也说是说最先被定义的defer语句最后执行。先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行；否则，如果前面先执行，那后面函数的依赖就没有了。

在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在defer定义时就把值传递给defer，并被cache起来；作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。

defer后面的语句在执行的时候，函数调用的参数会被保存起来，也就是复制了一份。真正执行的时候，实际上用到的是这个复制的变量，因此如果此变量是一个“值”，那么就和定义的时候是一致的。如果此变量是一个“引用”，那么就可能和定义的时候不一致。

举个例子：

```
func main() {
    var whatever [3]struct{}
    for i := range whatever {
        defer func() { 
            fmt.Println(i) 
        }()
    }
}
```

执行结果：

```
2
2
2
```

defer后面跟的是一个闭包（后面会讲到），i是“引用”类型的变量，最后i的值为2, 因此最后打印了三个2.

有了上面的基础，我们来检验一下成果：

```
type number int
func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }
func main() {
    var n number
    defer n.print()
    defer n.pprint()
    defer func() { n.print() }()
    defer func() { n.pprint() }()
    n = 3
}
```

执行结果是：

```
3
3
3
0
```

第四个defer语句是闭包，引用外部函数的n, 最终结果是3; 第三个defer语句同第四个； 第二个defer语句，n是引用，最终求值是3. 第一个defer语句，对n直接求值，开始的时候n=0, 所以最后是0。



**![img](data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==)**

#### **5、利用defer原理**



有些情况下，我们会故意用到defer的先求值，再延迟调用的性质。想象这样的场景：在一个函数里，需要打开两个文件进行合并操作，合并完后，在函数执行完后关闭打开的文件句柄。

```
func mergeFile() error {
    f, _ := os.Open("file1.txt")
    if f != nil {
        defer func(f io.Closer) {
            if err := f.Close(); err != nil {
                fmt.Printf("defer close file1.txt err %v\n", err)
            }
        }(f)
    }
    // ……
    f, _ = os.Open("file2.txt")
    if f != nil {
        defer func(f io.Closer) {
            if err := f.Close(); err != nil {
                fmt.Printf("defer close file2.txt err %v\n", err)
            }
        }(f)
    }
    return nil
}
```

上面的代码中就用到了defer的原理，defer函数定义的时候，参数就已经复制进去了，之后，真正执行close()函数的时候就刚好关闭的是正确的“文件”了，妙哉！可以想像一下如果不这样将f当成函数参数传递进去的话，最后两个语句关闭的就是同一个文件了，都是最后一个打开的文件。

不过在调用close()函数的时候，要注意一点：先判断调用主体是否为空，否则会panic. 比如上面的代码片段里，先判断 f不为空，才会调用 Close()函数，这样最安全。



**![img](data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==)**

#### **6、defer命令的拆解**



如果defer像上面介绍地那样简单（其实也不简单啦），这个世界就完美了。事情总是没这么简单，defer用得不好，是会跳进很多坑的。

理解这些坑的关键是这条语句：

```
return xxx
```

上面这条语句经过编译之后，变成了三条指令：

```
1. 返回值 = xxx
2. 调用defer函数
3. 空的return
```

1,3步才是Return 语句真正的命令，第2步是defer定义的语句，这里可能会操作返回值。

下面我们来看两个例子，试着将return语句和defer语句拆解到正确的顺序。

第一个例子：

```
func f() (r int) {
     t := 5
     defer func() {
       t = t + 5
     }()
     return t
}
```

拆解后：

```
func f() (r int) {
     t := 5
     // 1. 赋值指令
     r = t
     // 2. defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
     func() {        
         t = t + 5
     }
     // 3. 空的return指令
     return
}
```

这里第二步没有操作返回值r, 因此，main函数中调用f()得到5.

第二个例子：

```
func f() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
```

拆解后：

```
func f() (r int) {
     // 1. 赋值
     r = 1
     // 2. 这里改的r是之前传值传进去的r，不会改变要返回的那个r值
     func(r int) { 
          r = r + 5
     }(r)
     // 3. 空的return
     return
}
```

因此，main函数中调用f()得到1。



**![img](data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==)**

#### **7、defer语句的参数**



defer语句表达式的值在定义时就已经确定了。下面展示三个函数：

```GO
func f1() {
    var err error
    defer fmt.Println(err)
    err = errors.New("defer error")
    return
}
func f2() {
    var err error
    defer func() {
        fmt.Println(err)
    }()
    err = errors.New("defer error")
    return
}
func f3() {
    var err error
    defer func(err error) {
        fmt.Println(err)
    }(err)
    err = errors.New("defer error")
    return
}
func main() {
    f1()
    f2()
    f3()
}
```

运行结果：

```BASH
<nil>
defer error
<nil>
```

第1，3个函数是因为作为函数参数，定义的时候就会求值，定义的时候err变量的值都是nil, 所以最后打印的时候都是nil. 第2个函数的参数其实也是会在定义的时候求值，只不过，第2个例子中是一个闭包，它引用的变量err在执行的时候最终变成 defer error了。关于闭包在本文后面有介绍。

第3个函数的错误还比较容易犯，在生产环境中，很容易写出这样的错误代码。最后defer语句没有起到作用。



#### **8、闭包是什么？**



闭包是由函数及其相关引用环境组合而成的实体,即：

```GO
闭包=函数+引用环境
```

一般的函数都有函数名，但是匿名函数就没有。匿名函数不能独立存在，但可以直接调用或者赋值于某个变量。匿名函数也被称为闭包，一个闭包继承了函数声明时的作用域。在Golang中，所有的匿名函数都是闭包。

有个不太恰当的例子，可以把闭包看成是一个类，一个闭包函数调用就是实例化一个类。闭包在运行时可以有多个实例，它会将同一个作用域里的变量和常量捕获下来，无论闭包在什么地方被调用（实例化）时，都可以使用这些变量和常量。而且，闭包捕获的变量和常量是引用传递，不是值传递。

举个简单的例子：

```GO
func main() {
    var a = Accumulator()
    fmt.Printf("%d\n", a(1))
    fmt.Printf("%d\n", a(10))
    fmt.Printf("%d\n", a(100))
    fmt.Println("------------------------")
    var b = Accumulator()
    fmt.Printf("%d\n", b(1))
    fmt.Printf("%d\n", b(10))
    fmt.Printf("%d\n", b(100))
}
func Accumulator() func(int) int {
    var x int
    return func(delta int) int {
        fmt.Printf("(%+v, %+v) - ", &x, x)
        x += delta
        return x
    }
}
```

执行结果：

```BASH
(0xc420014070, 0) - 1
(0xc420014070, 1) - 11
(0xc420014070, 11) - 111
------------------------
(0xc4200140b8, 0) - 1
(0xc4200140b8, 1) - 11
(0xc4200140b8, 11) - 111
```

闭包引用了x变量，a,b可看作2个不同的实例，实例之间互不影响。实例内部，x变量是同一个地址，因此具有“累加效应”。



#### **9、defer配合recover**



Golang被诟病比较多的就是它的error, 经常是各种error满天飞。编程的时候总是会返回一个error, 留给调用者处理。如果是那种致命的错误，比如程序执行初始化的时候出问题，直接panic掉，省得上线运行后出更大的问题。

但是有些时候，我们需要从异常中恢复。比如服务器程序遇到严重问题，产生了panic, 这时我们至少可以在程序崩溃前做一些“扫尾工作”，如关闭客户端的连接，防止客户端一直等待等等。

panic会停掉当前正在执行的程序，不只是当前协程。在这之前，它会有序地执行完当前协程defer列表里的语句，其它协程里挂的defer语句不作保证。因此，我们经常在defer里挂一个recover语句，防止程序直接挂掉，这起到了 try...catch的效果。

注意，recover()函数只在defer的上下文中才有效（且只有通过在defer中用匿名函数调用才有效），直接调用的话，只会返回 nil.

```GO
func main() {
    defer fmt.Println("defer main")
    var user = os.Getenv("USER_")
    go func() {
        defer func() {
            fmt.Println("defer caller")
            if err := recover(); err != nil {
                fmt.Println("recover success. err: ", err)
            }
        }()
        func() {
            defer func() {
                fmt.Println("defer here")
            }()
            if user == "" {
                panic("should set user env.")
            }
            // 此处不会执行
            fmt.Println("after panic")
        }()
    }()
    time.Sleep(100)
    fmt.Println("end of main function")
}

```

上面的panic最终会被recover捕获到。这样的处理方式在一个http server的主流程常常会被用到。一次偶然的请求可能会触发某个bug, 这时用recover捕获panic, 稳住主流程，不影响其他请求。

程序员通过监控获知此次panic的发生，按时间点定位到日志相应位置，找到发生panic的原因，三下五除二，修复上线。一看四周，大家都埋头干自己的事，简直完美：偷偷修复了一个bug，没有发现！嘿嘿！



**![img](data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==)**

#### **10、后记**



defer非常好用，一般情况下不会有什么问题。

但是只有深入理解了defer的原理才会避开它的温柔陷阱。

掌握了它的原理后，就会写出易懂易维护的代码。



> 作者：饶全成，中科院计算所硕士，滴滴出行后端研发工程师。
>
> 声明：本文为作者投稿，版权归其个人所有。
>
> 免责声明：文章广告为微信自动匹配，与本平台无关，如遇假冒伪劣请联系微信进行举报。