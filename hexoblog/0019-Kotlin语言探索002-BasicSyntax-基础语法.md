---
title: Kotlin语言探索002-BasicSyntax-基础语法
date: 2018-10-29 12:45:04
updated: 2018-11-06 21:34:15
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

# 基础语法

## **1. `packages`(包) 的定义**

和java一样,包的定义语句一定要放在 源码文件的首行.

使用`package`关键字.

```Kotlin
package my.demo

import java.util.*

// ...
```


## **2. 函数定义**

1. 定义一个有两个`Int`类型参数并返回`Int`类型的函数.

    ```Kotlin
    fun sum(a: Int, b: Int): Int {
        return a + b
    }

    fun main(args: Array<String>) {
        print("sum of 3 and 5 is ")
        println(sum(3, 5))
    }

    ```

2. 一个函数体是一个表达式的函数,返回可以被自动判断.

    ```Kotlin
    fun sum(a: Int, b: Int) = a + b

    fun main(args: Array<String>) {
        println("sum of 19 and 23 is ${sum(19, 23)}")
    }
    ```

3. 函数可以返回一个没有意义的类型值

    ```Kotlin
    fun printSum(a: Int, b: Int): Unit {
        println("sum of $a and $b is ${a + b}")
    }

    fun main(args: Array<String>) {
        printSum(-1, 8)
    }

    ```

4. `Unit`这种无意义的类型时 ,可以省略,即不指定返回类型

    ```Kotlin
    fun printSum(a: Int, b: Int) {
    println("sum of $a and $b is ${a + b}")
    }
    ```

## **3. 定义变量**

1. 只读变量的定义(Assign-once (read-only) local variable):

    ```Kotlin
    fun main(args: Array<String>) {
        val a: Int = 1  // immediate assignment  立即分配
        val b = 2   // `Int` type is inferred Int类型会被自动推断出来
        val c: Int  
        // Type required when no initializer is provided 没有初始化前要指定类型.
        c = 3       // deferred assignment 延迟分配
        println("a = $a, b = $b, c = $c")
    }
    ```

2. 易变变量的定义(Mutable variable)

    ```Kotlin
    fun main(args: Array<String>) {
        var x = 5 // `Int` type is inferred
        x += 1
        println("x = $x")
    }
    ```

3. 顶级变量(Top-level variables):

    ```Kotlin
    val PI = 3.14
    var x = 0

    fun incrementX() {
        x += 1
    }

    fun main(args: Array<String>) {
        println("x = $x; PI = $PI")
        incrementX()
        println("incrementX()")
        println("x = $x; PI = $PI")
    }
    ```

 <!-- more -->

## **4. 注释(Comments)**

Just like Java and JavaScript, Kotlin supports end-of-line and block comments.
想`Java`和`JavaScript`一样,`Kotlin`支持 单行注释和块级注释.

```Kotlin

// This is an end-of-line comment
// 这是一个单行注释
/* This is a block comment
   on multiple lines. */

/*
 * 这是一个块级注释
 * 中间可以换行
 */

```


Unlike Java, block comments in Kotlin can be nested.
和`Java`不同的一点, 在`Kotlin`中块级注释可以嵌套 
> ⚠️ 没搞清楚它是怎么嵌套的.....


## **5. 使用字符串模板(Using string templates)**

```Kotlin
fun main(args: Array<String>) {
    var a = 1
    // simple name in template:
    val s1 = "a is $a" 
    a = 2
    // arbitrary expression in template:
    val s2 = "${s1.replace("is", "was")}, but now is $a"
    println(s2)
}
```


## **6. 使用条件语句(Using conditional expressions)**

```Kotlin
fun maxOf(a: Int, b: Int): Int {
    if (a > b) {
        return a
    } else {
        return b
    }
}
```

 使用`if`表达式(Using if as an expression):

```Kotlin
fun maxOf(a: Int, b: Int) = if (a > b) a else b

fun main(args: Array<String>) {
    println("max of 0 and 42 is ${maxOf(0, 42)}")
}
```


## **7. 使用`nullable`值和检查`null`值(Using nullable values and checking for null)

A reference must be explicitly marked as nullable when null value is possible.
一个引用而已返回`null`时必须显式声明.

当字符串`String`不能转换成`Int`值时返回`null`:


```Kotlin
fun parseInt(str: String): Int? {
    // ...
}
```

使用一个可返回`null`值得函数:

```Kotlin
fun parseInt(str: String): Int? {
    return str.toIntOrNull()
}
fun printProduct(arg1: String, arg2: String) {
    val x = parseInt(arg1)
    val y = parseInt(arg2)
​
    // Using `x * y` yields error because they may hold nulls.
    if (x != null && y != null) {
        // x and y are automatically cast to non-nullable after null check
        println(x * y)
    }
    else {
        println("either '$arg1' or '$arg2' is not a number")
    }    
}
fun main(args: Array<String>) {
    printProduct("6", "7")
    printProduct("a", "7")
    printProduct("a", "b")
}
```

