**垃圾回收统一定理**（Unified Theory of GC） [Bacon et al. 2004]。

内存中对象及其指针所组成了一个对象图，我们不妨令所有对象组成的集合为 VV， 设指针所指链接对象的多重集为 EE。由于我们不应该释放一个会在未来会被使用的对象， 如果我们不对任何赋值器进行分析而是进行保守地估计，则如果从栈或寄存器出发存在到达对象的路径， 则该对象将在未来被使用，记这些路径的起始对象组成的集合为根集合 RR。则对象的引用计数 ρ(v)（其中 v∈V）可以由下述的递归定点表示进行计算：

$ρ(v)=|[v:v∈R]|+|[(w,v):(w,v)∈E ∧ ρ(w)>0]|$



```go
// go 1.13.1
package main
import (
	"runtime"
	"time"
)
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	println("OK")
}
```

