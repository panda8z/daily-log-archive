---
title: android开发日志-data&file
date: 2019-10-16 14:00:30
tags: [ Android, Android基础, 开发日志 ]
categories: [ Android基础 ]
---



> 今天复习下Android的数据和文件的一些操作。

## 一、资料

官网文档：[App data and files  |  Android Developers](https://developer.android.com/guide/topics/data)

android官方的一个codelab：[Keep Sensitive Data Safe and Private](https://codelabs.developers.google.com/codelabs/android-storage-permissions/#0)

codelab的代码的github仓库：[googlecodelabs/android-storage-permissions: This codelab will cover the correct way to store data securely in an Android app, how to access data on the device securely and how to limit the amount of data that apps expose.](https://github.com/googlecodelabs/android-storage-permissions)





## 二、记录

![image-20191016140247241](https://tva1.sinaimg.cn/large/006y8mN6gy1g8007va2ctj31ik0u0dzr.jpg)



![image-20191016145609980](https://tva1.sinaimg.cn/large/006y8mN6gy1g801rd0a1rj324o0n0guq.jpg)





![image-20191016145515354](https://tva1.sinaimg.cn/large/006y8mN6gy1g801qfc6irj30st0kodj3.jpg)



```
## General
Request URL: https://dl.google.com/dl/android/maven2/master-index.xml
Request Method: GET
Status Code: 200 
Remote Address: 58.152.43.226:3389
Referrer Policy: no-referrer-when-downgrade

#Response Headers
accept-ranges: bytes
access-control-allow-origin: https://maven.google.com
access-control-expose-headers: Content-Disposition, Content-Length, Content-Range, Content-Security-Policy, Date, Etag
age: 0
alt-svc: quic=":443"; ma=2592000; v="46,43",h3-Q048=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000
cache-control: public,max-age=86400
content-length: 3826
content-type: application/xml
date: Wed, 16 Oct 2019 05:12:32 GMT
etag: "4787bd"
last-modified: Wed, 09 Oct 2019 17:00:00 GMT
server: downloads
status: 200
vary: Origin
x-content-type-options: nosniff
x-frame-options: SAMEORIGIN
x-xss-protection: 0

## request headers


Provisional headers are shown
Accept: application/json, text/plain, */*
Origin: https://maven.google.com
Referer: https://maven.google.com/web/index.html
Sec-Fetch-Mode: cors
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36
```

