[9.2 延迟语句 | Go 语言原本](https://changkun.de/golang/zh-cn/part2runtime/ch09lang/defer/)

# 9.4 Channel 与 Select

> 本节内容提供一个线上演讲：[YouTube 在线](https://www.youtube.com/watch?v=d7fFCGGn0Wc)，[Google Slides 讲稿](https://changkun.de/s/chansrc/)。

Go 语言中 Channel 与 Select 语句受到 1978 年 CSP 原始理论的启发。 在语言设计中，Goroutine 就是 CSP 理论中的并发实体， 而 Channel 则对应 CSP 中输入输出指令的消息信道，Select 语句则是 CSP 中守卫和选择指令的组合。 他们的区别在于 CSP 理论中通信是隐式的，而 Go 的通信则是显式的由程序员进行控制； CSP 理论中守卫指令只充当 Select 语句的一个分支，多个分支的 Select 语句由选择指令进行实现。

Channel 与 Select 是 Go 语言中提供的语言级的、基于消息传递的同步原语。

## Channel 的本质

### Channel 底层结构

实现 Channel 的结构并不神秘，本质上就是一个 `mutex` 锁加上一个环状缓存、 一个发送方队列和一个接收方队列：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 ` | `// src/runtime/chan.go type hchan struct { qcount   uint           // 队列中的所有数据数 dataqsiz uint           // 环形队列的大小 buf      unsafe.Pointer // 指向大小为 dataqsiz 的数组 elemsize uint16         // 元素大小 closed   uint32         // 是否关闭 elemtype *_type         // 元素类型 sendx    uint           // 发送索引 recvx    uint           // 接收索引 recvq    waitq          // recv 等待列表，即（ <-ch ） sendq    waitq          // send 等待列表，即（ ch<- ） lock mutex } type waitq struct { // 等待队列 sudog 双向队列 first *sudog last  *sudog } ` |
| ------------------------------------------------ | ------------------------------------------------------------ |
|                                                  |                                                              |

![img](Go语言原本-channel/chan.png)**图1：Channel 的结构**

其中 `recvq` 和 `sendq` 分别是 `sudog` 的一个链式队列， 其元素是一个包含当前包含队 Goroutine 及其要在 Channel 中发送的数据的一个封装， 如图 1 所示。

> 更多关于 sudog 的细节，请参考 [6.8 同步原语](https://changkun.de/golang/zh-cn/part2runtime/ch06sched/sync)。

### Channel 的创建

Channel 的创建语句由编译器完成如下翻译工作：

| `1 ` | `make(chan type, n) => makechan(type, n) ` |
| ---- | ------------------------------------------ |
|      |                                            |

将一个 `make` 语句转换为 `makechan` 调用。 而具体的 `makechan` 实现的本质是根据需要创建的元素大小， 对 `mallocgc` 进行封装， 因此，Channel 总是在堆上进行分配，它们会被垃圾回收器进行回收， 这也是为什么 Channel 不一定总是需要调用 `close(ch)` 进行显式地关闭。

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 ` | `// src/runtime/chan.go // 将 hchan 的大小对齐 const hchanSize = unsafe.Sizeof(hchan{}) + uintptr(-int(unsafe.Sizeof(hchan{}))&7) func makechan(t *chantype, size int) *hchan { elem := t.elem ... 	// 检查确认 channel 的容量不会溢出 mem, overflow := math.MulUintptr(elem.size, uintptr(size)) if overflow || mem > maxAlloc-hchanSize || size < 0 { 	panic("makechan: size out of range") } 	var c *hchan switch { case mem == 0: 	// 队列或元素大小为零 	c = (*hchan)(mallocgc(hchanSize, nil, true)) 	... case elem.ptrdata == 0: 	// 元素不包含指针 	// 在一个调用中分配 hchan 和 buf 	c = (*hchan)(mallocgc(hchanSize+mem, nil, true)) 	c.buf = add(unsafe.Pointer(c), hchanSize) default: 	// 元素包含指针 	c = new(hchan) 	c.buf = mallocgc(mem, elem, true) } 	c.elemsize = uint16(elem.size) c.elemtype = elem c.dataqsiz = uint(size) 	... return c } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

Channel 并不严格支持 `int64` 大小的缓冲，当 `make(chan type, n)` 中 n 为 `int64` 类型时， 运行时的实现仅仅只是将其强转为 `int`，提供了对 `int` 转型是否成功的检查：

| `1 2 3 4 5 6 7 8 9 ` | `// src/runtime/chan.go func makechan64(t *chantype, size int64) *hchan { if int64(int(size)) != size { 	panic("makechan: size out of range") } 	return makechan(t, int(size)) } ` |
| -------------------- | ------------------------------------------------------------ |
|                      |                                                              |

所以创建一个 Channel 最重要的操作就是创建 `hchan` 以及分配所需的 `buf` 大小的内存空间。

### 向 Channel 发送数据

发送数据完成的是如下的翻译过程：

| `1 ` | `ch <- v => chansend1(ch, v) ` |
| ---- | ------------------------------ |
|      |                                |

而本质上它会去调用更为通用的 `chansend`：

| `1 2 3 4 ` | `//go:nosplit func chansend1(c *hchan, elem unsafe.Pointer) { chansend(c, elem, true) } ` |
| ---------- | ------------------------------------------------------------ |
|            |                                                              |

下面我们来关注 `chansend` 的具体实现的第一个部分：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 ` | `func chansend(c *hchan, ep unsafe.Pointer, block bool) bool { // 当向 nil channel 发送数据时，会调用 gopark // 而 gopark 会将当前的 Goroutine 休眠，从而发生死锁崩溃 if c == nil { 	if !block { 		return false 	} 	gopark(nil, nil, waitReasonChanSendNilChan) 	throw("unreachable") } 	... } ` |
| --------------------------------- | ------------------------------------------------------------ |
|                                   |                                                              |

在这个部分中，我们可以看到，如果一个 Channel 为零值（比如没有初始化），这时候的发送操作会暂止当前的 Goroutine（`gopark`）。 而 gopark 会将当前的 Goroutine 休眠，从而发生死锁崩溃。

现在我们来看一切已经准备就绪，开始对 Channel 加锁：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 ` | `func chansend(c *hchan, ep unsafe.Pointer, block bool) bool { ... lock(&c.lock) 	// 持有锁之前我们已经检查了锁的状态， // 但这个状态可能在持有锁之前、该检查之后发生变化， // 因此还需要再检查一次 channel 的状态 if c.closed != 0 { // 不允许向已经 close 的 channel 发送数据 	unlock(&c.lock) 	panic(plainError("send on closed channel")) } 	// 1. channel 上有阻塞的接收方，直接发送 if sg := c.recvq.dequeue(); sg != nil { 	send(c, sg, ep, func() { unlock(&c.lock) }) 	return true } 	// 2. 判断 channel 中缓存是否有剩余空间 if c.qcount < c.dataqsiz { 	// 有剩余空间，存入 c.buf 	qp := chanbuf(c, c.sendx) 	... 	typedmemmove(c.elemtype, qp, ep) // 将要发送的数据拷贝到 buf 中 	c.sendx++ 	if c.sendx == c.dataqsiz { // 如果 sendx 索引越界则设为 0 		c.sendx = 0 	} 	c.qcount++ // 完成存入，记录增加的数据，解锁 	unlock(&c.lock) 	return true } if !block { 	unlock(&c.lock) 	return false } 	... } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

到目前位置，代码中考虑了当 Channel 上有接收方等待，可以直接将数据发送走，并返回（情况 1）；或没有接收方 但缓存中还有剩余空间来存放没有读取的数据（情况 2）。对于直接发送数据的情况，由 `send` 调用完成：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 ` | `func send(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func()) { ... if sg.elem != nil { 	sendDirect(c.elemtype, sg, ep) 	sg.elem = nil } gp := sg.g unlockf() // unlock(&c.lock) gp.param = unsafe.Pointer(sg) ... // 复始一个 Goroutine，放入调度队列等待被后续调度 goready(gp) // 将 gp 作为下一个立即被执行的 Goroutine } func sendDirect(t *_type, sg *sudog, src unsafe.Pointer) { dst := sg.elem ... // 为了确保发送的数据能够被立刻观察到，需要写屏障支持，执行写屏障，保证代码正确性 memmove(dst, src, t.size) // 直接写入接收方的执行栈！ } ` |
| ------------------------------------------------ | ------------------------------------------------------------ |
|                                                  |                                                              |

`send` 操作其实是隐含了有接收方阻塞在 Channel 上，换句话说有接收方已经被暂止， 当我们发送完数据后，应该让该接收方就绪（让调度器继续开始调度接收方）。

这个 `send` 操作其实是一种优化。原因在于，已经处于等待状态的 Goroutine 是没有被执行的， 因此用户态代码不会与当前所发生数据发生任何竞争。我们也更没有必要冗余的将数据写入到缓存， 再让接收方从缓存中进行读取。因此我们可以看到， `sendDirect` 的调用， 本质上是将数据直接写入接收方的执行栈。

最后我们来看第三种情况，如果既找不到接收方，`buf` 也已经存满， 这时我们就应该阻塞当前的 Goroutine 了：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 ` | `func chansend(c *hchan, ep unsafe.Pointer, block bool) bool { ... 	// 3. 阻塞在 channel 上，等待接收方接收数据 gp := getg() mysg := acquireSudog() ... c.sendq.enqueue(mysg) gopark(chanparkcommit, unsafe.Pointer(&c.lock)) // 将当前的 g 从调度队列移出 	// 因为调度器在停止当前 g 的时候会记录运行现场，当恢复阻塞的发送操作时候，会从此处继续开始执行 ... gp.waiting = nil gp.activeStackChans = false if gp.param == nil { 	if c.closed == 0 { // 正常唤醒状态，Goroutine 应该包含需要传递的参数，但如果没有唤醒时的参数，且 channel 没有被关闭，则为虚假唤醒 		throw("chansend: spurious wakeup") 	} 	panic(plainError("send on closed channel")) } gp.param = nil ... mysg.c = nil // 取消与之前阻塞的 channel 的关联 releaseSudog(mysg) // 从 sudog 中移除 return true } func chanparkcommit(gp *g, chanLock unsafe.Pointer) bool { // 具有未解锁的指向 gp 栈的 sudog。栈的复制必须锁住那些 sudog 的 channel gp.activeStackChans = true unlock((*mutex)(chanLock)) return true } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

简单总结一下，发送过程包含三个步骤：

1. 持有锁
2. 入队，拷贝要发送的数据
3. 释放锁

其中第二个步骤包含三个子步骤：

1. 找到是否有正在阻塞的接收方，是则直接发送
2. 找到是否有空余的缓存，是则存入
3. 阻塞直到被唤醒

### 从 Channel 接收数据

接收数据主要是完成以下翻译工作：

| `1 2 ` | `v <- ch      =>  chanrecv1(ch, v) v, ok <- ch  =>  ok := chanrecv2(ch, v) ` |
| ------ | ------------------------------------------------------------ |
|        |                                                              |

他们的本质都是调用 `chanrecv`：

| `1 2 3 4 5 6 7 8 9 ` | `//go:nosplit func chanrecv1(c *hchan, elem unsafe.Pointer) { chanrecv(c, elem, true) } //go:nosplit func chanrecv2(c *hchan, elem unsafe.Pointer) (received bool) { _, received = chanrecv(c, elem, true) return } ` |
| -------------------- | ------------------------------------------------------------ |
|                      |                                                              |

chanrecv 的具体实现如下，由于我们已经仔细分析过发送过程了， 我们不再详细分拆下面代码的步骤，其处理方式基本一致：

1. 上锁
2. 从缓存中出队，拷贝要接收的数据
3. 解锁

其中第二个步骤包含三个子步骤：

1. 如果 Channel 已被关闭，且 Channel 没有数据，立刻返回
2. 如果存在正在阻塞的发送方，说明缓存已满，从缓存队头取一个数据，再复始一个阻塞的发送方
3. 否则，检查缓存，如果缓存中仍有数据，则从缓存中读取，读取过程会将队列中的数据拷贝一份到接收方的执行栈中
4. 没有能接受的数据，阻塞当前的接收方 Goroutine

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 ` | `func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) { ... // nil channel，同 send，会导致两个 Goroutine 的死锁 if c == nil { 	if !block { 		return 	} 	gopark(nil, nil, waitReasonChanReceiveNilChan) 	throw("unreachable") } 	// 快速路径: 在不需要锁的情况下检查失败的非阻塞操作 // // 注意到 channel 不能由已关闭转换为未关闭，则 // 失败的条件是：1. 无 buf 时发送队列为空 2. 有 buf 时，buf 为空 // 此处的 c.closed 必须在条件判断之后进行验证， // 因为指令重排后，如果先判断 c.closed，得出 channel 未关闭，无法判断失败条件中 // channel 是已关闭还是未关闭（从而需要 atomic 操作） if !block && (c.dataqsiz == 0 && c.sendq.first == nil || 	c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) && 	atomic.Load(&c.closed) == 0 { 	return } 	... 	lock(&c.lock) 	// 1. channel 已经 close，且 channel 中没有数据，则直接返回 if c.closed != 0 && c.qcount == 0 { 	... 	unlock(&c.lock) 	if ep != nil { 		typedmemclr(c.elemtype, ep) 	} 	return true, false } 	// 2. channel 上有阻塞的发送方，直接接收 if sg := c.sendq.dequeue(); sg != nil { 	recv(c, sg, ep, func() { unlock(&c.lock) }) 	return true, true } 	// 3. channel 的 buf 不空 if c.qcount > 0 { 	// 直接从队列中接收 	qp := chanbuf(c, c.recvx) 	... 	if ep != nil { 		typedmemmove(c.elemtype, ep, qp) 	} 	typedmemclr(c.elemtype, qp) 	c.recvx++ 	if c.recvx == c.dataqsiz { 		c.recvx = 0 	} 	c.qcount-- 	unlock(&c.lock) 	return true, true } 	if !block { 	unlock(&c.lock) 	return false, false } 	// 4. 没有数据可以接收，阻塞当前 Goroutine gp := getg() mysg := acquireSudog() ... c.recvq.enqueue(mysg) gopark(chanparkcommit, unsafe.Pointer(&c.lock), waitReasonChanReceive) 	... // 被唤醒 gp.waiting = nil ... closed := gp.param == nil gp.param = nil mysg.c = nil releaseSudog(mysg) return true, !closed } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

接收数据同样包含直接往接收方的执行栈中拷贝要发送的数据，但这种情况当且仅当缓存大小为0时（即无缓冲 Channel）。

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 ` | `func recv(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) { if c.dataqsiz == 0 { 	... 	if ep != nil { 		// 直接从对方的栈进行拷贝 		recvDirect(c.elemtype, sg, ep) 	} } else { 	// 从缓存队列拷贝 	qp := chanbuf(c, c.recvx) 	... 	// 从队列拷贝数据到接收方 	if ep != nil { 		typedmemmove(c.elemtype, ep, qp) 	} 	// 从发送方拷贝数据到队列 	typedmemmove(c.elemtype, qp, sg.elem) 	c.recvx++ 	if c.recvx == c.dataqsiz { 		c.recvx = 0 	} 	c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz } sg.elem = nil gp := sg.g unlockf() gp.param = unsafe.Pointer(sg) ... goready(gp, skip+1) } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

到目前为止我们终于明白了为什么无缓冲 Channel 而言 `v <- ch` happens before `ch <- v` 了， 因为**无缓冲 Channel 的接收方会先从发送方栈拷贝数据后，发送方才会被放回调度队列中，等待重新调度**。

### Channel 的关闭

关闭 Channel 主要是完成以下翻译工作：

| `1 ` | `close(ch) => closechan(ch) ` |
| ---- | ----------------------------- |
|      |                               |

具体的实现中，首先对 Channel 上锁，而后依次将阻塞在 Channel 的 g 添加到一个 gList 中，当所有的 g 均从 Channel 上移除时，可释放锁，并唤醒 gList 中的所有接收方和发送方：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 ` | `func closechan(c *hchan) { if c == nil { // close 一个空的 channel 会 panic 	panic(plainError("close of nil channel")) } 	lock(&c.lock) if c.closed != 0 { // close 一个已经关闭的的 channel 会 panic 	unlock(&c.lock) 	panic(plainError("close of closed channel")) } 	... c.closed = 1 	var glist gList 	// 释放所有的接收方 for { 	sg := c.recvq.dequeue() 	if sg == nil { // 队列已空 		break 	} 	if sg.elem != nil { 		typedmemclr(c.elemtype, sg.elem) // 清零 		sg.elem = nil 	} 	... 	gp := sg.g 	gp.param = nil 	... 	glist.push(gp) } 	// 释放所有的发送方 for { 	sg := c.sendq.dequeue() 	if sg == nil { // 队列已空 		break 	} 	sg.elem = nil 	... 	gp := sg.g 	gp.param = nil 	... 	glist.push(gp) } // 释放 channel 的锁 unlock(&c.lock)     // 就绪所有的 G for !glist.empty() { 	gp := glist.pop() 	gp.schedlink = 0 	goready(gp, 3) } } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

当 Channel 关闭时，我们必须让所有阻塞的接收方重新被调度，让所有的发送方也重新被调度，这时候 的实现先将 Goroutine 统一添加到一个列表中（需要锁），然后逐个地进行复始（不需要锁）。

## Select 语句的本质

### 分支的随机化

Select 本身会被编译为 `selectgo` 调用。这与普通的多个 if 分支不同。 `selectgo` 则用于随机化每条分支的执行顺序，普通多个 if 分支的执行顺序始终是一致的。

| `  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246 247 248 249 250 251 252 253 254 255 256 257 258 259 260 261 262 263 264 265 ` | `type scase struct { c           *hchan         // chan elem        unsafe.Pointer // 数据元素 kind        uint16 ... } func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool) { ... 	cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0)) order1 := (*[1 << 17]uint16)(unsafe.Pointer(order0)) 	scases := cas1[:ncases:ncases] pollorder := order1[:ncases:ncases] lockorder := order1[ncases:][:ncases:ncases] 	// 替换零值的 channel for i := range scases { 	cas := &scases[i] 	if cas.c == nil && cas.kind != caseDefault { 		*cas = scase{} 	} } 	... 	// 生成随机顺序 for i := 1; i < ncases; i++ { 	j := fastrandn(uint32(i + 1)) 	pollorder[i] = pollorder[j] 	pollorder[j] = uint16(i) } 	// 根据 channel 的地址进行堆排序，决定加锁的顺序，避免死锁 for i := 0; i < ncases; i++ { 	... } ... 	// 依次加锁 sellock(scases, lockorder) 	var ( 	gp     *g 	sg     *sudog 	c      *hchan 	k      *scase 	sglist *sudog 	sgnext *sudog 	qp     unsafe.Pointer 	nextp  **sudog ) loop: // 1 遍历 channel，检查是否就绪（可发送/可接收） var dfli int var dfl *scase var casi int var cas *scase var recvOK bool for i := 0; i < ncases; i++ { 	casi = int(pollorder[i]) 	cas = &scases[casi] 	c = cas.c 	switch cas.kind { 	case caseNil: 		continue 	case caseRecv: 		sg = c.sendq.dequeue() 		if sg != nil { 			goto recv 		} 		if c.qcount > 0 { 			goto bufrecv 		} 		if c.closed != 0 { 			goto rclose 		} 	case caseSend: 		... 		if c.closed != 0 { 			goto sclose 		} 		sg = c.recvq.dequeue() 		if sg != nil { 			goto send 		} 		if c.qcount < c.dataqsiz { 			goto bufsend 		} 	case caseDefault: 		dfli = casi 		dfl = cas 	} } // 存在 default 分支，直接去 retc 执行 if dfl != nil { 	selunlock(scases, lockorder) 	casi = dfli 	cas = dfl 	goto retc } 	// 2 入队所有的 channel gp = getg() ... nextp = &gp.waiting for _, casei := range lockorder { 	casi = int(casei) 	cas = &scases[casi] 	if cas.kind == caseNil { 		continue 	} 	c = cas.c 	sg := acquireSudog() 	sg.g = gp 	sg.isSelect = true 	// 在 gp.waiting 上分配 elem 和入队 sg 之间没有栈分段，copystack 可以在其中找到它。 	sg.elem = cas.elem 	... 	sg.c = c 	// 按锁的顺序创建等待链表 	*nextp = sg 	nextp = &sg.waitlink 		switch cas.kind { 	case caseRecv: 		c.recvq.enqueue(sg) 		case caseSend: 		c.sendq.enqueue(sg) 	} } 	// 等待被唤醒 gp.param = nil // selparkcommit 根据等待列表依次解锁 gopark(selparkcommit, nil, waitReasonSelect) 	// 重新上锁 sellock(scases, lockorder) 	gp.selectDone = 0 sg = (*sudog)(gp.param) gp.param = nil 	// pass 3 - 从不成功的 channel 中出队 // 否则将它们堆到一个安静的 channel 上并记录所有成功的分支 // 我们按锁的顺序单向链接 sudog casi = -1 cas = nil sglist = gp.waiting // 从 gp.waiting 取消链接之前清除所有的 elem for sg1 := gp.waiting; sg1 != nil; sg1 = sg1.waitlink { 	sg1.isSelect = false 	sg1.elem = nil 	sg1.c = nil } gp.waiting = nil 	for _, casei := range lockorder { 	k = &scases[casei] 	if k.kind == caseNil { 		continue 	} 	... 	if sg == sglist { 		// sg 已经被唤醒我们的 G 出队了。 		casi = int(casei) 		cas = k 	} else { 		c = k.c 		if k.kind == caseSend { 			c.sendq.dequeueSudoG(sglist) 		} else { 			c.recvq.dequeueSudoG(sglist) 		} 	} 	sgnext = sglist.waitlink 	sglist.waitlink = nil 	releaseSudog(sglist) 	sglist = sgnext } 	if cas == nil { 	// 当一个参与在 select 语句中的 channel 被关闭时，我们可以在 gp.param == nil 时进行唤醒(所以 cas == nil) 	// 最简单的方法就是循环并重新运行该操作，然后就能看到它现在已经被关闭了 	// 也许未来我们可以显式的发送关闭信号， 	// 但我们就必须区分在接收方上关闭和在发送方上关闭这两种情况了 	// 最简单的方法是不复制代码并重新检查上面的代码。 	// 我们知道某些 channel 被关闭了，也知道某些可能永远不会被重新打开，因此我们不会再次阻塞 	goto loop } 	c = cas.c ... if cas.kind == caseRecv { 	recvOK = true } ... selunlock(scases, lockorder) goto retc bufrecv: // 可以从 buf 接收 ... recvOK = true qp = chanbuf(c, c.recvx) if cas.elem != nil { 	typedmemmove(c.elemtype, cas.elem, qp) } typedmemclr(c.elemtype, qp) c.recvx++ if c.recvx == c.dataqsiz { 	c.recvx = 0 } c.qcount-- selunlock(scases, lockorder) goto retc bufsend: // 可以发送到 buf ... typedmemmove(c.elemtype, chanbuf(c, c.sendx), cas.elem) c.sendx++ if c.sendx == c.dataqsiz { 	c.sendx = 0 } c.qcount++ selunlock(scases, lockorder) goto retc recv: // 可以从一个休眠的发送方 (sg)直接接收 recv(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2) ... recvOK = true goto retc rclose: // 在已经关闭的 channel 末尾进行读 selunlock(scases, lockorder) recvOK = false if cas.elem != nil { 	typedmemclr(c.elemtype, cas.elem) } ... goto retc send: // 可以向一个休眠的接收方 (sg) 发送 ... send(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2) ... goto retc retc: ... return casi, recvOK sclose: // 向已关闭的 channel 进行发送 selunlock(scases, lockorder) panic(plainError("send on closed channel")) } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

### 发送数据的分支

Select 的诸多用法其实本质上仍然是 Channel 操作，编译器会完成如下翻译工作：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 ` | `select { case c <- v: ... default: ... } => if selectnbsend(c, v) { ... } else { ... } ` |
| ------------------------------------ | ------------------------------------------------------------ |
|                                      |                                                              |

其中：

| `1 2 3 ` | `func selectnbsend(c *hchan, elem unsafe.Pointer) (selected bool) { return chansend(c, elem, false, getcallerpc()) } ` |
| -------- | ------------------------------------------------------------ |
|          |                                                              |

注意，这时 `chansend` 的第三个参数为 `false`，这与前面的普通 Channel 发送操作不同， 说明这时 Select 的操作是非阻塞的。

我们现在来关注 `chansend` 中当 block 为 `false` 的情况：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 ` | `func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool { 	... 	// 快速路径: 检查不需要加锁时失败的非阻塞操作 if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) || 	(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) { 	return false } 	... 	lock(&c.lock) 	... } ` |
| ------------------------------------------ | ------------------------------------------------------------ |
|                                            |                                                              |

这里的快速路径是一个优化，它发生在持有 Channel 锁之前。 这一连串检查不需要加锁有以下原因：

1. Channel 没有被关闭与 Channel 是否满的检查没有因果关系。换句话说，无论 Channel 是否被关闭，都不能得出 Channel 是否已满；Channel 是否满，也与 Channel 是否关闭无关，从而当发生指令重排时，这个检查也不会出错。
2. 当 Channel 已经被关闭、且缓存已满时，发送操作一定失败。

第二个关于 Select 的处理则是在当判断完 Channel 是否有 `buf` 可缓存当前的数据后， 如果没有读者阻塞在 Channel 上则会立即返回失败：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 ` | `func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool { 	... 	lock(&c.lock) 	... 	// 2. 判断 channel 中缓存是否仍然有空间剩余 if c.qcount < c.dataqsiz { 	// 有空间剩余，存入 buffer 	... 	unlock(&c.lock) 	return true } if !block { 	unlock(&c.lock) 	return false } 	... } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

因此这也是为什么，我们在没有配合 for 循环使用 Select 时，需要对发送失败进行处理，例如：

| ` 1 2 3 4 5 6 7 8 9 10 11 ` | `func main() { ch := make(chan interface{}) x := 1 select { case ch <- x: 	println("send success") // 如果初始化为有缓存 channel，则会发送成功 default: 	println("send failed") // 此时 send failed 会被输出 } return } ` |
| --------------------------- | ------------------------------------------------------------ |
|                             |                                                              |

如果读者进一步尝试没有 default 的例子：

| ` 1 2 3 4 5 6 7 8 9 10 11 ` | `// main.go package main func main() { ch := make(chan interface{}) x := 1 select { case ch <- x: 	println("send success") // 如果初始化为有缓存 channel，则会发送成功 } return } ` |
| --------------------------- | ------------------------------------------------------------ |
|                             |                                                              |

会发现，此时程序会发生 panic：

| `1 2 3 4 5 ` | `$ go run main.go fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]: main.main() ` |
| ------------ | ------------------------------------------------------------ |
|              |                                                              |

似乎与源码中发生的行为并不一致，因为按照之前的分析，当锁被解除后，并不会出现任何 panic。 这是为什么呢？事实上，编译器会特殊处理 **当 Select 语句只有一个分支的情况，即 `select` 关键字在只有一个分支时，没有被翻译成 `selectgo`。** 只有一个分支的情况下，`select` 与 `if` 是没有区别的，这种优化消除了只有一个分支情况下调用 `selectgo` 的性能开销：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 ` | `// src/cmd/compile/internal/gc/select.go func walkselectcases(cases *Nodes) []*Node { // 获取 case 分支的数量 n := cases.Len() 	// 优化: 没有 case 的情况 if n == 0 { 	// 翻译为：block() 	... 	return } 	// 优化: 只有一个 case 的情况 if n == 1 { 	// 翻译为：if ch == nil { block() }; n; 	... 	return } 	// 一般情况，调用 selecggo ... } ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

根据编译器的代码，我们甚至可以看到没有分支的 Select 会被编译成 `block` 的调用：

| `1 2 3 ` | `func block() { gopark(nil, nil, waitReasonSelectNoCases) // forever } ` |
| -------- | ------------------------------------------------------------ |
|          |                                                              |

即让整个 Goroutine 暂止。

### 接收数据的分支

对于接收数据而言，编译器会将这段语法：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 ` | `select { case v = <-c: ... default: ... } => if selectnbrecv(&v, c) { ... } else { ... } ` |
| ------------------------------------ | ------------------------------------------------------------ |
|                                      |                                                              |

而

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 ` | `select { case v, ok = <-c: ... foo default: ... bar } => if c != nil && selectnbrecv2(&v, &ok, c) { ... foo } else { ... bar } ` |
| ------------------------------------ | ------------------------------------------------------------ |
|                                      |                                                              |

其中：

| `1 2 3 4 5 6 7 8 ` | `func selectnbrecv(elem unsafe.Pointer, c *hchan) (selected bool) { selected, _ = chanrecv(c, elem, false) return } func selectnbrecv2(elem unsafe.Pointer, received *bool, c *hchan) (selected bool) { selected, *received = chanrecv(c, elem, false) return } ` |
| ------------------ | ------------------------------------------------------------ |
|                    |                                                              |

## Channel 的无锁实现

早在 2014 年时，Dmitry Vyukov 就已经提出实现无锁版本的 Channel [Vyukov, 2014a] [Vyukov, 2014b]， 但这提案虽然早年已经实现，但至今未被接受，其未被接收这一现实可以总结为以下三个原因。

早年的 Channel 实现基于比较交换的重试机制，换句话说：多个阻塞在同一 Channel 的 Goroutine 被唤醒时， 需要重新持有锁，这时谁抢到锁谁就能拿到数据。所以这些 Goroutine 被唤醒的顺序不是 FIFO，而是随机的， 最坏情况下可能存在一个 Goroutine 始终不会接受到数据。

后来 Russ Cox 希望 [Cox, 2015] 阻塞的 Goroutine 能够按照 FIFO 的顺序被唤醒 （虽然在语言层面上未定义多个 Goroutine 的唤醒顺序），保证得到数据的公平性，参与讨论的人中也表示支持这一提案。 但这一决定基本上抹杀了无锁 Channel 的实现机制 [Randall, 2015a]。 这是目前未使用无锁实现 Channel 的一个最主要的原因。

那在这个决定之前，无锁 Channel 早就已经实现了，为什么当时没有接受使用无锁版本的 Channel 呢？

第一个原因是提出的无锁 Channel 并非无等待算法，是否能有效提高 Channel 在大规模应用的性能并没有大规模测试的强有力的证据， 支撑性能表现的只有 Dmitry Vyukov 提交的性能测试； 与此同时，运行时调度器不是 NUMA-aware 的实现，在 CPU 核心与调度器 P 数量较多时， 一个社区实现的无锁 Channel [OneOfOne, 2016] 的性能测试结果 [Gjengset, 2016] 表明： 无锁版本的 Channel 甚至比基于 futex 加锁版本的 Channel 还要慢。 在后续对 Channel 性能优化的跟进中虽然没有采用无锁实现， 但仍然跟进了两个小成本的优化 [Vyukov, 2014d]：增加不需要锁时的快速路径和减少互斥锁的粒度。

第二个原因导致没有被接受的原因则在于：无锁版本的 Channel 可维护性大打折扣。 这里我们简单提一个由于无锁实现导致的维护性大打折扣的教训 [Randall, 2015b]。 在早年简化 Channel 实现的过程中，由于没有考虑到发送数据过程中， 对要发送数据的指针进行读取，将会与调度器对执行栈的伸缩发生竞争。这是因为 直接读取 Channel 的数据分为两个过程：1. 读取发送方的值的指针 2. 拷贝到要接收的位置。 然而在 1 和 2 这两个步骤之间，发送方的执行栈可能发生收缩，进而指针失效，成为竞争的源头。

虽然后来有人提出使用无锁编程的形式化验证工具 spin [Bell Labs, 1980] 来让调度器代码与形式验证的模型进行同步，但显然这需要更多的工作量，并没有人采取任何行动。

## 小结

Channel 的实现是一个典型的环形队列加上 `mutex` 锁的实现， 与 Channel 同步出现的 Select 更像是一个语法糖， 其本质仍然是一个 `chansend` 和 `chanrecv` 的两个通用实现。 但为了支持 Select 在不同分支上的非阻塞操作，`selectgo` 完成了这一需求。

考虑到整个 Channel 操作带锁的成本较高，官方也曾考虑过使用无锁 Channel 的设计， 但由于年代久远，该改进仍处于搁置状态 [Vyukov, 2014b]。

## 进一步阅读的参考文献

- [Vyukov, 2014a] [Dmitry Vyukov, Go channels on steroids, January 2014](https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub)
- [Vyukov, 2014b] [Dmitry Vyukov, runtime: lock-free channels, October 2014](https://github.com/golang/go/issues/8899)
- [Vyukov, 2014c] [Dmitry Vyukov, runtime: chans on steroids, October 2014](https://codereview.appspot.com/12544043)
- [Vyukov, 2014d] [update on “lock-free channels”, 2015](https://groups.google.com/forum/#!msg/golang-dev/0IElw_BbTrk/cGHMdNoHGQEJ)
- [Cox, 2015] [runtime: make sure blocked channels run operations in FIFO order](https://github.com/golang/go/issues/11506)
- [Randall, 2015a] [Keith Randall, runtime: simplify buffered channels, 2015](https://go-review.googlesource.com/c/go/+/9345/)
- [Randall, 2015b] [Keith Randall, runtime: simplify chan ops, take 2, 2015](https://go-review.googlesource.com/c/go/+/16740)
- [OneOfOne, 2016] [OneOfOne, A scalable lock-free channel, 2016](https://github.com/OneOfOne/lfchan)
- [Gjengset, 2016] [Jon Gjengset, Fix poor scalability to many (true-SMP) cores, 2016](https://github.com/OneOfOne/lfchan/issues/3)
- [Chenebault, 2017] [Benjamin Chenebault, runtime: select is not fair](https://github.com/golang/go/issues/21806)
- [Bell Labs, 1980] [Bell Labs, Verifying Multi-threaded Software with Spin, 1980](http://spinroot.com/spin/whatispin.html)



# 11.5 基准测试

> 本节内容提供一个线上演讲：[Google Slides 讲稿](https://changkun.de/s/gobench/)

TODO: 内容需要丰富描述

在《Software Testing: Printciples and Practices》一书中归纳的性能测试方法论：

- 搜集需求
- 编写测试用例
- 自动化性能测试用例
- 执行性能测试用例
- 分析性能测试结果
- 性能调优
- 性能基准测试（Performance Benchmarking）
- 向客户推荐合适的配置

## 11.5.1 可靠的测试环境

影响测试环境的软硬件因素

- 硬件：CPU 型号、温度、IO 等
- 软件：操作系统版本、当前系统调度的负载等

指导思想

- 单次测量结果毫无意义，统计意义下可对比的结果是关键
  - 分析测试的场景、多次测量、决定统计检验的类型
- 可对比的结果是在可控的环境下得到的
  - 笔记本电脑 CPU 的执行效率受电源管理等因素影响，连续测试同一段代码可能先得到短暂的性能提升，而后由于温度的上升导致性能下降
  - 虚拟机或（共享）云服务器上可能受到宿主机资源分配等因素导致测量结果不稳定

性能基准测试的两个基本目标：

- 可重复性：在其他外在条件不变的情况下，性能度量结果是稳定、可重复的（能复现的才叫 Bug）
- 可比较性：总是存在一个可以比较的基本线（有比较才有伤害）

## 11.5.2 benchstat 及其原理

benchstat 的功能非常简单，作用只是对性能测试结果进行统计分析，对测量结果进行假设检验，从而消除结果的观测误差（observational error）。

| ` 1 2 3 4 5 6 7 8 9 10 11 12 ` | `$ go get golang.org/x/perf/cmd/benchstat $ benchstat --help usage: benchstat [options] old.txt [new.txt] [more.txt ...] options:  -alpha α         设置显著性水平 α 的值（默认 0.05）  -delta-test test        设置显著性检验的类型，支持 utest/ttest/none（默认 utest）  -geomean        输出几何平均值  -sort order        对结果进行排序: [-]delta, [-]name, none (默认值 none) ` |
| ------------------------------ | ------------------------------------------------------------ |
|                                |                                                              |

当对一个性能基准测试 B 结果反复执行 n 次后，就能得到 b1, …, bn 个不同的结果； 在优化代码后，还能得到另外 m 个不同的结果 b1’, …, bm’。

一个稳定的基准测试，结果倾向于在某个值附近波动，于是通过通用的计算 1.5 倍四分位距法则（1.5 x InterQuartile Range Rule）的方法来消除异常值。

benchstat 的本质就是在消除异常值之后的两组浮点数之间进行假设检验（Hypothesis Testing）。

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 ` | `type Metrics struct {   Unit    string    // 性能测试的名称   Values  []float64 // 某个性能测试的度量值   RValues []float64 // 移除的异常值   Min     float64   // RValues 的最小值   Mean    float64   // RValues 的平均值   Max     float64   // RValues 的最大值 } func (m *Metrics) computeStats() {   values := stats.Sample{Xs: m.Values}   q1, q3 := values.Percentile(0.25), values.Percentile(0.75)   lo, hi := q1-1.5*(q3-q1), q3+1.5*(q3-q1) // 计算结果的四分位距，并移除异常值   for _, value := range m.Values {       if lo <= value && value <= hi { m.RValues = append(m.RValues, value) }   }   // 求统计量   m.Min, m.Max = stats.Bounds(m.RValues)   m.Mean = stats.Mean(m.RValues) } ... // 在 benchstat.Collection.Tables() 中 pval, _ := deltaTest(old, new) // 进行假设检验 ` |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

注意：

1. Q2 中位数，Q1 是最小值与 Q2 之间的中位数，Q3 是最大值和 Q2 之间的中位数
2. 异常值不是一个数学概念，因此四分位距法则只是一个「法则」，不是一个数学定理
3. 如果数据本身就是正态分布，那么等价于 μ ± 2.68 σ（因为 Q1 = μ + 0.67σ，Q3 = μ - 0.67σ，IQR = 1.34σ ⇒ Q3+1.5IQR = μ + 2.68 σ）

我们之后再来讨论假设检验。

## 11.5.3 性能基准测试示例

### 例 1：对 `sync.Map.Delete` 的一个优化

在 sync.Map 中存储一个值，然后再并发的删除该值：

| `1 2 3 4 5 6 7 8 ` | `func BenchmarkDeleteCollision(b *testing.B) {    benchMap(b, bench{        setup: func(_ *testing.B, m mapInterface) { m.LoadOrStore(0, 0) },        perG: func(b *testing.B, pb *testing.PB, i int, m mapInterface) {            for ; pb.Next(); i++ { m.Delete(0) }        },    }) } ` |
| ------------------ | ------------------------------------------------------------ |
|                    |                                                              |

在 sync.Map 中：

| `1 2 3 ` | `275 -delete(m.dirty, key) 275 +e, ok = m.dirty[key] 276 +m.missLocked() ` |
| -------- | ------------------------------------------------------------ |
|          |                                                              |

能够得到结果：

| `1 2 3 4 5 6 7 8 9 ` | `$ git stash $ go test -run=none -bench=BenchmarkDeleteCollision -count=20 | tee old.txt $ git stash pop $ go test -run=none -bench=BenchmarkDeleteCollision -count=20 | tee new.txt $ benchstat old.txt new.txt name                                      old time/op  new time/op  delta DeleteCollision/*sync_test.DeepCopyMap-8   104ns ± 0%   103ns ± 1%     ~     (p=0.383 n=20+20) DeleteCollision/*sync_test.RWMutexMap-8   67.6ns ± 2%  68.2ns ± 2%   +0.89%  (p=0.009 n=20+20) DeleteCollision/*sync.Map-8               94.2ns ± 2%   5.7ns ± 2%  -93.98%  (p=0.000 n=20+19) ` |
| -------------------- | ------------------------------------------------------------ |
|                      |                                                              |

- 可以观察到一个值被作为异常值消除了
- git stash 并不总是适用，它具有一定的局限性。

### 例 2：测试代码错误

创建一颗红黑树，并依次将 0 … n 插入到这颗红黑树中：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 ` | `func BenchmarkRBTree_PutWrong(b *testing.B) {    for size := 0; size < 1000; size += 100 {        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {            tree := ds.NewRBTree(func(a, b interface{}) bool {                if a.(int) < b.(int) { return true }                return false            })            for i := 0; i < b.N; i++ {                for n := 0; n < size; n++ { tree.Put(n, n) }            }        })    } } ` |
| --------------------------------- | ------------------------------------------------------------ |
|                                   |                                                              |

| ` 1 2 3 4 5 6 7 8 9 10 11 ` | `name                    time/op RBTree_PutWrong/size-0-8    0.65ns ± 0% RBTree_PutWrong/size-100-8  14.2µs ± 3% RBTree_PutWrong/size-200-8  30.5µs ± 0% RBTree_PutWrong/size-300-8  47.0µs ± 0% RBTree_PutWrong/size-400-8  63.3µs ± 0% RBTree_PutWrong/size-500-8  79.6µs ± 0% RBTree_PutWrong/size-600-8  97.3µs ± 0% RBTree_PutWrong/size-700-8   113µs ± 0% RBTree_PutWrong/size-800-8   131µs ± 0% RBTree_PutWrong/size-900-8   146µs ± 0% ` |
| --------------------------- | ------------------------------------------------------------ |
|                             |                                                              |

![img](Go语言原本-channel/benchstat-rbtree-wrong.png)**图 1: 红黑树插入性能（错误形式）**

为什么插入的性能是线性的？红黑树的插入性能不是 O(log(n)) 吗？ 代码写错了……吧……？

红黑树的插入性能是指当树的大小为 n-1 插入第 n 个值时的性能：

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 ` | `func BenchmarkRBTree_Put(b *testing.B) {   for size := 0; size < 1000; size += 100 {       b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {           tree := ds.NewRBTree(func(a, b interface{}) bool {               if a.(int) < b.(int) { return true }               return false           }) +          for n := 0; n < size-1; n++ { tree.Put(n, n) } +          b.ResetTimer()           for i := 0; i < b.N; i++ {               for n := 0; n < size; n++ { tree.Put(n, n) }           }       })   } } ` |
| --------------------------------------- | ------------------------------------------------------------ |
|                                         |                                                              |

```
name                    time/op
RBTree_Put/size-0-8    55.2ns ± 0%
RBTree_Put/size-100-8   158ns ± 0%
RBTree_Put/size-200-8   172ns ± 0%
RBTree_Put/size-300-8   181ns ± 0%
RBTree_Put/size-400-8   202ns ± 0%
RBTree_Put/size-500-8   203ns ± 0%
RBTree_Put/size-600-8   209ns ± 0%
RBTree_Put/size-700-8   209ns ± 0%
RBTree_Put/size-800-8   216ns ± 0%
RBTree_Put/size-900-8   217ns ± 0%
```

![img](Go语言原本-channel/benchstat-rbtree-correct.png)**图 2: 红黑树插入性能（正确形式）**

### 例 3：编译器优化

编译器优化产生的直接影响是测量的目标不准确，这一点在 C++ 编译器中相当严重。编译器优化是一个比较大的话题，我们不在此进行深入讨论，后续章节会进一步讨论此主题。只举比较常见的一例：

TODO:

## 11.5.4 假设检验的原理

### 统计学的基本概念

- 总体：所有满足某些共同性质的值的集合（共同性质：接口）
- 样本：从总体中随机抽取的个体
- 频率：n 次试验中，某个事件发生的次数除以总的试验次数
- 大数定理：当试验次数 n → ∞ 时，频率一定收敛到某个值
- 概率：频率收敛到的值，性质之一：0≤P(A)≤10≤P(A)≤1
- 独立：两个事件互不影响，性质之一：P(AB)=P(A)P(B)P(AB)=P(A)P(B)
- 随机变量：是一个函数，参数是所有可能的样本，返回值是这些样本的取值，例如 P(X=2)=0.25P(X=2)=0.25
- 期望：随机变量以其概率为权重的加权平均值，即 E(X)=∑ixipiE(X)=∑ixipi
- 方差：样本取值与期望之间的「距离」，距离定义为差的平方和，即 Var(X)=∑i(xi−E(X))2Var(X)=∑i(xi−E(X))2
- 概率密度函数：是一个函数，参数是随机变量取值，返回值是随机变量取得该值的概率
- 累积分布函数：随机变量取值小于某个值的概率
- 正态分布：一种特殊的概率密度函数 N(μ,σ2)N(μ,σ2)
- 中心极限定理：无穷多个独立的随机变量的和服从正态分布

考虑读者水平可能参差不齐，所以最基础的开始回顾。 这里的定义并不是概率论中的公理化定义。所以本书重述了概念本身，目的只是为了方便理解，严谨性有较大欠缺。

比如，概率是通过非负性、规范性、可列可加性定义的一个样本空间上的可测函数，需要解释严格的样本空间的定义、可测函数是什么，进而又牵扯出为什么需要可测，不可测又有什么问题等一些很严肃的、实变函数等数学专业课才会讨论的数学概念，进而严重偏离了主题本身，所以选择了使用大数定理所阐明的概率是频率的渐进值来定义概率；

再比如，随机变量概念本身其实很好理解，但严格的数学定义是概率空间到实数集的一个可测函数，为了保留随机变量的本质是一个函数，所以选择了从样本到取值的解释方式，等等其他概念也存在不同程度上的重述 … 但总体上，符合概念理解循序渐进的原则。

最后，中心极限定理存在多种形式，最为通用的（广义上的，没有同分布假设的版本）是 Lindeberg-Feller 的版本，也叫 Lindeberg-Feller 中心极限定理，中心极限定理成立的充要条件是 Lindeberg 条件。

### 假设检验的基本框架

统计是一套在总体分布函数完全未知或者只知道形式、不知参数的情况下，为了由样本推断总体的某些未知特性，形成的一套方法论。

多次抽样：对同一个性能基准测试运行多次，根据中心极限定理，如果理论均值存在，则抽样噪声服从正态分布的。

当重复执行完某个性能基准测试后， benchstat 先帮我们剔除掉了一些异常值，我们得到了关于某段代码在可控的环境条件 E 下的性能分布的一组样本。

现在的问题是：

- 非参数方法：剩下样本是否来自同一总体？总体是什么分布？两组样本在可控的测试环境下进行吗？
- 参数方法：如果总体分布已经确定，那么样本的变化是否显著？性能的基准测试前后，是否具有统计意义下的明显变化？

假设检验：利用样本判断对总体的假设是否成立的过程

零假设 H0：想要驳回的论点

备择假设 H1：拒绝零假设后的备用项，我们想要证明的论点

p 值：零假设发生的概率

显著性水平：可靠程度

例如：在性能基准测试中，

- H0：代码修改前后，性能没有提升
- H1：代码修改前后，性能有显著提升
- p < 0.05：H0 发生的概率小于 5%，在至少 95% 的把握下，性能有显著提升

| 真实情况   | 接受零假设   | 拒绝零假设   |
| ---------- | ------------ | ------------ |
| 零假设为真 | 正确         | 犯第一类错误 |
| 零假设为假 | 犯第二类错误 | 正确         |

第一类错误：把对的判断成错的；第二类错误：把错的判断成对的

当样本不变时，减少犯某类错误的概率择会增加犯另一类错误的概率。控制第一类错误的概率，让它小于某个 p 值（0.05）称之为显著性检验

- 零假设 H0: 代码性能测试的均值没有显著变化 μ0−μ1=0μ0−μ1=0

- 备择假设 H1: 代码性能有显著变化

   

  

  μ0−μ1≠0μ0−μ1≠0

  - 对性能提升持有保守态度，尽可能避免出现实际没有提升，但被判断为提升（第一类错误）
  - 在我们的场景下，应该拒绝零假设

### Welch T 检验和 Mann-Whitney U 检验

在 benchstat 的性能测试中提供了 Welch T 检验和 Man-Whitney U 检验，他们对数据的假设不同。

两个总体均值差的检验：H0:μ1−μ2=0,H1:μ1−μ2≠0H0:μ1−μ2=0,H1:μ1−μ2≠0

**T 检验**

参数检验，假设数据服从正态分布，且方差相同

**Welch T 检验**

参数检验，假设服从正态分布，方差一定不相同

**Mann-Whitney U 检验**

非参数检验，假设最少，最通用，只假设两组样本来自同一总体，只有均值上的差异（保守派）

当对数据的假设减少时，结论的不确定性就会增大，因此 p 值会相应的变大，进而使性能基准测试的条件更加严格。

## 11.5.5 性能基准测试的局限性

### 系统噪音

`perflock` 作用是限制 CPU 时钟频率，从而一定程度上消除系统对性能测试程序的影响，减少结果的噪声，进而性能测量的结果方差更小也更加可靠，仅支持 Linux。

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 ` | `$ go get github.com/aclements/perflock/cmd/perflock $ sudo install $GOPATH/bin/perflock /usr/bin/perflock $ sudo -b perflock -daemon $ perflock Usage of perflock:  perflock [flags] command...  perflock -list  perflock -daemon  -daemon        启动 perflock 守护进程  -governor percent        设置运行指令所占用的 CPU 频率比例，或 none 没有调整（默认 90%）  -list        列出当前正在等待执行的命令  -shared        在共享模式下获取锁 (默认独占模式) $ perflock -governor 70% go test -test=none -bench=.  ` |
| ------------------------------------------------ | ------------------------------------------------------------ |
|                                                  |                                                              |

perflock 的原理在于，在执行命令前，通知 perflock 守护进程，守护进程将 cpufreq 进行备份，调整到 perflock-client 指定的频率，再通知 perflock-client 开始执行 Benchmark：

![img](Go语言原本-channel/bench-perflock.png)**图 3: perflock 的原理**

进行如下修改：

```
/sys/devices/system/cpu/cpu0/cpufreq/scaling_max_freq
==
/sys/devices/system/cpu/cpu0/cpufreq/scaling_min_freq
==
target := (max-min)*percent/100 + min
```

一些需要注意的问题：

1. 不能在 Parallels 上使用（虚拟机都不行？未在其他虚拟机测试），原因在于无法获取 /sys/devices/system/cpu/cpu\d+$/cpufreq
2. 粗略看代码当不能获取 cpu domains 时，客户端会崩溃，但是实际上并不会
3. 不要在执行性能测试时强制 kill perflock daemon，否则 cpufreq 参数将不会恢复
4. 只锁定了系统的 CPU 频率，并没有限制与系统中其他资源的占用情况，该被打断的依然会被打断

还有一些其他的方法：

- 禁用地址空间随机化： `echo 0 > /proc/sys/kernel/randomize_va_space`
- 禁用 Intel Turbo 模式：`echo 1 > /sys/devices/system/cpu/intel_pstate/no_turbo`
- 禁用 CPU SMT pair：`echo 0 > /sys/devices/system/cpu/cpu*/online ⇐ 在 /sys/devices/system/cpu/cpu*/topology/thread_siblings_list`
- 使用 `cpuset：cset shield -c N1,N2 -k on` → 将所有线程移出 N1, N2，-k on 表示内核线程也会被移除
- 报告 perf：`cset shield --exec -- perf stat -r 10 ` → `--` 之后的命令将在隔离的 CPU 上运行。`perf` 将运行 `cmd` 10 次

### 多重比较谬误

广泛比较两个不同群体的所有差异，从中找出具有差异的特征，宣称是造成两个群体不同的原因。

在不对代码进行优化的情况下，反复对不同的性能测试结果样本进行显著性检验，直到找到能够使 p 值能够满足显著性水平，宣称性能得到了提升。

![img](Go语言原本-channel/bench-significant.png)**图 4: 多重比较谬误，图片来源：https://xkcd.com/882/**

### 机器过热

对结果进行回归，肉眼可见的性能下降：

| `1 2 3 ` | `#if _FP_W_TYPE_SIZE < 32 #error "Here's a nickel kid. Go buy yourself a real computer." #endif ` |
| -------- | ------------------------------------------------------------ |
|          |                                                              |

买台好电脑吧。

![img](Go语言原本-channel/cgo-go-c.png)**图 5: 性能基准测试随时间推移而导致变差（越小越好）**

## 11.5.6 总结

进行（严肃的）性能测试前的检查清单：

- 限制系统资源，降低测试噪声：perflock
  - 限制 CPU 时钟频率：perflock
  - （如果需要）限制 runtime 消耗的内存上限: runtime.SetMaxHeap
  - 关闭无关程序和进程等等……
- 确定测试代码的正确性
  - 考虑 Goroutine 的终止性，当某些并发的工作发生在基准测试结束后，那么测量是不准确的
  - 考虑编译器进行了过度优化或基准测试代码本身编写错误导致测量程序不正确
- 实施性能基准测试
  - （如果需要）计算需要采样的次数
  - 使用 git stash 记录并撤销代码的修改，执行测试得到修改前的性能测试结果
  - 使用 git stash pop 恢复代码的修改内容，执行测试得到修改后的性能测试结果
  - 使用 benchstat 对前后测量到的性能测量进行假设检验
  - 验证结果有效性，例如确认结果的波动，比较随时间推移造成的性能回归等等

## 进一步阅读的参考文献

- https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html

  - 这是一篇很早之前的关于Go程序性能测试的文章，里面讲述了相当多有关性能调优、测试的主题，不仅局限于这次分享的主题

- https://github.com/golang/go/issues/27400

  - 这是一个未解决的 Issue，目的是希望 Go 团队能够在 testing 包中使用文档来说明编译器优化的情况，进而避免基准测试测量不准确的问题

- https://github.com/golang/go/issues/23471

  - 这是一个未解决的 Issue，目的是希望 Go 团队能够发布一篇官方文档来详述如何科学的对 Go 程序进行性能测试 当然，本次分享的 PPT 其实解决了这个问题 :)

- A Review and Comparison of Methods for Detecting Outliers in Univariate Data Sets,

   

  http://d-scholarship.pitt.edu/7948/1/Seo.pdf

  - 这篇论文比较了统计学中的一些异常值检测的方法

- Mann, Henry B., and Donald R. Whitney. “On a test of whether one of two random variables is stochastically larger than the other.” The annals of mathematical statistics (1947): 50-60.

  - 这是 Mann-Whitney U 检验的原始论文

- Mytkowicz, Todd, et al. “Producing wrong data without doing anything obviously wrong!.” ACM Sigplan Notices 44.3 (2009): 265-276.

  - 这篇文章介绍了适用因果分析和随机化的方法来检测并避免测量误差