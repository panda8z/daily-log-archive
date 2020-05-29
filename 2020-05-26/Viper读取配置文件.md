



# Viper——Go语言写的配置文件读取写入工具神器



## 1. 资料搜集

一篇中文的使用入门教程：[Go语言配置管理神器——Viper中文教程 | 李文周的博客](https://www.liwenzhou.com/posts/Go/viper_tutorial/)

Viper官方仓库：[spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)



> ⚠️ ：这里强烈建议阅读  官方仓库的 README.md 和 中文入门教程。
>
> 本文提及的多数核心概念将来自以上两个资料。
>
> 本文主要是代码实践



## 2. 新建viper-demo项目

首先我们要建立一个 Go 语言工程测试 **Viper** 的用法。

```bash
mkdir viper-demo
cd viper-demo 
go mod init github.com/xxx/viper-demo
go get github.com/spf13/viper
```

至此我们新建了一个 Go 语言工程 `viper-demo` 文件夹. 接下来以此工程为基础我们来测试 **Viper** 的所有用法。



## 3. 读取 yaml 文件

在 **viper-demo** 文件夹里 创建两个文件 ： 

- `server.yaml`
- `main.go`

内容分别如下：

#### server.yaml

```yaml
name:
  first: panda
  last: 8z
age : 99
hobbies:
  - Coding
  - Movie
  - Swimming
```

#### main.go

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){
  //viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
  viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
  //viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
  //viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
		)
  
  // 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
}
```

## 3. 写入配置文件

从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:

- `WriteConfig` - 将当前的`viper`配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
- `SafeWriteConfig` - 将当前的`viper`配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
- `WriteConfigAs` - 将当前的`viper`配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
- `SafeWriteConfigAs` - 将当前的`viper`配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。

根据经验，标记为`safe`的所有方法都不会覆盖任何文件，而是直接创建（如果不存在），而默认行为是创建或截断。

示例： 将 `main.go` 的内容全部改成下面的代码：

##### main.go



```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){
	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
	viper.WriteConfigAs("new-server.yaml") // 直接写入，有内容就覆盖，没有文件就新建
}
```

这段代码执行完，会有一个新的文件出现，文件名： `new-server.yaml` 内容如下：

```yaml
age: 99
hobbies:
- Coding
- Movie
- Swimming
name:
  first: panda
  last: 8z
```

有没有注意到这个内容被 按字母顺序自动排序了。这可能是 **Viper**自动做了这件事。



##### 探索一下上面的程序使用 `SafeWriteConfigAs` 方法会怎样

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){
	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	
	err := viper.SafeWriteConfigAs("new-server.yaml") // 因为该配置文件已经存在，所以会报错
	if err != nil {
		fmt.Println(err)
	}
}
```


控制台输出： 
```bash
map[first:panda last:8z] 99 panda [Coding Movie Swimming]
Config File "new-server.yaml" Already Exists
```

## 4. 建立默认值

一个好的配置系统应该支持默认值。键不需要默认值，但如果没有通过配置文件、环境变量、远程配置或命令行标志（`flag`）设置键，则默认值非常有用。

重点代码：

```go
viper.SetDefault("name", "dogger")
viper.SetDefault("age", "18")
viper.SetDefault("class", map[string]string{"class01": "01", "class02": "02"})
```

完整代码：

##### main.go

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){

	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.SetDefault("name", "dogger")
	viper.SetDefault("age", "18")
	viper.SetDefault("class", map[string]string{"class01": "01", "class02": "02"})
	
	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	
	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
	viper.WriteConfigAs("server-04.yaml")
}

```

`server.yaml` 如文章开头所示。

新的文件 `server-04.yaml`是新生成的，内容如下：

##### server-04.yaml

```yaml
age: 99
class:
  class01: "01"
  class02: "02"
hobbies:
- Coding
- Movie
- Swimming
name:
  first: panda
  last: 8z

