---
title: Android创建定时器的几种方式
date: 2018-09-03 11:23:20
updated: 
comments: true	
tags: 
- Android 
- 2018 
- 技术
categories:	
- Android	
permalink:
---
# Android中几种常用的定时器和延时方法

## 一、三种常用的定时器

1. Handler类的postDelayed方法：

```java
Handler mHandler = new Handler();
        Runnable r = new Runnable() {

        @Override
        public void run() {
                //do something
                //每隔1s循环执行run方法
                mHandler.postDelayed(this, 1000);
        }
    };
// 主线程中调用：  mHandler.postDelayed(r, 100);//延时100毫秒
```

2. 用handler+timer+timeTask方法：

```java
      Handler handler = new Handler() {
        @Override
        public void handleMessage(Message msg) {
            if (msg.what == 1){
                //do something
            }
            super.handleMessage(msg);
        }
    };

    Timer timer = new Timer();
    TimerTask timerTask = new TimerTask() {
        @Override
        public void run() {
            Message message = new Message();
            message.what = 1;
            handler.sendMessage(message);
        }
    };
 // 主线程中调用：timer.schedule(timerTask,1000,500);//延时1s，每隔500毫秒执行一次run方法
```

3. Thread+handler方法：

<!-- more -->

```java
 Handler handler = new Handler() {
        @Override
        public void handleMessage(Message msg) {
            if (msg.what == 1){
                //do something
            }
            super.handleMessage(msg);

        }
    };

    class MyThread extends Thread {//这里也可用Runnable接口实现
        @Override
        public void run() {
            while (true){
                try {
                    Thread.sleep(1000);//每隔1s执行一次
                    Message msg = new Message();
                    msg.what = 1;
                    handler.sendMessage(msg);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

            }
        }
    }
//主线程中调用：new Thread(new MyThread()).start();
```

二、三种延时的快捷方法：

1. Handler的postDelayed方法：

```java
new Handler().postDelayed(new Runnable() {

        @Override
        public void run() {
              //do something
          }
       }, 1000);    //延时1s执行
```

2. timer + TimerTask方法：

```java
    timer = new Timer();
    timer.schedule(new TimerTask() {                   
    @Override
    public void run() {
            //do something
    }
},1000);//延时1s执行
```

3. Thread方法：

```java
new Thread(new MyThread()).start();
new Thread(new Runnable() {
    @Override
    public void run() {
        try {
            Thread.sleep(1000);//延时1s
            //do something
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}).start();
```
