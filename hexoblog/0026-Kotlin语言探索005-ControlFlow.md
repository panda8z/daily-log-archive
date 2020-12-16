---
title: 0026-Kotlin语言探索005-ControlFlow
date: 2018-11-13 14:26:34
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

# Control Flow: if, when, for, while **控制流: `if`, `when`, `for`, `while`**

## 1. If Expression `if`**表达式**

In Kotlin, *if* is an expression, i.e. it returns a value.
Therefore there is no ternary operator (condition ? then : else), because ordinary *if* works fine in this role.

在`Kotlin`中,`if`关键字是一个**表达式**,也就是它会返回一个值.
然后`Kotlin`中没有三元运算符(条件? 然后: 否则),因为普通的`if`表达式就能很好的胜任这个工作.

```kotlin
// Traditional usage 
var max = a 
if (a < b) max = b

// With else 
var max: Int
if (a > b) {
    max = a
} else {
    max = b
}
// As expression
val max = if (a > b) a else b
```

*if* branches can be blocks, and the last expression is the value of a block:

`if`的分支可以是一个语句块, 并且最终的表达式的值是语句块返回的值:

```kotlin
val max = if (a > b) {
    print("Choose a")
    a
} else {
    print("Choose b")
    b
}
```

If you're using *if* as an expression rather than a statement (for example, returning its value or
assigning it to a variable), the expression is required to have an `else` branch.

如果`if`关键字作为一个表达式而不是作为一个语句(将值返回给你一个变量的语句),表达式需要包含`eles`分支.

<!--- more --->

## When Expression

*when* replaces the switch operator of C-like languages. In the simplest form it looks like this

`when`替代了类C语言的`switch`语句,它看起来更加简洁:

```kotlin
when (x) {
    1 -> print("x == 1")
    2 -> print("x == 2")
    else -> { // Note the block
        print("x is neither 1 nor 2")
    }
}
```

*when* matches its argument against all branches sequentially until some branch condition is satisfied.
*when* can be used either as an expression or as a statement. If it is used as an expression, the value
of the satisfied branch becomes the value of the overall expression. If it is used as a statement, the values of
individual branches are ignored. (Just like with *if*, each branch can be a block, and its value
is the value of the last expression in the block.)

`when`表达式逐个匹配每个分支的条件知道某个分支的条件成立.
`when`可以作为一个表达式或者一个语句,当作为表达式的时候,条件匹配的分支的值就是表达式的值.(和`if`表达式一样,表达式的值是分支代码块最后一条语句的值)

The *else* branch is evaluated if none of the other branch conditions are satisfied.
If *when* is used as an expression, the *else* branch is mandatory,
unless the compiler can prove that all possible cases are covered with branch conditions (as, for example, with `enum` classentries and `sealed` classsubtypes).

如果不满足其他任何分支条件，则执行`else`分支。如果when用作表达式，则else分支是必需的，除非编译器能够证明所有可能的情况都包含在分支条件中

If many cases should be handled in the same way, the branch conditions may be combined with a comma:

如果应以相同方式处理许多情况，则可以将分支条件与逗号组合:

```kotlin
when (x) {
    0, 1 -> print("x == 0 or x == 1")
    else -> print("otherwise")
}
```

We can use arbitrary expressions (not only constants) as branch conditions.

我们可以使用任意表达式（不仅仅是常量）作为分支条件:

```kotlin
when (x) {
    parseInt(s) -> print("s encodes x")
    else -> print("s does not encode x")
}
```

We can also check a value for being *in* or *!in* a `range` or a `collection`:

我们还可以检查一个值是否在(`in`)或不在(`!in`)某个范围或集合中:

```kotlin
when (x) {
    in 1..10 -> print("x is in the range")
    in validNumbers -> print("x is valid")
    !in 10..20 -> print("x is outside the range")
    else -> print("none of the above")
}
```

Another possibility is to check that a value *is* or *!is* of a particular type. Note that,
due to `smart` , you can access the methods and properties of the type without
any extra checks.

另一种可能性是检查值是(`is`)或不是(`!is`)某种特定类型的值。请注意，由于智能强制转换，您可以访问该类型的方法和属性，而无需任何额外的检查。

```kotlin
fun hasPrefix(x: Any) = when(x) {
    is String -> x.startsWith("prefix")
    else -> false
}
```

*when* can also be used as a replacement for an *if*-*else* *if* chain.

If no argument is supplied, the branch conditions are simply boolean expressions, and a branch is executed when its condition is true:

`when`也可以用作`if-else if`的替代语句。
如果没有提供参数，则分支条件只是布尔表达式，并且当条件为真时执行分支

```kotlin
when {
    x.isOdd() -> print("x is odd")
    x.isEven() -> print("x is even")
    else -> print("x is funny")
}
```

Since Kotlin 1.3, it is possible to capture *when*{: .keyword} subject in a variable using following syntax:

```kotlin
fun Request.getBody() =
        when (val response = executeRequest()) {
            is Success -> response.body
            is HttpError -> throw HttpException(response.status)
        }
```

Scope of variable, introduced in *when*{: .keyword} subject, is restricted to *when*{: .keyword} body.

See the grammar for `when`

## For Loops

*for* loop iterates through anything that provides an iterator. This is equivalent
to the `foreach` loop in languages like C#. The syntax is as follows:

```kotlin
for (item in collection) print(item)
```

The body can be a block.

```kotlin
for (item: Int in ints) {
    // ...
}
```

As mentioned before, *for* iterates through anything that provides an iterator, i.e.

* has a member- or extension-function `iterator()`, whose return type
  * has a member- or extension-function `next()`, and
  * has a member- or extension-function `hasNext()` that returns `Boolean`.

All of these three functions need to be marked as `operator`.

To iterate over a range of numbers, use a `range` expression]

```kotlin
fun main(args: Array<String>) {
//sampleStart
for (i in 1..3) {
    println(i)
}
for (i in 6 downTo 0 step 2) {
    println(i)
}
//sampleEnd
}
```

A `for` loop over a range or an array is compiled to an index-based loop that does not create an iterator object.

If you want to iterate through an array or a list with an index, you can do it this way:



```kotlin
fun main(args: Array<String>) {
val array = arrayOf("a", "b", "c")
//sampleStart
for (i in array.indices) {
    println(array[i])
}
//sampleEnd
}
```

Alternatively, you can use the `withIndex` library function:

```kotlin
fun main(args: Array<String>) {
val array = arrayOf("a", "b", "c")
//sampleStart
for ((index, value) in array.withIndex()) {
    println("the element at $index is $value")
}
//sampleEnd
}
```


## While Loops

*while* and *do*..*while* work as usual

```kotlin
while (x > 0) {
    x--
}

do {
    val y = retrieveData()
} while (y != null) // y is visible here!
```

## Break and continue in loops |  **`break`和`continue`语句**

Kotlin支持传统的在循环中使用的`break`和`continue`语句,详情请看下一节:
{% post_link 0027-Kotlin语言探索006-ReturnsAndJumps 0027-Kotlin语言探索006-ReturnsAndJumps %}

Kotlin supports traditional *break* and *continue* operators in loops. 
