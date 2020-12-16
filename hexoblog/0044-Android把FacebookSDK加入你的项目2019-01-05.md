---
title: 把 Facebook SDK 加入你的项目2019-01-05
date: 2019-01-05 10:28:59
tags:
---

# 把 Facebook SDK 加入你的项目

1. 导入 SDK
2. 添加 SDK
3. 应用信息
4. 密钥散列
5. 应用事件
6. 完成

## 1. 导入SDK

如需在项目中使用 Facebook SDK，请把它添加为版本依赖关系并导入。如果你在创建新项目，请按以下所有步骤操作。如需把 Facebook SDK 加入现有项目，请从第 3 步开始操作。

1. 前往`Android Studio | New Project | Minimum SDK`
2. 选择`“API 15: Android 4.0.3”`或更高版本来创建新项目。
3. 在你的项目中，打开`your_app | Gradle Scripts | build.gradle` 
4. 把 Maven Central Repository 添加到dependencies:前的build.gradle

    ```java
        repositories {
                mavenCentral()
            }
    ```

5. 把compile 'com.facebook.android:facebook-android-sdk:[4,5)'加入你的build.gradle依赖关系。
6. 创建你的项目。
7. 把 Facebook SDK 导入你的应用：`import com.facebook.FacebookSdk;`

## 2. 添加 Facebook 应用编号

请为应用添加 Facebook 应用编号，并更新你的 Android 清单。

1. 打开你的strings.xml文件，例如：`/app/src/main/res/values/strings.xml`。

2. 添加名为facebook_app_id的新字串，字串名称应包含你的 Facebook 应用编号：

    ```xml
    <string name="facebook_app_id">223xxxxxxxxx989</string>
    ```

3. 打开AndroidManifest.xml。
4. 把uses-permission元素加入清单：`<uses-permission android:name="android.permission.INTERNET"/>`
5. 把一项meta-data元素添加到application元素：

    ```xml
    <application android:label="@string/app_name" ...>
        ...
        <meta-data android:name="com.facebook.sdk.ApplicationId" android:value="@string/facebook_app_id"/>
        ...
    </application>
    ```

## 3. 请介绍你的 Android 项目

1. 包名

    包名是你 Android 应用的独特标识符。如果用户还未安装应用，我们就会使用这个包名让他们从 Google Play 下载你的应用。你可以在Android Manifest中找到你的包名。

    `com.xxx.xxxx`
2. 默认动态类别名称

    这是处理深度链接的操作的完整类别名称。我们从 Facebook 应用深度链接到你的应用时会使用此信息。你还可以在你的Android Manifest中找到此类别名称

    `MainActivity`


## 4. 添加你的开发和发布密钥散列

为确保你的应用和 Facebook 之间互动的真实性，你需要向我们提供你的开发环境的 Android 密钥散列。如果你的应用已发布，你还需要提供你的发布密钥散列。

### 1.Android开发密钥散列

每个 Android 开发环境都会获得唯一的开发密钥散列。如需在 Mac 设备上生成开发密钥散列，请执行以下命令：

`keytool -exportcert -alias androiddebugkey -keystore ~/.android/debug.keystore | openssl sha1 -binary | openssl base64`

请在 Windows 中执行这一命令：

`keytool -exportcert -alias androiddebugkey -keystore %HOMEPATH%\.android\debug.keystore | openssl sha1 -binary | openssl base64`

此命令将为你的开发环境生成长为 28 个字符的专属密钥散列。

你需要为应用的每位工作人员提供开发环境的开发密钥散列。

### 2. Android发布密钥散列

如果你的应用已发布，你还应该添加发布密钥的散列。
Android 应用必须使用发布密钥作为数字签名才能上传到应用商店。
如需生成发布密钥散列，请在 Mac 或 Windows 中执行以下命令，替换你的发布密钥别名和 keystore 路径：

`keytool -exportcert -alias YOUR_RELEASE_KEY_ALIAS -keystore YOUR_RELEASE_KEY_PATH | openssl sha1 -binary | openssl base64`

这会生成由 28 个字符组成的字串，请把字串复制粘贴到下方空格中。另外，请查看[Android documentation](https://developer.android.com/studio/publish/app-signing?fbclid=IwAR0IT92ThMzXxZCfpkom6bjnlJ9DRmtytSIITQm5zq8JAca7WmGqdJ0_yP4)来了解如何为应用签名。

### 3. 密钥散列示例

xwKxxxxxxz2fzlzzz3Hln82I=xx

## 5. 应用事件

### 1. 追踪应用安装量和应用打开次数

通过应用事件，你可以衡量移动应用广告促成的安装量，创建高价值目标受众，查看用户人口统计等分析数据。部分事件是在你为应用添加并配置 Facebook SDK 之后记录的。这些事件包括安装、启用、和停用事件。详情请见[App Events Guide](https://developers.facebook.com/docs/app-events/getting-started-app-events-android)。
用户安装你的应用或与之互动时，你便可以在应用[Insights dashboard](https://www.facebook.com/analytics/)中看到相应数据。

## 6. 结束

至此, Facebook的接入已经完成.

恭喜！你已将 Facebook SDK 加入你的项目。你可以开始将应用集成到 Facebook 了。你希望下一步做什么？[Skip to Developer Dashboard](https://developers.facebook.com/apps)或[Documentation](https://developers.facebook.com/docs/)

#### 分享

添加分享对话框，让用户向好友分享你的内容。

#### 登录

添加 Facebook 登录功能，让用户快速、轻松地登录你的应用。

#### 广告

通过 Facebook 应用广告来吸引更多用户。

#### 创收

通过 Facebook 广告为移动应用创收