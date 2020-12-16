---
title: 0025-Kotlin语言探索004-(package & import)
date: 2018-11-13 11:44:58
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
# Packages(包)

A source file may start with a package declaration:

```kotlin
package foo.bar

fun baz() { ... }
class Goo { ... }

// ...
```

All the contents (such as classes and functions) of the source file are contained by the package declared.
So, in the example above, the full name of `baz()` is `foo.bar.baz`, and the full name of `Goo` is `foo.bar.Goo`. 
 
If the package is not specified, the contents of such a file belong to "default" package that has no name.

---

源文件的开头可能是`包声明`语句:

```kotlin
package foo.bar

fun baz() { ... }
class Goo { ... }

// ...
```

所有源文件里的内容都包含在`包声明`之下.所以,在上面的例子中,`baz()`的全名是`foo.bar.baz`,`Goo` 的全名是 `foo.bar.Goo`. 

如果没有定义`包名`,源文件里的内容则属于没有名字的默认包名.

<!--- more --->

## Default Imports (缺省的导入包)

A number of packages are imported into every Kotlin file by default:

- [kotlin.*](/api/latest/jvm/stdlib/kotlin/index.html)
- [kotlin.annotation.*](/api/latest/jvm/stdlib/kotlin.annotation/index.html)
- [kotlin.collections.*](/api/latest/jvm/stdlib/kotlin.collections/index.html)
- [kotlin.comparisons.*](/api/latest/jvm/stdlib/kotlin.comparisons/index.html)  (since 1.1)
- [kotlin.io.*](/api/latest/jvm/stdlib/kotlin.io/index.html)
- [kotlin.ranges.*](/api/latest/jvm/stdlib/kotlin.ranges/index.html)
- [kotlin.sequences.*](/api/latest/jvm/stdlib/kotlin.sequences/index.html)
- [kotlin.text.*](/api/latest/jvm/stdlib/kotlin.text/index.html)

Additional packages are imported depending on the target platform:

- JVM:
  - java.lang.*
  - [kotlin.jvm.*](/api/latest/jvm/stdlib/kotlin.jvm/index.html)

- JS:    
  - [kotlin.js.*](/api/latest/jvm/stdlib/kotlin.js/index.html)

---

一下的几个包都会在每个Kotlin文件中自动导入:

- [kotlin.*](/api/latest/jvm/stdlib/kotlin/index.html)
- [kotlin.annotation.*](/api/latest/jvm/stdlib/kotlin.annotation/index.html)
- [kotlin.collections.*](/api/latest/jvm/stdlib/kotlin.collections/index.html)
- [kotlin.comparisons.*](/api/latest/jvm/stdlib/kotlin.comparisons/index.html)  (since 1.1)
- [kotlin.io.*](/api/latest/jvm/stdlib/kotlin.io/index.html)
- [kotlin.ranges.*](/api/latest/jvm/stdlib/kotlin.ranges/index.html)
- [kotlin.sequences.*](/api/latest/jvm/stdlib/kotlin.sequences/index.html)
- [kotlin.text.*](/api/latest/jvm/stdlib/kotlin.text/index.html)

附加的`包`取决于不同的编译环境自动导入:

- JVM:
  - java.lang.*
  - [kotlin.jvm.*](/api/latest/jvm/stdlib/kotlin.jvm/index.html)

- JS:    
  - [kotlin.js.*](/api/latest/jvm/stdlib/kotlin.js/index.html)

## Imports (导包)

Apart from the default imports, each file may contain its own import directives.
Syntax for imports is described in the [grammar](grammar.html#import).
除了缺省导包,每个文件可能需要自己特定的导包.kotlin语法中使用`import`关键字导包[grammar](grammar.html#import).

We can import either a single name, e.g.
我们可以导入一个具体的函数名字, 例如:

```kotlin
import foo.Bar // Bar is now accessible without qualification
```

or all the accessible contents of a scope (package, class, object etc):
或者导入一个类或包里所有的内容使用`*` (package, class, object etc):

```kotlin
import foo.* // everything in 'foo' becomes accessible
```

If there is a name clash, we can disambiguate by using *as* keyword to locally rename the clashing entity:
遇到名字重复的情况可以使用`as`关键字改名字

```kotlin
import foo.Bar // Bar is accessible
import bar.Bar as bBar // bBar stands for 'bar.Bar'
```

The `import` keyword is not restricted to importing classes; you can also use it to import other declarations:
`import`关键字不仅可以导入类,也可以用它导入其他的声明:

* top-level functions and properties; 顶级的方法和声明
* functions and properties declared in [object declarations](object-declarations.html#object-declarations);在对象声明里的方法和声明.
* [enum constants](enum-classes.html).枚举对象

Unlike Java, Kotlin does not have a separate ["import static"](https://docs.oracle.com/javase/8/docs/technotes/guides/language/static-import.html) syntax; all of these declarations are imported using the regular `import` keyword.
和Java不同, Kotlin没有 ["import static"](https://docs.oracle.com/javase/8/docs/technotes/guides/language/static-import.html) 这种用法,所有的情况都是用标准的`import` 关键字完成.

## Visibility of Top-level Declarations(顶级声明的可见性)

If a top-level declaration is marked *private*{: .keyword }, it is private to the file it's declared in (see [Visibility Modifiers](visibility-modifiers.html)).
如果顶级声明被标记为`private`他只是对声明他的文件私有了而已.

