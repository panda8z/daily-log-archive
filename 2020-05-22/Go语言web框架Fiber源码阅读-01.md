

## app.go 实体定义

```go
// Map 专门用来返回JSON用的数据结构
// Map is a shortcut for map[string]interface{}, usefull for JSON returns
type Map map[string]interface{}

// 参数解析器 保存着 路径字段和参数字段
// paramsParser holds the path segments and param names
type parsedParams struct {
	segs   []paramSeg 
	params []string
}

// 参数字段 保存着参数的元数据
// paramsSeg holds the segment metadata
type paramSeg struct {
	Param      string
	Const      string
	IsParam    bool
	IsOptional bool
	IsLast     bool
}


// Layer 中文叫 页层。 这个结构体保存着每一个注册的 handler 的所有 元数据 信息。
// Layer is a struct that holds all metadata for each registered handler
type Layer struct {
	// Internal fields // 内部字段
	use    bool         // USE matches path prefixes // 使用 匹配的 路径前缀
	star   bool         // Path equals '*' or '/*' // 路径 是 '*' 或者 '/*'
	root   bool         // Path equals '/' // 路径是 '/'
	parsed parsedParams // parsed contains parsed params segments // 已经解析的 参数字段

	// External fields for ctx.Route() method // 外部字段, 给 ctx.Route() 方法准备的
	Path    string     // Registered route path // 注册的路由路径
	Method  string     // HTTP method // http 方法
	Params  []string   // Slice containing the params names // 保存参数名的 切片
	Handler func(*Ctx) // Ctx handler // Ctx 处理函数
}

// Fiber应用程序定义成App
// App denotes the Fiber application.
type App struct {
	// Layer stack //页层栈
	stack [][]*Layer
	// Fasthttp server // fasthttp 服务器
	server *fasthttp.Server
	mutex  sync.Mutex // 同步锁
	// App settings // Fiber应用配置的设置信息
	Settings *Settings
}

// Settings 是一个结构体，保存着服务器的配置。
// Settings holds is a struct holding the server settings
type Settings struct {
  
  // 这将允许多个Go进程在同一端口上侦听
	// This will spawn multiple Go processes listening on the same port
	Prefork bool // default: false // 默认false 。
  
  // 严格路由使能。当设置为true 的时候 "/foo"  和 "/foo/"  是不同的路由。
	// Enable strict routing. When enabled, the router treats "/foo" and "/foo/" as different.
	StrictRouting bool // default: false // 默认false 。
  
  // 大小写使能。 当设置为true 的时候 "/foo"  和 "/Foo"  是不同的路由。
	// Enable case sensitivity. When enabled, "/Foo" and "/foo" are different routes.
	CaseSensitive bool // default: false
  
  //  "Server: value" HTTP 头的使能
	// Enables the "Server: value" HTTP header.
	ServerHeader string // default: ""
  
  // handler值不可变的使能。
	// Enables handler values to be immutable even if you return from handler
	Immutable bool // default: false
	// Enable or disable ETag header generation, since both weak and strong etags are generated
	// using the same hashing method (CRC-32). Weak ETags are the default when enabled.
	// Optional. Default value false
	ETag bool
	// Max body size that the server accepts
	BodyLimit int // default: 4 * 1024 * 1024
	// Maximum number of concurrent connections.
	Concurrency int // default: 256 * 1024
	// Disable keep-alive connections, the server will close incoming connections after sending the first response to client
	DisableKeepalive bool // default: false
	// When set to true causes the default date header to be excluded from the response.
	DisableDefaultDate bool // default: false
	// When set to true, causes the default Content-Type header to be excluded from the Response.
	DisableDefaultContentType bool // default: false
	DisableHeaderNormalizing  bool // default: false
	// When set to true, it will not print out the fiber ASCII and "listening" on message
	DisableStartupMessage bool
	// Folder containing template files
	TemplateFolder string // default: ""
	// Template engine: html, amber, handlebars , mustache or pug
	TemplateEngine func(raw string, bind interface{}) (string, error) // default: nil
	// Extension for the template files
	TemplateExtension string // default: ""
	// The amount of time allowed to read the full request including body.
	ReadTimeout time.Duration // default: unlimited
	// The maximum duration before timing out writes of the response.
	WriteTimeout time.Duration // default: unlimited
	// The maximum amount of time to wait for the next request when keep-alive is enabled.
	IdleTimeout time.Duration // default: unlimited
}

// Group struct
type Group struct {
	prefix string
	app    *App
}

// Static struct
type Static struct {
	// Transparently compresses responses if set to true
	// This works differently than the github.com/gofiber/compression middleware
	// The server tries minimizing CPU usage by caching compressed files.
	// It adds ".fiber.gz" suffix to the original file name.
	// Optional. Default value false
	Compress bool
	// Enables byte range requests if set to true.
	// Optional. Default value false
	ByteRange bool
	// Enable directory browsing.
	// Optional. Default value false.
	Browse bool
	// Index file for serving a directory.
	// Optional. Default value "index.html".
	Index string
}
```

