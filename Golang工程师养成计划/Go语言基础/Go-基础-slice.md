# Go语言基础-Slice表达式

> 资料：
>
> [The Go Programming Language Specification - Slice_expressions](https://golang.google.cn/ref/spec#Slice_expressions)
>
> [The Go Programming Language Specification - Making_slices_maps_and_channels](https://golang.google.cn/ref/spec#Making_slices_maps_and_channels)
>
> [The Go Programming Language Specification - Appending_and_copying_slices](https://golang.google.cn/ref/spec#Appending_and_copying_slices)
>
> [The Go Programming Language Specification - Slice_types](https://golang.google.cn/ref/spec#Slice_types)

### Slice types

A slice is a descriptor for a contiguous segment of an *underlying array* and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`.

```
SliceType = "[" "]" ElementType .
```

The length of a slice `s` can be discovered by the built-in function [`len`](https://golang.google.cn/ref/spec#Length_and_capacity); unlike with arrays it may change during execution. The elements can be addressed by integer [indices](https://golang.google.cn/ref/spec#Index_expressions) 0 through `len(s)-1`. The slice index of a given element may be less than the index of the same element in the underlying array.

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying a slice may extend past the end of the slice. The *capacity* is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by [*slicing*](https://golang.google.cn/ref/spec#Slice_expressions) a new one from the original slice. The capacity of a slice `a` can be discovered using the built-in function [`cap(a)`](https://golang.google.cn/ref/spec#Length_and_capacity).

A new, initialized slice value for a given element type `T` is made using the built-in function [`make`](https://golang.google.cn/ref/spec#Making_slices_maps_and_channels), which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with `make` always allocates a new, hidden array to which the returned slice value refers. That is, executing

```go
make([]T, length, capacity)
```

produces the same slice as allocating an array and [slicing](https://golang.google.cn/ref/spec#Slice_expressions) it, so these two expressions are equivalent:

```go
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

### Slice expressions

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity.

#### Simple slice expressions

For a string, array, pointer to array, or slice `a`, the primary expression

```go
a[low : high]
```

constructs a substring or slice. The *indices* `low` and `high` select which elements of operand `a` appear in the result. The result has indices starting at 0 and length equal to `high` - `low`. After slicing the array `a`

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

the slice `s` has type `[]int`, length 3, capacity 4, and elements

```go
s[0] == 2
s[1] == 3
s[2] == 4
```

For convenience, any of the indices may be omitted. A missing `low` index defaults to zero; a missing `high` index defaults to the length of the sliced operand:

```go
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

If `a` is a pointer to an array, `a[low : high]` is shorthand for `(*a)[low : high]`.

For arrays or strings, the indices are *in range* if `0` <= `low` <= `high` <= `len(a)`, otherwise they are *out of range*. For slices, the upper index bound is the slice capacity `cap(a)` rather than the length. A [constant](https://golang.google.cn/ref/spec#Constants) index must be non-negative and [representable](https://golang.google.cn/ref/spec#Representability) by a value of type `int`; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy `low <= high`. If the indices are out of range at run time, a [run-time panic](https://golang.google.cn/ref/spec#Run_time_panics) occurs.

Except for [untyped strings](https://golang.google.cn/ref/spec#Constants), if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type `string`. If the sliced operand is an array, it must be [addressable](https://golang.google.cn/ref/spec#Address_operators) and the result of the slice operation is a slice with the same element type as the array.

If the sliced operand of a valid slice expression is a `nil` slice, the result is a `nil` slice. Otherwise, if the result is a slice, it shares its underlying array with the operand.

```go
var a [10]int
s1 := a[3:7]   // underlying array of s1 is array a; &s1[2] == &a[5]
s2 := s1[1:4]  // underlying array of s2 is underlying array of s1 which is array a; &s2[1] == &a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42; they all refer to the same underlying array element
```

#### Full slice expressions

For an array, pointer to array, or slice `a` (but not a string), the primary expression

```
a[low : high : max]
```

constructs a slice of the same type, and with the same length and elements as the simple slice expression `a[low : high]`. Additionally, it controls the resulting slice's capacity by setting it to `max - low`. Only the first index may be omitted; it defaults to 0. After slicing the array `a`

```go
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

the slice `t` has type `[]int`, length 2, capacity 4, and elements

```go
t[0] == 2
t[1] == 3
```

As for simple slice expressions, if `a` is a pointer to an array, `a[low : high : max]` is shorthand for `(*a)[low : high : max]`. If the sliced operand is an array, it must be [addressable](https://golang.google.cn/ref/spec#Address_operators).

The indices are *in range* if `0 <= low <= high <= max <= cap(a)`, otherwise they are *out of range*. A [constant](https://golang.google.cn/ref/spec#Constants) index must be non-negative and [representable](https://golang.google.cn/ref/spec#Representability) by a value of type `int`; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a [run-time panic](https://golang.google.cn/ref/spec#Run_time_panics) occurs.

### Making slices, maps and channels

The built-in function `make` takes a type `T`, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions. It returns a value of type `T` (not `*T`). The memory is initialized as described in the section on [initial values](https://golang.google.cn/ref/spec#The_zero_value).

```go
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```

Each of the size arguments `n` and `m` must be of integer type or an untyped [constant](https://golang.google.cn/ref/spec#Constants). A constant size argument must be non-negative and [representable](https://golang.google.cn/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. If both `n` and `m` are provided and are constant, then `n` must be no larger than `m`. If `n` is negative or larger than `m` at run time, a [run-time panic](https://golang.google.cn/ref/spec#Run_time_panics) occurs.

```go
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
```

Calling `make` with a map type and size hint `n` will create a map with initial space to hold `n` map elements. The precise behavior is implementation-dependent.

### Appending to and copying slices

The built-in functions `append` and `copy` assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps.

The [variadic](https://golang.google.cn/ref/spec#Function_types) function `append` appends zero or more values `x` to `s` of type `S`, which must be a slice type, and returns the resulting slice, also of type `S`. The values `x` are passed to a parameter of type `...T` where `T` is the [element type](https://golang.google.cn/ref/spec#Slice_types) of `S` and the respective [parameter passing rules](https://golang.google.cn/ref/spec#Passing_arguments_to_..._parameters) apply. As a special case, `append` also accepts a first argument assignable to type `[]byte` with a second argument of string type followed by `...`. This form appends the bytes of the string.

```
append(s S, x ...T) S  // T is the element type of S
```

If the capacity of `s` is not large enough to fit the additional values, `append` allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, `append` re-uses the underlying array.

```go
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   // t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents   b == []byte{'b', 'a', 'r' }
```

The function `copy` copies slice elements from a source `src` to a destination `dst` and returns the number of elements copied. Both arguments must have [identical](https://golang.google.cn/ref/spec#Type_identity) element type `T` and must be [assignable](https://golang.google.cn/ref/spec#Assignability) to a slice of type `[]T`. The number of elements copied is the minimum of `len(src)` and `len(dst)`. As a special case, `copy` also accepts a destination argument assignable to type `[]byte` with a source argument of a string type. This form copies the bytes from the string into the byte slice.

```go
copy(dst, src []T) int
copy(dst []byte, src string) int
```

Examples:

```go
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```