

# ory/kratos-01-快速入门

> 快速入门官方英文文档： [Quickstart | **ORY Kratos**](https://www.ory.sh/kratos/docs/quickstart/)

**ORY Kratos** 有几个正在经历改动的组件，因此从一开始就使所有问题都变得困难。

本入门指南将帮助你安装 **ORY Kratos** 和一些其他依赖项，以便你了解 **ORY Kratos** 的工作方式。

请注意，本指南不能代替研究文档。你必须了解核心概念和 API 才能有效地使用 **ORY Kratos**。这篇文章仅是入门指导。

## 用例：你想要登录并注册你的应用程序

本节为你提供了一些有关我们要实现的目标以及为此所需的工具的背景信息。你还将了解我们为本指南选择的网络设置。

本快速入门指南的前提是假设我们正在编写一个称为 **SecureApp** 的 NodeJS 应用程序。 这个程序没有花哨的东西仅有一些 **ExpressJS** 和一些使用 **Handlebars** 的 **HTML** 模板。 我们使用 **TypeScript**，但这只是因为它更具可读性，并不是因为我们在做任何与众不同的事情！

在 UI 层，你当然可以选择任何技术。 它可以与 **Swift**，**ReactJS** 或 **Angular**（客户端）以及 **PHP**，**Ruby**，**Python**，**Java**（服务器端）或一起使用，总之随便你怎么组合。我们之所以选择 **NodeJS** + **TypeScript**，是因为我们认为它最容易理解，并且因为**JavaScript** 和 **NodeJS** 被普遍理解并且易于安装。

我们真的不知道 **SecureApp** 有一天它会被扩展成什么样子。 但是我们确切知道它将具有某种类型的仪表板并且需要用户，因此我们需要：

- 登录
- 登出
- 注册
- 个人资料管理（“更新名字”，“更新头像...”）
- 凭据管理（“添加新的辅助邮箱”，“更改密码”，“ ...”）
- 帐户恢复（“密码重置”）
- 使用Google Authenticator进行两因素身份验证
- “使用Google登录”和“使用GitHub登录”

当然还有在仪表盘显示一些信息
- 比如用户信息：“`你好 {{firstName}} {{lastName}}，你的生日在{{Birthday}}！`”。 仅在用户登录时可见！


## 设置

你可能已经知道，**ORY Kratos** 仅限于 API。 它没有 UI 或 HTML 模板引擎。 我们将在 NodeJS **SecureApp** 中实现所有面向用户的 UI（仪表盘，登录，注册...）。

为了确保没有未经身份验证（登录）的人都无法访问仪表板，我们可以使用一小段代码（此处为 ExpressJS）来执行此操作：

```js
// Import the ORY Kratos SDK. SDKs are available for all popular programming
// languages!
//
// We will add examples for other programming languages here soon!
import { KratosPublicSDK } from '@oryd/kratos-client';

// You can use protect as a middleware for expressJS:
//
//   import express from 'express'
//   const app = express()
//   app.get("/dashboard", needsLogin, dashboard)
//
const needsLogin = (req, res, next) => {
  new KratosPublicSDK('https://public.ory-kratos')
    .whoami(req)
    .then(({ body }) => {
      req.user = { session: body };
      next();
    })
    .catch(() => {
      res.redirect('/login');
    });
};
```

> **ORY Kratos**不仅是一个API，它还使用Cookie，HTTP重定向，Anti-CSRF令牌等等，因此这些你都不用烦心了！

由于我们的 **SecureApp** 和 **ORY Kratos** 需要共享 cookie，为了使 Anti-CSRF 令牌和登录会话正常工作，我们将设置转发请求到 ORY Krato 的 Public API 的路径。
 如果发出了对 `https://my-secureap/.ory/kratos/public/self-service/browser/flows/login` 的 HTTP 请求，我们会将请求转发给 `https://public.ory-kratos/self-service/browser/flows/login` （像代理一样） 并将响应通过管道传递回初始 HTTP 请求：

```js

import express from 'express';
import request from 'request';
const app = express();

const pathPrefix = '/.ory/kratos/public';
app.use(pathPrefix + '/', (req, res) => {
  const url = 'https://public.ory-kratos' + req.url.replace(pathPrefix, '');

  // Uses the request library to forward the request to **ORY Kratos**
  req.pipe(request(url, { followRedirect: false })).pipe(res);
});

// ...
// app.get("/dashboard", needsLogin, dashboard)
// ...

```


**ORY Kratos** 不附带管理用户界面。 你必须自己实现，或选择 ORY Cloud 产品（待发行）。 
在本快速入门中，我们将使用 **ORY Kratos** CLI（命令行界面）与 **ORY Kratos** 的管理 API 进行交互。

快速入门还随附 **MailSlurper**，这是演示程序用来演示如何进行演示的模拟 SMTP 服务器。 电子邮件验证有效。

## 克隆**ORY Kratos**并运行docker

为了使该示例正常工作，你将需要在系统上安装 **Git** 和 **Docker**，以及 **Docker Compose**。 不需要其他依赖项。 在开始之前，请确保 **Docker** 具有足够的磁盘空间。

> Tips: 本教程使用 **Docker-Compose** 卷，这些卷据报道已用完磁盘空间。 使用 `docker system df` 检查剩余磁盘空间。 如果卷高于85％的阈值，请在开始之前修剪旧的 Docker 对象！

> Tips: 如果遇到构建错误（例如，网络超时），请确保网络运行正常并再次运行 `macke docker`。 如果问题仍然存在，请随时打开问题。

让我们克隆 **ORY Kratos** 并运行 docker-compose：
```bash
git clone https://github.com/ory/kratos.git
# or if you have git+ssh set up:
#  git clone git@github.com:ory/kratos.git
cd kratos
git checkout v0.3.0-alpha.1
make quickstart

# or if you don't have make installed:
docker pull oryd/kratos:latest-sqlite
docker pull oryd/kratos-selfservice-ui-node:latest
docker-compose -f quickstart.yml -f quickstart-standalone.yml up --build --force-recreate
```


这可能需要一两分钟。
一旦输出变慢并且日志表明系统运行正常，就可以开始运行了！
状况良好的系统将显示以下内容（消息的顺序可能相反）：

```bash
kratos_1                      | time="2020-01-20T14:52:13Z" level=info msg="Starting the admin httpd on: 0.0.0.0:4434"
kratos_1                      | time="2020-01-20T14:52:13Z" level=info msg="Starting the public httpd on: 0.0.0.0:4433"
```


> 获得完整功能的系统有两个重要因素：
> 1. 需要确保端口 `4455`、`4433`、`4434`和`4436`空闲。
> 2. 确保始终使用 `127.0.0.1` 作为主机名，切勿使用 `localhost` ！这很重要，因为浏览器将这两个视为单独的域，因此在设置和使用正确的 `cookie` 时会遇到问题。

你可能会注意到在此示例中未使用任何数据库。 **ORY Kratos** 支持 **SQLite**，**PostgreSQL**，**MySQL** 和 **CockroachDB** 作为数据库后端。
为了快速入门，我们安装了一个永久卷来存储 **SQLite** 数据库。
以后的指南将解释如何设置生产系统！

## 网络架构
该演示利用了多种服务 Docker 镜像：

1. **ORY Kratos**
1. **SecureApp**-用 **NodeJS** 编写的示例应用程序，用于实现登录，注册，注销，...和仪表板屏幕。
1. **ORY Kratos**可以通过其发送电子邮件的SMTP服务器。我们将使用 **MailHog**，这是具有简单 UI 的简约 SMTP 穿越服务器。

为了更好地理解所有内容的连接方式，让我们看一下网络配置。假设你至少对Docker（组合）网络的工作方式有一定的了解：
![Docker组织起来的网络架构](https://img-blog.csdnimg.cn/20200520235850143.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)

如前所述，我们正在将请求代理到 **ORY Kratos** 的 Public API。

我们这样做是因为这样所有请求都将来自同一主机名。这样可以避免 Cookie 常见的跨域问题。

## 用户执行注册，注销，登录
已经有足够的理论支持了，是时候让我们开始下一步了！ 

让我们首先尝试打开仪表板-转到 `http://127.0.0.1:4455/dashboard`。你可能会注意到，你最终将登录到端点：
![登录](https://img-blog.csdnimg.cn/20200520235929758.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)

**SecureApp** 的登录屏幕
查看网络堆栈，你会看到两个重定向发生：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200520235952660.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)
**SecureApp** 的网络 trace。

因为在一开始浏览器还没有有效的身份验证（登录）会话建立，所以才有了从 `http://127.0.0.1:4445/dashboard` 到 `http://127.0.0.1:4455/.ory/kratos/public/self-service/browser/flows/login` 的首次重定向。

重定向到 某个**ORY Kratos** API ，记录**浏览器应用程序**的身份 ID，这里就是 request 后面的参数 。 

**ORY Kratos** 会进行一些安全检查，并准备表单数据，然后将浏览器重定向到 `http://127.0.0.1:4445/auth/login`，并附加一个 `?request =`  查询参数。

然后，`/auth/login`（属于**SecureApp**）从 **ORY Kratos** 的 Admin API 中获取对呈现表单至关重要的数据：

```bash

$ curl http://127.0.0.1:4434/self-service/browser/flows/requests/login?request=<request-id>
{
    "id": "27aa98bc-a074-418f-96fa-8b8146050209",
    "expires_at": "2020-01-20T21:10:12.7365393Z",
    "issued_at": "2020-01-20T21:00:12.7365532Z",
    "request_url": "http://127.0.0.1:4455/self-service/browser/flows/login",
    "methods": {
        "password": {
            "method": "password",
            "config": {
                "action": "http://127.0.0.1:4455/.ory/kratos/public/auth/browser/methods/password/login?request=27aa98bc-a074-418f-96fa-8b8146050209",
                "method": "POST",
                "fields": [
                    {
                        "name": "csrf_token",
                        "type": "hidden",
                        "required": true,
                        "value": "Ii8iIEdnn12vVQ2vyz2YaHjmXMUK5eSQgw9pgENGxPjXi1PHC9gOG51x61o2GT9LGvC81ddvmNXYeLvlPxA04g=="
                    },
                    {
                        "name": "identifier",
                        "type": "text",
                        "required": true
                    },
                    {
                        "name": "password",
                        "type": "password",
                        "required": true
                    }
                ]
            }
        }
    }
}

```
然后将此数据呈现为 HTML 表单。 此流程也可用于**单页应用程序（SPA）**和 **Angular** 或 **ReactJS** 之类的框架。 有关特定流程（登录，注册，注销等）的更多详细信息，请转至[概念章节](https://www.ory.sh/kratos/docs/concepts/index)。

让我们继续下一个流程-注册！ 单击“创建帐户”，这将启动类似于我们刚刚使用的流程：
![](https://img-blog.csdnimg.cn/2020052100002237.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)
到目前为止，网络重定向看起来很熟悉：
![](https://img-blog.csdnimg.cn/20200521000040827.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)
如果我们尝试使用 `123456` 这样的密码进行注册，那么 **ORY Kratos** 的密码策略将会拒绝：
![在这里插入图片描述](https://img-blog.csdnimg.cn/2020052100005749.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)

错误消息直接来自**ORY Kratos** 的 API：

```bash
$ curl http://127.0.0.1:4434/self-service/browser/flows/requests/registration?request=<request-id>
{
    "id": "79349cbd-c785-476a-8db8-d0d71c5b003c",
    "expires_at": "2020-01-20T21:17:00.5077381Z",
    "issued_at": "2020-01-20T21:07:00.5077527Z",
    "request_url": "http://127.0.0.1:4455/self-service/browser/flows/registration",
    "methods": {
        "password": {
            "method": "password",
            "config": {
                "action": "http://127.0.0.1:4455/.ory/kratos/public/auth/browser/methods/password/registration?request=79349cbd-c785-476a-8db8-d0d71c5b003c",
                "method": "POST",
                "fields": [
                    {
                        "name": "csrf_token",
                        "type": "hidden",
                        "required": true,
                        "value": "+ZQ8x5cVgdtt4xtPIRJXQPKMVU5c/S2Mj2MuudP32vsMME0g26oQnV/H/brcNvBjkJq1XoF3UcnUFPzcr6Eq4Q=="
                    },
                    {
                        "name": "password",
                        "type": "password",
                        "required": true
                    },
                    {
                        "name": "traits.email",
                        "type": "text",
                        "value": "hello@ory.sh"
                    },
                    {
                        "name": "traits.full_name",
                        "type": "text"
                    }
                ]
            }
        }
    }
}
```

设置不违反这些策略的密码，我们将立即重定向到仪表板：

![](https://img-blog.csdnimg.cn/2020052100012544.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)

通过使用“注销”，你将再次被重定向到登录屏幕，并且将能够使用刚刚设置的登录凭据！

## 了解登录和注册的工作方式
请转到[自助服务流程一章](https://www.ory.sh/kratos/docs/self-service/flows/index.md)，深入了解各个流程的工作原理。

## 电子邮件验证
在你注册后，一封电子邮件已发送到你使用的电子邮件地址。 因为快速入门使用了伪造的 SMTP 服务器，所以电子邮件没有到达你的收件箱。 但是，你可以通过在 [`127.0.0.1:4436`](127.0.0.1:4436) 打开 **MailSlurper UI** 来检索电子邮件。

你应该会看到以下内容：
![](https://img-blog.csdnimg.cn/20200521000150292.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)
如果不是，请硬刷新选项卡或单击菜单栏中的主页图标。

接下来，点击验证链接。

你最终将在仪表板中获得一个经过验证的电子邮件地址（请检查 JSON Payload 中的 `authenticated` 和 `authenticated_at` 字段）：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200521000213753.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3BhbmRhXzg=,size_16,color_FFFFFF,t_70)

要重新请求验证电子邮件，请填写表格 `127.0.0.1:4455/verify`。

## 使用的配置
你可以在  `./contrib/quickstart` 和  `./quickstart.yml` 中找到此快速入门指南的所有配置文件。 要了解每个单独的配置文件的功能，必须查阅本文档的其他章节。

> 要运行 **ORY Kratos** 的最低版本，你需要设置配置项 `identity.traits.default_schema_url` 和 `dsn`。 你还应该配置 `urls.*_ui`，因为否则你的用户将以回退状态结束。

将来，本指南将支持更多用例，例如：
- 使用 GitHub 登录并注册

## 清理Docker
最后，要清理所有内容，需要关闭 Docker Compose 环境并删除所有已安装的卷。

```bash
docker-compose -f quickstart.yml down -v
docker-compose -f quickstart.yml rm -f -s -v
```