或者

```Kotlin
// ...
if (x == null) {
    println("Wrong number format in arg1: '$arg1'")
    return
}
if (y == null) {
    println("Wrong number format in arg2: '$arg2'")
    return
}
​
// x and y are automatically cast to non-nullable after null check
println(x * y)
```


## **8. 类型检查和类型自动转换(Using type checks and automatic casts)**

The `is` operator checks if an expression is an instance of a type. If an immutable local variable or property is checked for a specific type, there's no need to cast it explicitly:

`is` 操作符可以检查表达式是否是某种类型的实例. 如果本地变量或者属性是指定的类型,将不用显式转换.

```Kotlin
fun getStringLength(obj: Any): Int? {
    if (obj is String) {
        // `obj` is automatically cast to `String` in this branch
        // Any类型自动被转换成String类型
        return obj.length
    }
​
    // `obj` is still of type `Any` outside of the type-checked branch
    // obj 在类型判断语句外面的时候依然是 Any类型
    return null
}

fun main(args: Array<String>) {
    fun printLength(obj: Any) {
        println("'$obj' string length is ${getStringLength(obj) ?: "... err, not a string"} ")
    }
    printLength("Incomprehensibilities")
    printLength(1000)
    printLength(listOf(Any()))
}

```

or

```Kotlin
fun getStringLength(obj: Any): Int? {
    if (obj !is String) return null
​
    // `obj` is automatically cast to `String` in this branch
    return obj.length
}
```
or even

```Kotlin
fun getStringLength(obj: Any): Int? {
    // `obj` is automatically cast to `String` on the right-hand side of `&&`
    if (obj is String && obj.length > 0) {
        return obj.length
    }
​
    return null
}
```


## **9. `For`循环(Using a for loop)**

```Kotlin
val items = listOf("apple", "banana", "kiwifruit")
for (item in items) {
    println(item)
}
```
or

```Kotlin
val items = listOf("apple", "banana", "kiwifruit")
for (index in items.indices) {
    println("item at $index is ${items[index]}")
}
```


## **10. `While`循环（Using a while loop）**

```Kotlin
val items = listOf("apple", "banana", "kiwifruit")
var index = 0
while (index < items.size) {
    println("item at $index is ${items[index]}")
    index++
}
```


## **11. `when`表达式 （ Using when expression）**


`when` 表达式像java里面的`switch`语句一样.
但是函数式编程写法看起来挺棒.

```Kotlin
fun describe(obj: Any): String =
    when (obj) {
        1          -> "One"
        "Hello"    -> "Greeting"
        is Long    -> "Long"
        !is String -> "Not a string"
        else       -> "Unknown"
    }
```


## **12. Using ranges**

使用`in`操作符检查一个数字是否在给定的范围之内:Check if a number is within a range using in operator:

```Kotlin
val x = 10
val y = 9
if (x in 1..y+1) {
    println("fits in range")
}
```

检查一个数字是否在给定的范围之外:Check if a number is out of range:


```Kotlin
val list = listOf("a", "b", "c")
​
if (-1 !in 0..list.lastIndex) {
    println("-1 is out of range")
}
if (list.size !in list.indices) {
    println("list size is out of valid list indices range too")
}
```
在一个范围内迭代 Iterating over a range:


```Kotlin
for (x in 1..5) {
    print(x)
}
```

或者 指定进度: or over a progression:

```
for (x in 1..10 step 2) {
    print(x)
}
println()
for (x in 9 downTo 0 step 3) {
    print(x)
}
```


## **13. 使用集合（Using collections)**

迭代一个集合:Iterating over a collection:


```Kotlin
for (item in items) {
    println(item)
}
```

使用 `in` 操作符 检索集合中是否包含某个对象: Checking if a collection contains an object using in operator:

```
when {
    "orange" in items -> println("juicy")
    "apple" in items -> println("apple is fine too")
}
```

使用`lambda`表达式迭代，遍历集合: Using lambda expressions to filter and map collections:


```Kotlin
val fruits = listOf("banana", "avocado", "apple", "kiwifruit")
fruits
  .filter { it.startsWith("a") }
  .sortedBy { it }
  .map { it.toUpperCase() }
  .forEach { println(it) }
Target platform: JVMRunning on kotlin v. 1.2.71
```

## **14. 创建基本类的实例 （Creating basic classes and their instances）**

    ```Kotlin
    val rectangle = Rectangle(5.0, 2.0) //no 'new' keyword required
    val triangle = Triangle(3.0, 4.0, 5.0)
    ```
