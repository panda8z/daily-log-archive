---
title: 0042-Kotlin语言探索021-KotlinStantardLibrary
date: 2018-11-21 19:57:45
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

# Kotlin Standard Library

The Kotlin Standard Library provides living essentials for everyday work with kotlin. These include:

Higher-order functions implementing idiomatic patterns (let, apply, use, synchronized, etc).
Extension functions providing querying operations for collections (eager) and sequences (lazy).
Various utilities for working with strings and char sequences.
Extensions for JDK classes making it convenient to work with files, IO, and threading.

# Packages

Common
JVM
JS
Native
1.0
## kotlin
Core functions and types, available on all supported platforms.

Common
JVM
JS
Native
1.0
## kotlin.annotation
Library support for the Kotlin annotation facility.

JS
1.1
## kotlin.browser
Access to top-level properties (document, window etc.) in the browser environment.

Common
JVM
JS
Native
1.0
## kotlin.collections
Collection types, such as Iterable, Collection, List, Set, Map and related top-level and extension functions.

Common
JVM
JS
Native
1.0
## kotlin.comparisons
Helper functions for creating Comparator instances.

JVM
1.0
## kotlin.concurrent
Utility functions for concurrent programming.

Common
JVM
JS
Native
1.3
## kotlin.contracts
Experimental DSL for declaring custom function contracts.

Common
JVM
JS
Native
1.3
## kotlin.coroutines
Basic primitives for creating and suspending coroutines: Continuation, CoroutineContext interfaces, coroutine creation and suspension top-level functions.

Common
JVM
JS
Native
1.1
## kotlin.coroutines.experimental
Deprecated support for experimental coroutines, provided for compatibility. It's recommended to migrate to ## kotlin.coroutines API.

Common
JVM
JS
Native
1.1
## kotlin.coroutines.experimental.intrinsics
Deprecated support for experimental coroutines, provided for compatibility. It's recommended to migrate to ## kotlin.coroutines.intrinsics API.

Common
JVM
JS
Native
1.3
## kotlin.coroutines.intrinsics
Low-level building blocks for libraries that provide coroutine-based APIs.

JS
1.1
## kotlin.dom
Utility functions for working with the browser DOM.

Common
JVM
JS
Native
1.1
## kotlin.experimental
Experimental APIs, subject to change in future versions of ## kotlin.

Common
JVM
JS
Native
1.0
## kotlin.io
IO API for working with files and streams.

Common
JS
1.0
## kotlin.js
Functions and other APIs specific to the JavaScript platform.

Common
JVM
JS
1.0
## kotlin.jvm
Functions and annotations specific to the Java platform.

Common
JVM
JS
Native
1.2
## kotlin.math
Mathematical functions and constants.

Native
1.3
## kotlin.native
Native
1.3
## kotlin.native.concurrent
Native
1.3
## kotlin.native.ref
Common
JVM
JS
Native
1.0
## kotlin.properties
Standard implementations of delegates for delegated properties and helper functions for implementing custom delegates.

Common
JVM
JS
Native
1.3
## kotlin.random
Provides the default generator of pseudo-random values, the repeatable generator, and a base class for other RNG implementations.

Common
JVM
JS
Native
1.0
## kotlin.ranges
Ranges, Progressions and related top-level and extension functions.

Common
JVM
JS
Native
1.0
## kotlin.reflect
Runtime API for Kotlin reflection

JVM
1.1
## kotlin.reflect.full
Extensions for Kotlin reflection provided by kotlin-reflect library.

JVM
1.0
## kotlin.reflect.jvm
Runtime API for interoperability between Kotlin reflection and Java reflection provided by kotlin-reflect library.

Common
JVM
JS
Native
1.0
## kotlin.sequences
Sequence type that represents lazily evaluated collections. Top-level functions for instantiating sequences and extension functions for sequences.

JVM
JRE8
1.2
## kotlin.streams
Utility functions for working with Java 8 streams.

JVM
Native
1.0
## kotlin.system
System-related utility functions.

Common
JVM
JS
Native
1.0
## kotlin.text
Functions for working with text and regular expressions.

Native
1.3
kotlinx.cinterop
Native
1.3
kotlinx.wasm.jsinterop
JS
1.1
org.khronos.webgl
Kotlin JavaScript wrappers for the WebGL API.

JS
1.1
org.w3c.dom
Kotlin JavaScript wrappers for the DOM API.

JS
1.1
org.w3c.dom.css
Kotlin JavaScript wrappers for the DOM CSS API.

JS
1.1
org.w3c.dom.events
Kotlin JavaScript wrappers for the DOM events API.

JS
1.1
org.w3c.dom.parsing
Kotlin JavaScript wrappers for the DOM parsing API.

JS
1.1
org.w3c.dom.svg
Kotlin JavaScript wrappers for the DOM SVG API.

JS
1.1
org.w3c.dom.url
Kotlin JavaScript wrappers for the DOM URL API.

JS
1.1
org.w3c.fetch
Kotlin JavaScript wrappers for the W3C fetch API.

JS
1.1
org.w3c.files
Kotlin JavaScript wrappers for the W3C file API.

JS
1.1
org.w3c.notifications
Kotlin JavaScript wrappers for the Web Notifications API.

JS
1.1
org.w3c.performance
Kotlin JavaScript wrappers for the Navigation Timing API.

JS
1.1
org.w3c.workers
Kotlin JavaScript wrappers for the Web Workers API.

JS
1.1
org.w3c.xhr
Kotlin JavaScript wrappers for the XMLHttpRequest API.

