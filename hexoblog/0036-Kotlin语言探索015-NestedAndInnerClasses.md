---
title: 0036-Kotlin语言探索015-NestedAndInnerClasses
date: 2018-11-21 19:44:35
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
# Nested and Inner Classes

Classes can be nested in other classes:


```kotlin
class Outer {
    private val bar: Int = 1
    class Nested {
        fun foo() = 2
    }
}

val demo = Outer.Nested().foo() // == 2
```


## Inner classes

A class may be marked as *inner*{: .keyword } to be able to access members of outer class. Inner classes carry a reference to an object of an outer class:


```kotlin
class Outer {
    private val bar: Int = 1
    inner class Inner {
        fun foo() = bar
    }
}

val demo = Outer().Inner().foo() // == 1
```


See [Qualified *this*{: .keyword } expressions](this-expressions.html) to learn about disambiguation of *this*{: .keyword } in inner classes.

## Anonymous inner classes

Anonymous inner class instances are created using an [object expression](object-declarations.html#object-expressions):


```kotlin
window.addMouseListener(object: MouseAdapter() {

    override fun mouseClicked(e: MouseEvent) { ... }
                                                                                                            
    override fun mouseEntered(e: MouseEvent) { ... }
})
```


If the object is an instance of a functional Java interface (i.e. a Java interface with a single abstract method),
you can create it using a lambda expression prefixed with the type of the interface:


```kotlin
val listener = ActionListener { println("clicked") }
```

