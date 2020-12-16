---
title: Kotlin语言探索003-BasicTypes
date: 2018-11-10 13:18:44
updated:
comments: true
tags: 
- Android 
- 2018 
- 技术
- Kotlin
categories:	
- Kotlin
permalink:
---

# Basic Types 基础类型

In Kotlin, everything is an object in the sense that we can call member functions and properties on any variable.
Some of the types can have a special internal representation - for example, numbers, characters and booleans can be
represented as primitive values at runtime - but to the user they look like ordinary classes. 
In this section we describe the basic types used in Kotlin: numbers, characters, booleans, arrays, and strings.

 在`Kotli`中所有的东西都是对象,不论是成员函数还是各种变量属性都是如此.
有一些类型可以用一种特定的内部对象标示,比如,数值型,字符型,布尔型等,它们可以在运行时使用原始值,但用户看起来他们和普通的类没有区别.
在这一章节,我们描述一些`Kotlin`里面使用的基本类型: `numbers`(数值型),`characters`(字符型), `booleans`(布尔型), `arrarys`(数组)和 `string`(字符串)

<!-- more -->

## Numbers 数值型

Kotlin handles numbers in a way close to Java, but not exactly the same. For example, there are no implicit widening conversions for numbers, and literals are slightly different in some cases.

Kotlin provides the following built-in types representing numbers (this is close to Java):

| Type   | Bit width |
| ------ | --------- |
| Double | 64        |
| Float  | 32        |
| Long   | 64        |
| Int    | 32        |
| Short  | 16        |
| Byte   | 8         |

Note that characters are not numbers in Kotlin.

在Kotlin中数值型的表示方式和Java类似,但也有不同之处. 例如,这里的`numbers`没有隐式的扩大转换,在某些情况下数值型字面值略有不同。

Kotlin provides the following built-in types representing numbers (this is close to Java):
Kotlin提供了一下的内置类型来表示数值型(这些几乎和java相同):

| 类型   | 位宽 |
| ------ | ---- |
| Double | 64   |
| Float  | 32   |
| Long   | 64   |
| Int    | 32   |
| Short  | 16   |
| Byte   | 8    |

请注意,请注意字符型不是数值型.

### Literal Constants 字面值常量

There are the following kinds of literal constants for integral values:

* Decimals: `123`
  * Longs are tagged by a capital `L`: `123L`
* Hexadecimals: `0x0F`
* Binaries: `0b00001011`

NOTE: Octal literals are not supported.

Kotlin also supports a conventional notation for floating-point numbers:
 
* Doubles by default: `123.5`, `123.5e10`
* Floats are tagged by `f` or `F`: `123.5f`
  
一下是数值类型的几种整数值常量:

* 十进制: `123`
  * 长整型由大写 `L`做后缀表示: `123L`
* 十六进制: `0x0F`
* 二进制: `0b00001011`

注意: 八进制字面值在Kotlin中不支持.

Kotlin还支持浮点数的传统表示法:

* 双精度浮点型: `123.5`, `123.5e10`
* 浮点型由后缀 `f` 或 `F`表示: `123.5f`

### Underscores in numeric literals (since 1.1) 在数值字面量中使用下划线(since 1.1)

You can use underscores to make number constants more readable:
你可以在数值字面量中使用下划线使更容易阅读:

```kotlin
val oneMillion = 1_000_000
val creditCardNumber = 1234_5678_9012_3456L
val socialSecurityNumber = 999_99_9999L
val hexBytes = 0xFF_EC_DE_5E
val bytes = 0b11010010_01101001_10010100_10010010
```

### Representation 表示法

On the Java platform, numbers are physically stored as JVM primitive types, unless we need a nullable number reference (e.g. `Int?`) or generics are involved. 
In the latter cases numbers are boxed.

Note that boxing of numbers does not necessarily preserve identity:

在Java平台,数值类型是被JVM的原始类型物理表示出来的,现在我们的数值类型中有了非空类型和泛型,后者数值型是被包装起来的.