```

##### 总结：

我们代码的顺序是：

1. 设置 配置文件搜索路径
2. 设置 配置文件名称
3. 设置 默认值
4. 读取文件
5. 打印一部分读取到的值。
6. 重新写入文件

**能明显看出读取后的文件覆盖了默认设置。**

**同时也在重新写入文件这一步看到了 `class` 被加入了新的配置。**

**这也验证了 Viper 对各种值的优先级处理。**



## 5. 优先级

**Viper** 会按照下面的优先级。每个项目的优先级都高于它下面的项目:

1. 显示调用`Set`设置值
2. 命令行参数（`flag`）
3. 环境变量
4. 配置文件
5. `key/value` 存储
6. 默认值

**重要：** 目前 **Viper** 配置的键（**Key**）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。



## 6. key/value 存储

上一小节我们看了 **Viper** 处理 配置参数的默认优先级。前文我们已经讲了配置文件和默认值两种方式。

接下来我们将 **key/value 存储**。乍一看可能不知道这是什么东西。

在 **Viper** 中启用远程支持，需要在代码中匿名导入`viper/remote`这个包。

```
import _ "github.com/spf13/viper/remote"
```

**Viper** 将读取从**Key/Value 存储**（例如 **etcd** 或 **Consul**）中的路径检索到的配置字符串（如`JSON`、`TOML`、`YAML`、`HCL`、`envfile` 和 `Java properties` 格式）。这些值的优先级高于默认值，但是会被从磁盘、`flag` 或环境变量检索到的配置值覆盖。（译注：也就是说 **Viper** 加载配置值的优先级为：磁盘上的配置文件 > 命令行参数 > 环境变量 > 远程 **Key/Value 存储 ** > 默认值。）

Viper使用 [crypt](https://github.com/bketelsen/crypt) 从 **K/V 存储** 中检索配置，这意味着如果你有正确的 **gpg 密匙**，你可以将配置值加密存储并自动解密。加密是可选的。

你可以将远程配置与本地配置结合使用，也可以独立使用。

  `crypt` 有一个命令行助手，你可以使用它将配置放入 **K/V 存储中**。`crypt` 默认使用在[http://127.0.0.1:4001](http://127.0.0.1:4001/)的etcd。

```bash
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

确认值已经设置：

```bash
$ crypt get -plaintext /config/hugo.json
```

