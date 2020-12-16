---
title: 0043-Kotlin语言探索022-kotlin-test
date: 2018-11-21 19:57:53
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

# kotlin.test

kotlin.test library provides annotations to mark test functions and a set of utility functions for performing assertions in tests, independently of the test framework being used.

The test framework is abstracted through the Asserter class. A basic Asserter implementation is provided out of the box. Note that the class is not intended to be used directly from tests, use instead the top-level assertion functions which delegate to the Asserter.

The library consists of the modules:

kotlin-test-common – assertions for use in common code;
kotlin-test-annotations-common – test annotations for use in common code;
kotlin-test – a JVM implementation of assertions from kotlin-test-common;
kotlin-test-junit – provides an implementation of Asserter on top of JUnit and maps the test annotations from kotlin-test-annotations-common to JUnit test annotations;
kotlin-test-junit5 – provides an implementation of Asserter on top of JUnit 5 and maps the test annotations from kotlin-test-annotations-common to JUnit 5 test annotations;
kotlin-test-testng – provides an implementation of Asserter on top of TestNG and maps the test annotations from kotlin-test-annotations-common to TestNG test annotations;
kotlin-test-js – a JS implementation of common test assertions and annotations with the out-of-the-box support for Jasmine, Mocha, and Jest testing frameworks, and an experimental way to plug in a custom unit testing framework.
Packages
Common
JVM
JUnit
JUnit5
TestNG
JS
Native
1.0
## kotlin.test
Annotations to mark test functions and top-level functions for performing assertions in tests.

JUnit
1.0
## kotlin.test.junit
JUnit5
1.0
## kotlin.test.junit5
TestNG
1.0
## kotlin.test.testng
