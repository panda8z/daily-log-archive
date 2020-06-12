## package flag

```
import "flag"
```

flag包实现了命令行参数的解析。

要求：

使用flag.String(), Bool(), Int()等函数注册flag，下例声明了一个整数flag，解析结果保存在*int指针ip里：

```
import "flag"
var ip = flag.Int("flagname", 1234, "help message for flagname")
```

如果你喜欢，也可以将flag绑定到一个变量，使用Var系列函数：

```
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

或者你可以自定义一个用于flag的类型（满足Value接口）并将该类型用于flag解析，如下：

```
flag.Var(&flagVal, "name", "help message for flagname")
```

对这种flag，默认值就是该变量的初始值。

在所有flag都注册之后，调用：

```
flag.Parse()
```

来解析命令行参数写入注册的flag里。

解析之后，flag的值可以直接使用。如果你使用的是flag自身，它们是指针；如果你绑定到了某个变量，它们是值。

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

解析后，flag后面的参数可以从flag.Args()里获取或用flag.Arg(i)单独获取。这些参数的索引为从0到flag.NArg()-1。

命令行flag语法：

```
-flag
-flag=x
-flag x  // 只有非bool类型的flag可以
```

可以使用1个或2个'-'号，效果是一样的。最后一种格式不能用于bool类型的flag，因为如果有文件名为0、false等时,如下命令：

```
cmd -x *
```

其含义会改变。你必须使用-flag=false格式来关闭一个bool类型flag。

Flag解析在第一个非flag参数（单个"-"不是flag参数）之前停止，或者在终止符"--"之后停止。

整数flag接受1234、0664、0x1234等类型，也可以是负数。bool类型flag可以是：

```
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

时间段flag接受任何合法的可提供给time.ParseDuration的输入。

默认的命令行flag集被包水平的函数控制。FlagSet类型允许程序员定义独立的flag集，例如实现命令行界面下的子命令。FlagSet的方法和包水平的函数是非常类似的。

Example

### Index