请注意，被包装起来的数值不一定保存了相同的标识：

```kotlin
fun main(args: Array<String>) {
    val a: Int = 10000
    println(a === a) // Prints 'true'
    val boxedA: Int? = a
    val anotherBoxedA: Int? = a
    println(boxedA === anotherBoxedA) // !!!Prints 'false'!!!
}
```

On the other hand, it preserves equality:

然而,他们确保存了相同的值.

```kotlin
fun main(args: Array<String>) {
    val a: Int = 10000
    println(a == a) // Prints 'true'
    val boxedA: Int? = a
    val anotherBoxedA: Int? = a
    println(boxedA == anotherBoxedA) // Prints 'true'
}
```

### Explicit Conversions 显式转换

Due to different representations, smaller types are not subtypes of bigger ones.
If they were, we would have troubles of the following sort:

由于不同的标识方法,更小的类型并不是更大类型的子类型.如果是的话,就会遇到一下类型错误:

```kotlin
// Hypothetical code, does not actually compile:
val a: Int? = 1 // A boxed Int (java.lang.Integer)
val b: Long? = a // implicit conversion yields a boxed Long (java.lang.Long)
print(b == a) // Surprise! This prints "false" as Long's equals() checks whether the other is Long as well
```

So equality would have been lost silently all over the place, not to mention identity.
As a consequence, smaller types are NOT implicitly converted to bigger types.
This means that we cannot assign a value of type `Byte` to an `Int` variable without an explicit conversion

所以值相等就会在整个地方默默地消失，更不用说唯一标识了。
因此，较小的类型不会隐式转换为更大的类型。
这意味着在没有明确的显式转换条件下不能把一个`Byte`类型的值赋值给`Int`类型的变量.

```kotlin
fun main(args: Array<String>) {

    val b: Byte = 1 // OK, literals are checked statically
    val i: Int = b // ERROR

}
```

We can use explicit conversions to widen numbers

我们可以用显式转换来扩展数值类型的大小.

```kotlin
fun main(args: Array<String>) {
    val b: Byte = 1

    val i: Int = b.toInt() // OK: explicitly widened
    print(i)

}
```

Every number type supports the following conversions:
每一个数值类型都支持一下的转换方法:

* `toByte(): Byte`
* `toShort(): Short`
* `toInt(): Int`
* `toLong(): Long`
* `toFloat(): Float`
* `toDouble(): Double`
* `toChar(): Char`

Absence of implicit conversions is rarely noticeable because the type is inferred from the context, and arithmetical operations are overloaded for appropriate conversions, for example

缺少隐式转换很少会引起注意，因为类型是从上下文中推断出来的，并且算术运算会因适当的转换而过载，例如

```kotlin
val l = 1L + 3 // Long + Int => Long
```

### Operations 运算符

Kotlin supports the standard set of arithmetical operations over numbers, which are declared as members of appropriate classes (but the compiler optimizes the calls down to the corresponding instructions).
See [Operator overloading](operator-overloading.html).

As of bitwise operations, there're no special characters for them, but just named functions that can be called in infix form, for example:

Kotlin支持标准的数值型算术运算,这些已经在数值类型的相关勒种声明(但是编译器会优化调用的相关指令.),
对于位运算，它们没有特殊操作符，只是可以以中缀形式调用的命名函数，例如：

```kotlin
val x = (1 shl 2) and 0x000FF000
```

Here is the complete list of bitwise operations (available for `Int` and `Long` only):

