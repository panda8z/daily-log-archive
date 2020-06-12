## package sort

```
import "sort"
```

sort包提供了排序切片和用户自定义数据集的函数。

### Index

[返回首页](https://studygolang.com/static/pkgdoc/main.html)

[type Interface](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Interface)

[type IntSlice](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice)

- [func (p IntSlice) Len() int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice.Len)
- [func (p IntSlice) Less(i, j int) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice.Less)
- [func (p IntSlice) Search(x int) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice.Search)
- [func (p IntSlice) Sort()](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice.Sort)
- [func (p IntSlice) Swap(i, j int)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntSlice.Swap)

[type Float64Slice](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice)

- [func (p Float64Slice) Len() int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice.Len)
- [func (p Float64Slice) Less(i, j int) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice.Less)
- [func (p Float64Slice) Search(x float64) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice.Search)
- [func (p Float64Slice) Sort()](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice.Sort)
- [func (p Float64Slice) Swap(i, j int)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64Slice.Swap)

[type StringSlice](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice)

- [func (p StringSlice) Len() int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice.Len)
- [func (p StringSlice) Less(i, j int) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice.Less)
- [func (p StringSlice) Search(x string) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice.Search)
- [func (p StringSlice) Sort()](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice.Sort)
- [func (p StringSlice) Swap(i, j int)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringSlice.Swap)

[func Ints(a [\]int)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Ints)

[func IntsAreSorted(a [\]int) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IntsAreSorted)

[func SearchInts(a [\]int, x int) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#SearchInts)

[func Float64s(a [\]float64)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64s)

[func Float64sAreSorted(a [\]float64) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Float64sAreSorted)

[func SearchFloat64s(a [\]float64, x float64) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#SearchFloat64s)

[func Strings(a [\]string)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Strings)

[func StringsAreSorted(a [\]string) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#StringsAreSorted)

[func SearchStrings(a [\]string, x string) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#SearchStrings)

[func Sort(data Interface)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Sort)

[func Stable(data Interface)](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Stable)

[func Reverse(data Interface) Interface](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Reverse)

[func IsSorted(data Interface) bool](https://studygolang.com/static/pkgdoc/pkg/sort.htm#IsSorted)

[func Search(n int, f func(int) bool) int](https://studygolang.com/static/pkgdoc/pkg/sort.htm#Search)



### type [Interface](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#12)

```
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
```

一个满足sort.Interface接口的（集合）类型可以被本包的函数进行排序。方法要求集合中的元素可以被整数索引。

### type [IntSlice](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#233)

```
type IntSlice []int
```

IntSlice给[]int添加方法以满足Interface接口，以便排序为递增序列。

#### func (IntSlice) [Len](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#235)

```
func (p IntSlice) Len() int
```

#### func (IntSlice) [Less](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#236)

```
func (p IntSlice) Less(i, j int) bool
```

#### func (IntSlice) [Swap](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#237)

```
func (p IntSlice) Swap(i, j int)
```

#### func (IntSlice) [Sort](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#240)

```
func (p IntSlice) Sort()
```

Sort等价于调用Sort(p)

#### func (IntSlice) [Search](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#106)

```
func (p IntSlice) Search(x int) int
```

Search等价于调用SearchInts(p, x)

### type [Float64Slice](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#243)

```
type Float64Slice []float64
```

Float64Slice给[]float64添加方法以满足Interface接口，以便排序为递增序列。

#### func (Float64Slice) [Len](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#245)

```
func (p Float64Slice) Len() int
```

#### func (Float64Slice) [Less](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#246)

```
func (p Float64Slice) Less(i, j int) bool
```

#### func (Float64Slice) [Swap](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#247)

```
func (p Float64Slice) Swap(i, j int)
```

#### func (Float64Slice) [Sort](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#255)

```
func (p Float64Slice) Sort()
```

Sort等价于调用Sort(p)

#### func (Float64Slice) [Search](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#109)

```
func (p Float64Slice) Search(x float64) int
```

Search等价于调用SearchFloat64s(p, x)

### type [StringSlice](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#258)

```
type StringSlice []string
```

StringSlice给[]string添加方法以满足Interface接口，以便排序为递增序列。

#### func (StringSlice) [Len](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#260)

```
func (p StringSlice) Len() int
```

#### func (StringSlice) [Less](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#261)

```
func (p StringSlice) Less(i, j int) bool
```

#### func (StringSlice) [Swap](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#262)

```
func (p StringSlice) Swap(i, j int)
```

#### func (StringSlice) [Sort](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#265)

```
func (p StringSlice) Sort()
```

Sort等价于调用Sort(p)

#### func (StringSlice) [Search](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#112)

```
func (p StringSlice) Search(x string) int
```

Search等价于调用SearchStrings(p, x)

### func [Sort](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#192)

```
func Sort(data Interface)
```

Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。

### func [Stable](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#317)

```
func Stable(data Interface)
```

Stable排序data，并保证排序的稳定性，相等元素的相对次序不变。

它调用1次data.Len，O(n*log(n))次data.Less和O(n*log(n)*log(n))次data.Swap。

### func [IsSorted](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#220)

```
func IsSorted(data Interface) bool
```

IsSorted报告data是否已经被排序。

### func [Reverse](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#215)

```
func Reverse(data Interface) Interface
```

Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。

Example

### func [Search](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#59)

```
func Search(n int, f func(int) bool) int
```

Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。也就是说，Search函数希望f在输入位于区间[0, n)的前面某部分（可以为空）时返回假，而在输入位于剩余至结尾的部分（可以为空）时返回真；Search函数会返回满足f(i)==true的最小值i。如果没有该值，函数会返回n。注意，未找到时的返回值不是-1，这一点和strings.Index等函数不同。Search函数只会用区间[0, n)内的值调用f。

一般使用Search找到值x在插入一个有序的、可索引的数据结构时，应插入的位置。这种情况下，参数f（通常是闭包）会捕捉应搜索的值和被查询的数据集。

例如，给定一个递增顺序的切片，调用Search(len(data), func(i int) bool { return data[i] >= 23 })会返回data中最小的索引i满足data[i] >= 23。如果调用者想要知道23是否在切片里，它必须另外检查data[i] == 23。

搜索递减顺序的数据时，应使用<=运算符代替>=运算符。

下列代码尝试在一个递增顺序的整数切片中找到值x：

```
x := 23
i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
if i < len(data) && data[i] == x {
	// x is present at data[i]
} else {
	// x is not present in data,
	// but i is the index where it would be inserted.
}
```

一个更古怪的例子，下面的程序会猜测你持有的数字：

```
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
```

### func [Ints](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#270)

```
func Ints(a []int)
```

Ints函数将a排序为递增顺序。

Example

### func [IntsAreSorted](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#279)

```
func IntsAreSorted(a []int) bool
```

IntsAreSorted检查a是否已排序为递增顺序。

### func [SearchInts](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#83)

```
func SearchInts(a []int, x int) int
```

SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。

### func [Float64s](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#273)

```
func Float64s(a []float64)
```

Float64s函数将a排序为递增顺序。

### func [Float64sAreSorted](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#282)

```
func Float64sAreSorted(a []float64) bool
```

Float64sAreSorted检查a是否已排序为递增顺序。

### func [SearchFloat64s](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#92)

```
func SearchFloat64s(a []float64, x float64) int
```

SearchFloat64s在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。

### func [Strings](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#276)

```
func Strings(a []string)
```

Strings函数将a排序为递增顺序。

### func [StringsAreSorted](https://github.com/golang/go/blob/master/src/sort/sort.go?name=release#285)

```
func StringsAreSorted(a []string) bool
```

StringsAreSorted检查a是否已排序为递增顺序。

### func [SearchStrings](https://github.com/golang/go/blob/master/src/sort/search.go?name=release#101)

```
func SearchStrings(a []string, x string) int
```

SearchStrings在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。

[Go语言中文网](http://studygolang.com/) | [Go Language](http://golang.org/)[Back to top](https://studygolang.com/static/pkgdoc/pkg/sort.htm#)