[返回首页](https://studygolang.com/static/pkgdoc/main.html)

[Variables](https://studygolang.com/static/pkgdoc/pkg/flag.htm#pkg-variables)

[type Value](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Value)

[type Getter](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Getter)

[type ErrorHandling](https://studygolang.com/static/pkgdoc/pkg/flag.htm#ErrorHandling)

[type Flag](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Flag)

[type FlagSet](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet)

- [func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet](https://studygolang.com/static/pkgdoc/pkg/flag.htm#NewFlagSet)
- [func (f *FlagSet) Init(name string, errorHandling ErrorHandling)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Init)
- [func (f *FlagSet) NFlag() int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.NFlag)
- [func (f *FlagSet) Lookup(name string) *Flag](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Lookup)
- [func (f *FlagSet) NArg() int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.NArg)
- [func (f *FlagSet) Args() [\]string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Args)
- [func (f *FlagSet) Arg(i int) string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Arg)
- [func (f *FlagSet) PrintDefaults()](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.PrintDefaults)
- [func (f *FlagSet) SetOutput(output io.Writer)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.SetOutput)
- [func (f *FlagSet) Bool(name string, value bool, usage string) *bool](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Bool)
- [func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.BoolVar)
- [func (f *FlagSet) Int(name string, value int, usage string) *int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Int)
- [func (f *FlagSet) IntVar(p *int, name string, value int, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.IntVar)
- [func (f *FlagSet) Int64(name string, value int64, usage string) *int64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Int64)
- [func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Int64Var)
- [func (f *FlagSet) Uint(name string, value uint, usage string) *uint](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Uint)
- [func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.UintVar)
- [func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Uint64)
- [func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Uint64Var)
- [func (f *FlagSet) Float64(name string, value float64, usage string) *float64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Float64)
- [func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Float64Var)
- [func (f *FlagSet) String(name string, value string, usage string) *string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.String)
- [func (f *FlagSet) StringVar(p *string, name string, value string, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.StringVar)
- [func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Duration)
- [func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.DurationVar)
- [func (f *FlagSet) Var(value Value, name string, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Var)
- [func (f *FlagSet) Set(name, value string) error](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Set)
- [func (f *FlagSet) Parse(arguments [\]string) error](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Parse)
- [func (f *FlagSet) Parsed() bool](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Parsed)
- [func (f *FlagSet) Visit(fn func(*Flag))](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.Visit)
- [func (f *FlagSet) VisitAll(fn func(*Flag))](https://studygolang.com/static/pkgdoc/pkg/flag.htm#FlagSet.VisitAll)

[func NFlag() int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#NFlag)

[func Lookup(name string) *Flag](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Lookup)

[func NArg() int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#NArg)

[func Args() [\]string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Args)

[func Arg(i int) string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Arg)

[func PrintDefaults()](https://studygolang.com/static/pkgdoc/pkg/flag.htm#PrintDefaults)

[func Bool(name string, value bool, usage string) *bool](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Bool)

[func BoolVar(p *bool, name string, value bool, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#BoolVar)

[func Int(name string, value int, usage string) *int](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Int)

[func IntVar(p *int, name string, value int, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#IntVar)

[func Int64(name string, value int64, usage string) *int64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Int64)

[func Int64Var(p *int64, name string, value int64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Int64Var)

[func Uint(name string, value uint, usage string) *uint](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Uint)

[func UintVar(p *uint, name string, value uint, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#UintVar)

[func Uint64(name string, value uint64, usage string) *uint64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Uint64)

[func Uint64Var(p *uint64, name string, value uint64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Uint64Var)

[func Float64(name string, value float64, usage string) *float64](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Float64)

[func Float64Var(p *float64, name string, value float64, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Float64Var)

[func String(name string, value string, usage string) *string](https://studygolang.com/static/pkgdoc/pkg/flag.htm#String)

[func StringVar(p *string, name string, value string, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#StringVar)

[func Duration(name string, value time.Duration, usage string) *time.Duration](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Duration)

[func DurationVar(p *time.Duration, name string, value time.Duration, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#DurationVar)

[func Var(value Value, name string, usage string)](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Var)

[func Set(name, value string) error](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Set)

[func Parse()](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Parse)

[func Parsed() bool](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Parsed)

[func Visit(fn func(*Flag))](https://studygolang.com/static/pkgdoc/pkg/flag.htm#Visit)

[func VisitAll(fn func(*Flag))](https://studygolang.com/static/pkgdoc/pkg/flag.htm#VisitAll)

#### Examples

[返回首页](https://studygolang.com/static/pkgdoc/main.html)

[package](https://studygolang.com/static/pkgdoc/pkg/flag.htm#example-package)

### Variables

```
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

CommandLine是默认的命令行flag集，用于解析os.Args。包水平的函数如BoolVar、Arg等都是对其方法的封装。

```
var ErrHelp = errors.New("flag: help requested")
```

当flag -help被调用但没有注册这个flag时，就会返回ErrHelp。

```
var Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    PrintDefaults()
}
```

Usage打印到标准错误输出一个使用信息，记录了所有注册的flag。本函数是一个包变量，可以将其修改为指向自定义的函数。

### type [Value](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#237)

```
type Value interface {
    String() string
    Set(string) error
}
```

Value接口是用于将动态的值保存在一个flag里。（默认值被表示为一个字符串）

如果Value接口具有IsBoolFlag() bool方法，且返回真。命令行解析其会将-name等价为-name=true而不是使用下一个命令行参数。

### type [Getter](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#246)

```
type Getter interface {
    Value
    Get() interface{}
}
```

Gette接口使可以取回Value接口的内容。本接口包装了Value接口而不是作为Value接口的一部分，因为本接口是在Go 1之后出现，出于兼容的考虑才如此。本包所有的满足Value接口的类型都同时满足Getter接口。

### type [ErrorHandling](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#252)

```
type ErrorHandling int
```

ErrorHandling定义如何处理flag解析错误。

```
const (
    ContinueOnError ErrorHandling = iota
    ExitOnError
    PanicOnError
)
```

### type [Flag](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#278)

```
type Flag struct {
    Name     string // flag在命令行中的名字
    Usage    string // 帮助信息
    Value    Value  // 要设置的值
    DefValue string // 默认值（文本格式），用于使用信息
}
```

Flag类型代表一条flag的状态。

### type [FlagSet](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#262)

```
type FlagSet struct {
    // Usage函数在解析flag出现错误时会被调用
    // 该字段为一个函数（而非采用方法），以便修改为自定义的错误处理函数
    Usage func()
    // 内含隐藏或非导出字段
}
```

FlagSet代表一个已注册的flag的集合。FlagSet零值没有名字，采用ContinueOnError错误处理策略。

#### func [NewFlagSet](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#835)

```
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```

NewFlagSet创建一个新的、名为name，采用errorHandling为错误处理策略的FlagSet。

#### func (*FlagSet) [Init](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#846)

```
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

Init设置flag集合f的名字和错误处理属性。FlagSet零值没有名字，默认采用ContinueOnError错误处理策略。

#### func (*FlagSet) [NFlag](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#415)

```
func (f *FlagSet) NFlag() int
```

NFlag返回解析时进行了设置的flag的数量。

#### func (*FlagSet) [Lookup](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#343)

```
func (f *FlagSet) Lookup(name string) *Flag
```

返回已经f中已注册flag的Flag结构体指针；如果flag不存在的话，返回nil。

#### func (*FlagSet) [NArg](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#436)

```
func (f *FlagSet) NArg() int
```

NArg返回解析flag之后剩余参数的个数。

#### func (*FlagSet) [Args](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#442)

```
func (f *FlagSet) Args() []string
```

返回解析之后剩下的非flag参数。（不包括命令名）

#### func (*FlagSet) [Arg](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#422)

```
func (f *FlagSet) Arg(i int) string
```

返回解析之后剩下的第i个参数，从0开始索引。

#### func (*FlagSet) [PrintDefaults](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#377)

```
func (f *FlagSet) PrintDefaults()
```

PrintDefault打印集合中所有注册好的flag的默认值。除非另外配置，默认输出到标准错误输出中。

#### func (*FlagSet) [SetOutput](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#310)

```
func (f *FlagSet) SetOutput(output io.Writer)
```

设置使用信息和错误信息的输出流，如果output为nil，将使用os.Stderr。

#### func (*FlagSet) [Bool](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#461)

```
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

Bool用指定的名称、默认值、使用信息注册一个bool类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [BoolVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#449)

```
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar用指定的名称、默认值、使用信息注册一个bool类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Int](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#487)

```
func (f *FlagSet) Int(name string, value int, usage string) *int
```

Int用指定的名称、默认值、使用信息注册一个int类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [IntVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#475)

```
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Int64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#513)

```
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

Int64用指定的名称、默认值、使用信息注册一个int64类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [Int64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#501)

```
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var用指定的名称、默认值、使用信息注册一个int64类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Uint](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#539)

```
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

Uint用指定的名称、默认值、使用信息注册一个uint类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [UintVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#527)

```
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

UintVar用指定的名称、默认值、使用信息注册一个uint类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Uint64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#565)

```
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

Uint64用指定的名称、默认值、使用信息注册一个uint64类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [Uint64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#553)

```
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var用指定的名称、默认值、使用信息注册一个uint64类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Float64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#617)

```
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

Float64用指定的名称、默认值、使用信息注册一个float64类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [Float64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#605)

```
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var用指定的名称、默认值、使用信息注册一个float64类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [String](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#591)

```
func (f *FlagSet) String(name string, value string, usage string) *string
```

String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [StringVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#579)

```
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Duration](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#643)

```
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration用指定的名称、默认值、使用信息注册一个time.Duration类型flag。返回一个保存了该flag的值的指针。

#### func (*FlagSet) [DurationVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#631)

```
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar用指定的名称、默认值、使用信息注册一个time.Duration类型flag，并将flag的值保存到p指向的变量。

#### func (*FlagSet) [Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#661)

```
func (f *FlagSet) Var(value Value, name string, usage string)
```

Var方法使用指定的名字、使用信息注册一个flag。该flag的类型和值由第一个参数表示，该参数应实现了Value接口。例如，用户可以创建一个flag，可以用Value接口的Set方法将逗号分隔的字符串转化为字符串切片。

#### func (*FlagSet) [Set](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#354)

```
func (f *FlagSet) Set(name, value string) error
```

设置已注册的flag的值。

#### func (*FlagSet) [Parse](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#788)

```
func (f *FlagSet) Parse(arguments []string) error
```

从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。

#### func (*FlagSet) [Parsed](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#812)

```
func (f *FlagSet) Parsed() bool
```

返回是否f.Parse已经被调用过。

#### func (*FlagSet) [Visit](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#330)

```
func (f *FlagSet) Visit(fn func(*Flag))
```

按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签。

#### func (*FlagSet) [VisitAll](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#316)

```
func (f *FlagSet) VisitAll(fn func(*Flag))
```

按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数会遍历所有标签，不管解析时有无进行设置。

### func [NFlag](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#418)

```
func NFlag() int
```

NFlag返回已被设置的flag的数量。

### func [Lookup](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#349)

```
func Lookup(name string) *Flag
```

返回已经已注册flag的Flag结构体指针；如果flag不存在的话，返回nil。。

### func [NArg](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#439)

```
func NArg() int
```

NArg返回解析flag之后剩余参数的个数。

### func [Args](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#445)

```
func Args() []string
```

返回解析之后剩下的非flag参数。（不包括命令名）

### func [Arg](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#431)

```
func Arg(i int) string
```

返回解析之后剩下的第i个参数，从0开始索引。

### func [PrintDefaults](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#389)

```
func PrintDefaults()
```

PrintDefault会向标准错误输出写入所有注册好的flag的默认值。

### func [Bool](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#469)

```
func Bool(name string, value bool, usage string) *bool
```

Bool用指定的名称、默认值、使用信息注册一个bool类型flag。返回一个保存了该flag的值的指针。

### func [BoolVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#455)

```
func BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar用指定的名称、默认值、使用信息注册一个bool类型flag，并将flag的值保存到p指向的变量。

### func [Int](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#495)

```
func Int(name string, value int, usage string) *int
```

Int用指定的名称、默认值、使用信息注册一个int类型flag。返回一个保存了该flag的值的指针。

### func [IntVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#481)

```
func IntVar(p *int, name string, value int, usage string)
```

IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。

### func [Int64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#521)

```
func Int64(name string, value int64, usage string) *int64
```

Int64用指定的名称、默认值、使用信息注册一个int64类型flag。返回一个保存了该flag的值的指针。

### func [Int64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#507)

```
func Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var用指定的名称、默认值、使用信息注册一个int64类型flag，并将flag的值保存到p指向的变量。

### func [Uint](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#547)

```
func Uint(name string, value uint, usage string) *uint
```

Uint用指定的名称、默认值、使用信息注册一个uint类型flag。返回一个保存了该flag的值的指针。

### func [UintVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#533)

```
func UintVar(p *uint, name string, value uint, usage string)
```

UintVar用指定的名称、默认值、使用信息注册一个uint类型flag，并将flag的值保存到p指向的变量。

### func [Uint64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#573)

```
func Uint64(name string, value uint64, usage string) *uint64
```

Uint64用指定的名称、默认值、使用信息注册一个uint64类型flag。返回一个保存了该flag的值的指针。

### func [Uint64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#559)

```
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var用指定的名称、默认值、使用信息注册一个uint64类型flag，并将flag的值保存到p指向的变量。

### func [Float64](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#625)

```
func Float64(name string, value float64, usage string) *float64
```

Float64用指定的名称、默认值、使用信息注册一个float64类型flag。返回一个保存了该flag的值的指针。

### func [Float64Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#611)

```
func Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var用指定的名称、默认值、使用信息注册一个float64类型flag，并将flag的值保存到p指向的变量。

### func [String](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#599)

```
func String(name string, value string, usage string) *string
```

String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针。

### func [StringVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#585)

```
func StringVar(p *string, name string, value string, usage string)
```

StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。

### func [Duration](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#651)

```
func Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration用指定的名称、默认值、使用信息注册一个time.Duration类型flag。返回一个保存了该flag的值的指针。

### func [DurationVar](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#637)

```
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar用指定的名称、默认值、使用信息注册一个time.Duration类型flag，并将flag的值保存到p指向的变量。

### func [Var](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#687)

```
func Var(value Value, name string, usage string)
```

Var方法使用指定的名字、使用信息注册一个flag。该flag的类型和值由第一个参数表示，该参数应实现了Value接口。例如，用户可以创建一个flag，可以用Value接口的Set方法将逗号分隔的字符串转化为字符串切片。

### func [Set](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#371)

```
func Set(name, value string) error
```

设置已注册的flag的值。

### func [Parse](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#818)

```
func Parse()
```

从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。

### func [Parsed](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#824)

```
func Parsed() bool
```

返回是否Parse已经被调用过。

### func [Visit](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#338)

```
func Visit(fn func(*Flag))
```

按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签。

### func [VisitAll](https://github.com/golang/go/blob/master/src/flag/flag.go?name=release#324)

```
func VisitAll(fn func(*Flag))
```

按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数会遍历所有标签，不管解析时有无进行设置。

[Go语言中文网](http://studygolang.com/) | [Go Language](http://golang.org/)[Back to top](https://studygolang.com/static/pkgdoc/pkg/flag.htm#)