* `shl(bits)` – signed shift left (Java's `<<`)
* `shr(bits)` – signed shift right (Java's `>>`)
* `ushr(bits)` – unsigned shift right (Java's `>>>`)
* `and(bits)` – bitwise and
* `or(bits)` – bitwise or
* `xor(bits)` – bitwise xor
* `inv()` – bitwise inversion

---

Kotlin支持标准的数值型算术运算,这些已经在数值类型的相关勒种声明(但是编译器会优化调用的相关指令.),
对于位运算，它们没有特殊操作符，只是可以以中缀形式调用的命名函数，例如：

```kotlin
val x = (1 shl 2) and 0x000FF000
```

位运算操作符列表 (只有 `Int` 和 `Long` 可以使用位运算.):

* `shl(bits)` – 有符号左移 (Java's `<<`)
* `shr(bits)` – 有符号右移 (Java's `>>`)
* `ushr(bits)` – 无符号右移 (Java's `>>>`)
* `and(bits)` – 按位与
* `or(bits)` – 按位或
* `xor(bits)` – 按位与非
* `inv()` – 按位翻转

### Floating Point Numbers Comparison 浮点数比较

The operations on floating point numbers discussed in this section are:

* Equality checks: `a == b` and `a != b`
* Comparison operators: `a < b`, `a > b`, `a <= b`, `a >= b`
* Range instantiation and range checks: `a..b`, `x in a..b`, `x !in a..b`

When the operands `a` and `b` are statically known to be `Float` or `Double` or their nullable counterparts (the type is 
declared or inferred or is a result of a [smart cast](typecasts.html#smart-casts)), the operations on the 
numbers and the range that they form follow the IEEE 754 Standard for Floating-Point Arithmetic. 

However, to support generic use cases and provide total ordering, when the operands are **not** statically typed as 
floating point numbers (e.g. `Any`, `Comparable<...>`, a type parameter), the operations use the 
`equals` and `compareTo` implementations for `Float` and `Double`, which disagree with the standard, so that:

* `NaN` is considered equal to itself
* `NaN` is considered greater than any other element including `POSITIVE_INFINITY`
* `-0.0` is considered less than `0.0`

---
以下是这一章节提到的浮点数运算符

* 相等运算符: `a == b` and `a != b`
* 比较运算符: `a < b`, `a > b`, `a <= b`, `a >= b`
* 范围实例和范围检查: `a..b`, `x in a..b`, `x !in a..b`

当显式明确操作数a和b是Float或Double或它们的nullable值（声明或推断类型或者是智能转换的结果）时，
对数值型和它们形成的范围的操作遵循`IEEE 754`浮点运算标准。 
但是，为了支持通用用例并提供总排序，
当操作数不是静态类型为浮点数（例如Any，Comparable <...>，类型参数）时，
不同意标准操作使用Float和Double 的`equals`和`compareTo`实现 ,
根据一下原则:

* `NaN` 和自己相等
* `NaN` 比其他任意可表示整数更大 `POSITIVE_INFINITY`
* `-0.0` 比 `0.0` 小

## Characters 字符型

Characters are represented by the type `Char`. They can not be treated directly as numbers

```kotlin
fun check(c: Char) {
    if (c == 1) { // ERROR: incompatible types
        // ...
    }
}
```

Character literals go in single quotes: `'1'`.
Special characters can be escaped using a backslash.
The following escape sequences are supported: `\t`, `\b`, `\n`, `\r`, `\'`, `\"`, `\\` and `\$`.
To encode any other character, use the Unicode escape sequence syntax: `'\uFF00'`.

We can explicitly convert a character to an `Int` number:

```kotlin
fun decimalDigitValue(c: Char): Int {
    if (c !in '0'..'9')
        throw IllegalArgumentException("Out of range")
    return c.toInt() - '0'.toInt() // Explicit conversions to numbers
}
```

Like numbers, characters are boxed when a nullable reference is needed. Identity is not preserved by the boxing operation.

---

字符型使用`Char`表示,它可以直接使用数值型表示.

```kotlin
fun check(c: Char) {
    if (c == 1) { // ERROR: incompatible types
        // ...
    }
}
```

字符型字面量使用单引号包裹: `'1'`;
特殊字符要使用转义字符.
这些是一部分支持的转义字符: `\t`, `\b`, `\n`, `\r`, `\'`, `\"`, `\\` 和 `\$`.
可以使用Unicode方式编码任意字符: `'\uFF00'`.

我们可以将一个字符型的值显式转换为`Int`类型:

```kotlin
fun decimalDigitValue(c: Char): Int {
    if (c !in '0'..'9')
        throw IllegalArgumentException("Out of range")
    return c.toInt() - '0'.toInt() // Explicit conversions to numbers
}
```

像数值型一样字符型在被当做nullable引用时被装箱,被包装后的字符型可能不会存储相同的标识.但会保存同样的值.

## Booleans 布尔型

The type `Boolean` represents booleans, and has two values: *true*{: .keyword } and *false*{: .keyword }.

Booleans are boxed if a nullable reference is needed.

Built-in operations on booleans include

* `||` – lazy disjunction
* `&&` – lazy conjunction
* `!` - negation

---
布尔型使用 `Boolean` 表示, 只有两个值: *true* 和 *false*
布尔型会在需要nullable的时候被包装起来.

内置的布尔值运算符包括:

* `||` – 短路或
* `&&` – 短路与
* `!` - 非

## Arrays

Arrays in Kotlin are represented by the `Array` class, that has `get` and `set` functions (that turn into `[]` by operator overloading conventions), and `size` property, along with a few other useful member functions:
数组在Kotlin里使用`Array`类表示,它有`get`,`set`方法和`size`属性,另外还有一下的几个方法:

```kotlin
class Array<T> private constructor() {
    val size: Int
    operator fun get(index: Int): T
    operator fun set(index: Int, value: T): Unit

    operator fun iterator(): Iterator<T>
    // ...
}
```

To create an array, we can use a library function `arrayOf()` and pass the item values to it, so that `arrayOf(1, 2, 3)` creates an array `[1, 2, 3]`.
Alternatively, the `arrayOfNulls()` library function can be used to create an array of a given size filled with null elements.

Another option is to use the `Array` constructor that takes the array size and the function that can return the initial value
of each array element given its index:

```kotlin
fun main(args: Array<String>) {

    // Creates an Array<String> with values ["0", "1", "4", "9", "16"]
    val asc = Array(5, { i -> (i * i).toString() })
    asc.forEach { println(it) }

}
```

As we said above, the `[]` operation stands for calls to member functions `get()` and `set()`.

Note: unlike Java, arrays in Kotlin are invariant. This means that Kotlin does not let us assign an `Array<String>`
to an `Array<Any>`, which prevents a possible runtime failure (but you can use `Array<out Any>`, 
see [Type Projections](generics.html#type-projections)).

Kotlin also has specialized classes to represent arrays of primitive types without boxing overhead: `ByteArray`,
`ShortArray`, `IntArray` and so on. These classes have no inheritance relation to the `Array` class, but they
have the same set of methods and properties. Each of them also has a corresponding factory function:

```kotlin
val x: IntArray = intArrayOf(1, 2, 3)
x[0] = x[1] + x[2]
```

---
数组在Kotlin里使用`Array`类表示,它有`get`,`set`方法和`size`属性,另外还有一下的几个方法:

```kotlin
class Array<T> private constructor() {
    val size: Int
    operator fun get(index: Int): T
    operator fun set(index: Int, value: T): Unit

    operator fun iterator(): Iterator<T>
    // ...
}
```

创建一个数组,可以使用内置方法 `arrayOf()`,然后给它入各项的值, 例如:  `arrayOf(1, 2, 3)` 创建了一个数组 `[1, 2, 3]`.
也可以使用 `arrayOfNulls()` 库方法创建一个保存了null值元素的数组.

另一个操作就是使用`Array`的构造方法创建数组,它接受一个数组大小和一个方法用以根据索引返回各项的初始值.

```kotlin
fun main(args: Array<String>) {

    // Creates an Array<String> with values ["0", "1", "4", "9", "16"]
    val asc = Array(5, { i -> (i * i).toString() })
    asc.forEach { println(it) }

}
```

如上所述，[]操作代表对成员函数get（）和set（）的调用。

注意：与Java不同，Kotlin中的数组是不变的。这意味着Kotlin不允许我们将Array <String>分配给Array <Any>，
这可以防止可能的运行时故障（但您可以使用`Array <out Any>`，请参阅[Type Projections](generics.html#type-projections)).

Kotlin还有专门的类来表示原始类型的数组而没有装箱开销：ByteArray，ShortArray，IntArray等。
这些类与Array类没有继承关系，但它们具有相同的方法和属性集。
他们各自也有相应的工厂功能：

```kotlin
val x: IntArray = intArrayOf(1, 2, 3)
x[0] = x[1] + x[2]
```

## Unsigned integers 无符号整型

> Unsigned types are available only since Kotlin 1.3 and currently are *experimental*. See details [below](#experimental-status-of-unsigned-integers) 
{:.note}

Kotlin introduces following types for unsigned integers:

* `kotlin.UByte`: an unsigned 8-bit integer, ranges from 0 to 255
* `kotlin.UShort`: an unsigned 16-bit integer, ranges from 0 to 65535
* `kotlin.UInt`: an unsigned 32-bit integer, ranges from 0 to 2^32 - 1
* `kotlin.ULong`: an unsigned 64-bit integer, ranges from 0 to 2^64 - 1

Unsigned types support most of the operations of their signed counterparts.

> Note that changing type from unsigned type to signed counterpart (and vice versa) is a *binary incompatible* change
{:.note}

Unsigned types are implemented using another experimental feature, namely [inline classes](inline-classes.html).

---

> 无符号类型从 Kotlin 1.3 开始,并且仍然在 *实验功能*. 详请请看:[实验功能-无符号类型详情](#experimental-status-of-unsigned-integers) 

Kotlin 有以下类型的无符号整数 :

* `kotlin.UByte`: 无符号 8位 整数, 范围从0 到 255
* `kotlin.UShort`: 无符号 16位 整数, 范围从0 到 65535
* `kotlin.UInt`: 无符号 32位 整数, 范围从0 到 2^32 - 1
* `kotlin.ULong`: 无符号 64位 整数, 范围从0 到 2^64 - 1

无符号整数支持大多数有符号整数的运算符

>警告: 无符号向有符号的转换是不被允许的转换.

无符号被实现已使用另一个*实验功能*, 名字叫 [inline classes](inline-classes.html).

### Specialized classes 无符号类型专用类

Same as for primitives, each of unsigned type has corresponding type that represents array, specialized for that unsigned type:

* `kotlin.UByteArray`: an array of unsigned bytes
* `kotlin.UShortArray`: an array of unsigned shorts
* `kotlin.UIntArray`: an array of unsigned ints
* `kotlin.ULongArray`: an array of unsigned longs

Same as for signed integer arrays, they provide similar API to `Array` class without boxing overhead. 

Also, [ranges and progressions](ranges.html) supported for `UInt` and `ULong` by classes `kotlin.ranges.UIntRange`, `kotlin.ranges.UIntProgression`, `kotlin.ranges.ULongRange`, `kotlin.ranges.ULongProgression` 

---


与原始类型相同，每个无符号类型都有对应的表示数组的类型，专门用于该无符号类型

* `kotlin.UByteArray`: an array of unsigned bytes
* `kotlin.UShortArray`: an array of unsigned shorts
* `kotlin.UIntArray`: an array of unsigned ints
* `kotlin.ULongArray`: an array of unsigned longs

和有符号整型一样,它们提供相似的 `Array` API,而没有装箱的开销. 

同时, [ranges and progressions](ranges.html)支持 `UInt` 和 `ULong` 在以下几个类中:

* `kotlin.ranges.UIntRange`
* `kotlin.ranges.UIntProgression`
* `kotlin.ranges.ULongRange`
* `kotlin.ranges.ULongProgression`

### Literals 无符号类型字面值

To make unsigned integers easier to use, Kotlin provides an ability to tag an integer literal with a suffix indicating a specific unsigned type (similarly to Float/Long):

* suffixes `u` and `U` tag literal as unsigned. Exact type will be determined based on the expected type. If no expected type is provided, `UInt` or `ULong` will be chosen based on the size of literal 

```kotlin
val b: UByte = 1u  // UByte, expected type provided
val s: UShort = 1u // UShort, expected type provided
val l: ULong = 1u  // ULong, expected type provided

val a1 = 42u // UInt: no expected type provided, constant fits in UInt
val a2 = 0xFFFF_FFFF_FFFFu // ULong: no expected type provided, constant doesn't fit in UInt
```

* suffixes `uL` and `UL` explicitly tag literal as unsigned long.

```kotlin
val a = 1UL // ULong, even though no expected type provided and constant fits into UInt
```

---

为了使无符号整数更易于使用，Kotlin提供了一种标记整数文字的功能，后缀表示特定的无符号类型（类似于Float / Long）

* 后缀 `u` 和 `U` 的字面值表示无符号类型.确切类型将根据预期类型确定. 如果没有预期的类型提供, `UInt` 或 `ULong` 将根据字面值长度被使用

```kotlin
val b: UByte = 1u  // UByte, expected type provided
val s: UShort = 1u // UShort, expected type provided
val l: ULong = 1u  // ULong, expected type provided

val a1 = 42u // UInt: no expected type provided, constant fits in UInt
val a2 = 0xFFFF_FFFF_FFFFu // ULong: no expected type provided, constant doesn't fit in UInt
```

* 后缀 `uL` 和 `UL` 显式表示无符号长整型.

```kotlin
val a = 1UL // ULong, even though no expected type provided and constant fits into UInt
```

### Experimental status of unsigned integers 无符号类型的实验状态

The design of unsigned types is experimental, meaning that this feature is moving fast and no compatibility guarantees are given. When using unsigned arithmetics in Kotlin 1.3+, warning will be reported, indicating that this feature is experimental. To remove warning, you have to opt-in for experimental usage of unsigned types. 

There are two possible ways to opt-in for unsigned types: with marking your API as experimental too, or without doing that.

- to propagate experimentality, either annotate declarations which use unsigned integers with `@ExperimentalUnsignedTypes` or pass `-Xexperimental=kotlin.ExperimentalUnsignedTypes` to the compiler (note that the latter will make *all* declaration in compiled module experimental)
- to opt-in without propagating experimentality, either annotate declarations with `@UseExperimental(ExperimentalUnsignedTypes::class)` or pass `-Xuse-experimental=kotlin.ExperimentalUnsignedTypes`

It's up to you to decide if your clients have to explicitly opt-in into usage of your API, but bear in mind that unsigned types are an experimental feature, so API which uses them can be suddenly broken due to changes in language. 

See also or Experimental API [KEEP](https://github.com/Kotlin/KEEP/blob/master/proposals/experimental.md) for technical details.

### Further discussion  更多关于无符号类型的讨论

See [language proposal for unsigned types](https://github.com/Kotlin/KEEP/blob/master/proposals/unsigned-types.md) for technical details and further discussion.

## Strings  字符串类型

Strings are represented by the type `String`. Strings are immutable.
Elements of a string are characters that can be accessed by the indexing operation: `s[i]`.
A string can be iterated over with a *for*{: .keyword }-loop:

```kotlin
fun main(args: Array<String>) {
    val str = "abcd"
    for (c in str) {
        println(c)
    }
}
```

You can concatenate strings using the `+` operator. This also works for concatenating strings with values of other types, as long
as the first element in the expression is a string:

```kotlin
fun main(args: Array<String>) {

    val s = "abc" + 1
    println(s + "def")

    }
```

Note that in most cases using [string templates](#string-templates) or raw strings is preferable to string concatenation.

---
字符串由String类型表示。字符串是不可变的。字符串的元素是可以通过索引操作访问的字符：s [i]。可以使用for循环迭代字符串

```kotlin
fun main(args: Array<String>) {
    val str = "abcd"
    for (c in str) {
        println(c)
    }
}
```

您可以使用+运算符连接字符串。这也适用于将字符串与其他类型的值连接，只要表达式中的第一个元素是字符串：

```kotlin
fun main(args: Array<String>) {

    val s = "abc" + 1
    println(s + "def")

    }
```

请注意，在大多数情况下，使用字符串模板或原始字符串比字符串连接更可取。

### String Literals 字符串字面量

Kotlin has two types of string literals: escaped strings that may have escaped characters in them and raw strings that can contain newlines and arbitrary text. An escaped string is very much like a Java string:

```kotlin
val s = "Hello, world!\n"
```

Escaping is done in the conventional way, with a backslash. See [Characters](#characters) above for the list of supported escape sequences.

A raw string is delimited by a triple quote (`"""`), contains no escaping and can contain newlines and any other characters:

```kotlin
val text = """
    for (c in "foo")
        print(c)
"""
```

You can remove leading whitespace with [`trimMargin()`](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.text/trim-margin.html) function:

```kotlin
val text = """
    |Tell me and I forget.
    |Teach me and I remember.
    |Involve me and I learn.
    |(Benjamin Franklin)
    """.trimMargin()
```

By default `|` is used as margin prefix, but you can choose another character and pass it as a parameter, like `trimMargin(">")`.

---

Kotlin有两种类型的字符串文字：可能在其中包含转义字符的转义字符串以及可以包含换行符和任意文本的原始字符串。转义字符串非常类似于Java字符串：

```kotlin
val s = "Hello, world!\n"
```

以传统方式进行转义，使用反斜杠。请参阅上面的字符以获取支持的转义序列列表。

原始字符串由三引号（“”“）分隔，不包含转义，可以包含换行符和任何其他字符：

```kotlin
val text = """
    for (c in "foo")
        print(c)
"""
```

您可以使用[`trimMargin()`](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.text/trim-margin.html) 函数删除前导空格：

```kotlin
val text = """
    |Tell me and I forget.
    |Teach me and I remember.
    |Involve me and I learn.
    |(Benjamin Franklin)
    """.trimMargin()
```

默认情况下|用作边距前缀，但您可以选择另一个字符并将其作为参数传递，如trimMargin（“>”）。

### String Templates 字符串模板

Strings may contain template expressions, i.e. pieces of code that are evaluated and whose results are concatenated into the string.
A template expression starts with a dollar sign ($) and consists of either a simple name:

```kotlin
fun main(args: Array<String>) {
    val i = 10
    println("i = $i") // prints "i = 10"

    }
```

or an arbitrary expression in curly braces:

```kotlin
fun main(args: Array<String>) {
    val s = "abc"
    println("$s.length is ${s.length}") // prints "abc.length is 3"

    }
```

Templates are supported both inside raw strings and inside escaped strings.
If you need to represent a literal `$` character in a raw string (which doesn't support backslash escaping), you can use the following syntax:

```kotlin
val price = """
${'$'}9.99
"""
```

--

字符串可以包含模板表达式，即评估的代码片段，其结果连接到字符串中。模板表达式以美元符号（$）开头，由一个简单的名称组成：

```kotlin
fun main(args: Array<String>) {
    val i = 10
    println("i = $i") // prints "i = 10"

    }
```

或花括号中的任意表达式：

```kotlin
fun main(args: Array<String>) {
    val s = "abc"
    println("$s.length is ${s.length}") // prints "abc.length is 3"

    }
```

原始字符串和转义字符串内部都支持模板。如果需要在原始字符串中表示文字$字符（不支持反斜杠转义），可以使用以下语法：

```kotlin
val price = """
${'$'}9.99
"""
```