有关如何设置加密值或如何使用Consul的示例，请参见 [crypt](https://github.com/bketelsen/crypt) 文档。

##### 总结：

对于分布式场景和微服务场景中可能会用到。方式方法几乎和我们前文测试的和从配置文件读取一样。

涉及到 etcd 或 Consul 等外部组件，这里不额外做代码测试。

*仅在下方搬运李文周先生博客里的代码片段*。

##### 远程Key/Value存储示例-未加密

###### etcd

```go
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

###### Consul

你需要 Consul Key/Value存储中设置一个Key保存包含所需配置的JSON值。例如，创建一个key`MY_CONSUL_KEY`将下面的值存入Consul key/value 存储：

```json
{
    "port": 8080,
    "hostname": "liwenzhou.com"
}
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") // 需要显示设置成json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // liwenzhou.com
```

###### Firestore

```go
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // 配置的格式: "json", "toml", "yaml", "yml"
err := viper.ReadRemoteConfig()
```

当然，你也可以使用`SecureRemoteProvider`。

###### 远程Key/Value存储示例-加密

```go
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

###### 监控etcd中的更改-未加密

```go
// 或者你可以创建一个新的viper实例
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

// 第一次从远程读取配置
err := runtime_viper.ReadRemoteConfig()

// 反序列化
runtime_viper.Unmarshal(&runtime_conf)

// 开启一个单独的goroutine一直监控远端的变更
go func(){
	for {
	    time.Sleep(time.Second * 5) // 每次请求后延迟一下

	    // 目前只测试了etcd支持
	    err := runtime_viper.WatchRemoteConfig()
	    if err != nil {
	        log.Errorf("unable to read remote config: %v", err)
	        continue
	    }

	    // 将新配置反序列化到我们运行时的配置结构体中。你还可以借助channel实现一个通知系统更改的信号
	    runtime_viper.Unmarshal(&runtime_conf)
	}
}()
```



## 7. 显示调用`Set`设置值

###### 覆盖设置

这些可能来自命令行标志，也可能来自你自己的应用程序逻辑。

```go
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
```

## 8. 命令行参数（`flag`）

Viper 具有绑定到标志的能力。具体来说，Viper支持[Cobra](https://github.com/spf13/cobra)库中使用的`Pflag`。

与`BindEnv`类似，该值不是在调用绑定方法时设置的，而是在访问该方法时设置的。这意味着你可以根据需要尽早进行绑定，即使在`init()`函数中也是如此。

对于单个标志，`BindPFlag()`方法提供此功能。

**例如：**

```go
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

###### 绑定一组现有的pflags

你还可以绑定一组现有的pflags （pflag.FlagSet）：

**举个例子：**

```go
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // 从viper而不是从pflag检索值
```

在 Viper 中使用 pflag 并不阻碍其他包中使用标准库中的 flag 包。pflag 包可以通过导入这些 flags 来处理flag包定义的flags。这是通过调用pflag包提供的便利函数`AddGoFlagSet()`来实现的。

**例如：**

```go
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	// 使用标准库 "flag" 包
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // 从 viper 检索值

	...
}
```

###### flag接口

如果你不使用`Pflag`，Viper 提供了两个Go接口来绑定其他 flag 系统。

`FlagValue`表示单个flag。这是一个关于如何实现这个接口的非常简单的例子：

```go
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }
```

一旦你的 flag 实现了这个接口，你可以很方便地告诉Viper绑定它：

```go
viper.BindFlagValue("my-flag-name", myFlag{})
```

`FlagValueSet`代表一组 flags 。这是一个关于如何实现这个接口的非常简单的例子:

```go
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
```

一旦你的flag set实现了这个接口，你就可以很方便地告诉Viper绑定它：

```go
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)
```

## 9. 环境变量

Viper完全支持环境变量。这使`Twelve-Factor App`开箱即用。有五种方法可以帮助与ENV协作:

- `AutomaticEnv()`
- `BindEnv(string...) : error`
- `SetEnvPrefix(string)`
- `SetEnvKeyReplacer(string...) *strings.Replacer`
- `AllowEmptyEnv(bool)`

*使用ENV变量时，务必要意识到Viper将ENV变量视为区分大小写。*

Viper提供了一种机制来确保ENV变量是惟一的。通过使用`SetEnvPrefix`，你可以告诉Viper在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用这个前缀。

`BindEnv`使用一个或两个参数。第一个参数是键名称，第二个是环境变量的名称。环境变量的名称区分大小写。如果没有提供ENV变量名，那么Viper将自动假设ENV变量与以下格式匹配：前缀+ “_” +键名全部大写。当你显式提供ENV变量名（第二个参数）时，它 **不会** 自动添加前缀。例如，如果第二个参数是“id”，Viper将查找环境变量“ID”。

在使用ENV变量时，需要注意的一件重要事情是，每次访问该值时都将读取它。Viper在调用`BindEnv`时不固定该值。

`AutomaticEnv`是一个强大的助手，尤其是与`SetEnvPrefix`结合使用时。调用时，Viper会在发出`viper.Get`请求时随时检查环境变量。它将应用以下规则。它将检查环境变量的名称是否与键匹配（如果设置了`EnvPrefix`）。

`SetEnvKeyReplacer`允许你使用`strings.Replacer`对象在一定程度上重写 Env 键。如果你希望在`Get()`调用中使用`-`或者其他什么符号，但是环境变量里使用`_`分隔符，那么这个功能是非常有用的。可以在`viper_test.go`中找到它的使用示例。

或者，你可以使用带有`NewWithOptions`工厂函数的`EnvKeyReplacer`。与`SetEnvKeyReplacer`不同，它接受`StringReplacer`接口，允许你编写自定义字符串替换逻辑。

默认情况下，空环境变量被认为是未设置的，并将返回到下一个配置源。若要将空环境变量视为已设置，请使用`AllowEmptyEnv`方法。

###### Env 示例：

```go
SetEnvPrefix("spf") // 将自动转为大写
BindEnv("id")

os.Setenv("SPF_ID", "13") // 通常是在应用程序之外完成的

id := Get("id") // 13
```

### 10.监控并重新读取配置文件

Viper支持在运行时实时读取配置文件的功能。

需要重新启动服务器以使配置生效的日子已经一去不复返了，viper驱动的应用程序可以在运行时读取配置文件的更新，而不会错过任何消息。

只需告诉viper实例watchConfig。可选地，你可以为Viper提供一个回调函数，以便在每次发生更改时运行。

**确保在调用`WatchConfig()`之前添加了所有的配置路径。**

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
	fmt.Println("Config file changed:", e.Name)
})
```