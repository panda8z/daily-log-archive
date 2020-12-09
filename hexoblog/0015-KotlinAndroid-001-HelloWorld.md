---
title: KotlinAndroid-001-HelloWorld
date: 2018-10-24 23:15:06
updated: 2018-11-06 21:34:34
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
# KotlinAndroid-001-HelloWorld

## 1. 创建Kotlin支持的Android工程

{% asset_img  QQ20181025-001116.jpg  QQ20181025-001116.jpg %}

## 2.各重要文件里的kotlin

1. Project的build.gradle

在 Project级别的build.gradle里面多了一个classpath`classpath "org.jetbrains.kotlin:kotlin-gradle-plugin:$kotlin_version"`.

**kotlin_version**使用`ext.kotlin_version='1.2.30'`定义.

```groovy

// Top-level build file where you can add configuration options common to all sub-projects/modules.

buildscript {
    ext.kotlin_version = '1.2.30'
    repositories {
        google()
        jcenter()
    }
    dependencies {
        classpath 'com.android.tools.build:gradle:3.1.2'
        classpath "org.jetbrains.kotlin:kotlin-gradle-plugin:$kotlin_version"

        // NOTE: Do not place your application dependencies here; they belong
        // in the individual module build.gradle files
    }
}

allprojects {
    repositories {
        google()
        jcenter()
    }
}

task clean(type: Delete) {
    delete rootProject.buildDir
}

```

<!-- more -->

2. Module级别的build.gradle

在这个级别的`build.gradle`里面多了两个插件和一个标准库

- 两个插件
  
    `apply plugin: 'kotlin-android'`

    `apply plugin: 'kotlin-android-extensions'`

- 一个标准库

    `implementation"org.jetbrains.kotlin:kotlin-stdlib-jre7:$kotlin_version"`

```groovy

apply plugin: 'com.android.application'

apply plugin: 'kotlin-android'

apply plugin: 'kotlin-android-extensions'

android {
    compileSdkVersion 27
    defaultConfig {
        applicationId "top.moreme.pandakotlindemo.pandakotlindemo"
        minSdkVersion 21
        targetSdkVersion 27
        versionCode 1
        versionName "1.0"
        testInstrumentationRunner "android.support.test.runner.AndroidJUnitRunner"
    }
    buildTypes {
        release {
            minifyEnabled false
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
        }
    }
}

dependencies {
    implementation fileTree(dir: 'libs', include: ['*.jar'])
    implementation"org.jetbrains.kotlin:kotlin-stdlib-jre7:$kotlin_version"
    implementation 'com.android.support:appcompat-v7:27.1.1'
    implementation 'com.android.support.constraint:constraint-layout:1.1.3'
    testImplementation 'junit:junit:4.12'
    androidTestImplementation 'com.android.support.test:runner:1.0.2'
    androidTestImplementation 'com.android.support.test.espresso:espresso-core:3.0.2'
}


```

## 3. 核心Activity代码部分

**1. `MainActivity.kt`**


**源码:**
```kotlin

package top.moreme.pandakotlindemo.pandakotlindemo

import android.support.v7.app.AppCompatActivity
import android.os.Bundle

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }
}

```

通过源码可以得到以下信息:

1. package 分包机制和java相同.
2. import 导报包机制和java相同.
3. 类声明不用加public,直接 `class + [类名] {    }`即可.
4. 类的继承使用冒号`:` 代替了 `extends`.
5. 类的继承被继承的类要在类名后面加`()`.
6. 函数声明使用: `fun`前缀 + `[类名]`+ `(参数列表)`+ `{ }`形式.
7. 函数重写直接在`fun`前缀前面加`override`, 格式: `override fun [类名](参数列表) {    }`.
8. 变量类型使用: `[变量名] : [类型]?`格式声明.
9. 函数中调用父类同名函数和java一样使用`super.[父类函数名](参数列表)`的形式.
10. 调用函数和java一样使用: `[函数名](参数列表的形式)`,但是没有分号结尾.
11. kotlin的一条语句和java的一条语句的定义可能不一样,kotlin没有分号作为语句的结尾,一行就是一条语句.(这个待确定.TODO)
