# Go语言基础-上-Go命令行

> 关于命令行命令的使用这里有完整的官方文档：[go - The Go Programming Language](https://golang.google.cn/cmd/go/#hdr-Initialize_new_module_in_current_directory)
>
> 此文仅做简单 释义 和 翻译。

## 1. `go` 命令


#### 1.1  go 命令详情

**`go`是一个管理Go语言源码的工具**

在命令行 执行命令 `go` 打印如下信息：

> Tips: 保留原始信息的基础上在每一行英文的上面是相应的翻译。

```bash
$ go  
// go 是一个管理go源代码的工具🔧。
Go is a tool for managing Go source code.

// 使用方式
Usage:

        go <command> [arguments]

// 各 命令 如下
The commands are:
		
        // 开始一个bug报告。
        bug         start a bug report 
        
        // 编译 包 和 依赖。
        build       compile packages and dependencies 
        
        // 清理 对象文件 和 缓存文件。 
        clean       remove object files and cached files  
        
        // 显示 包或标识的 文档。
        doc         show documentation for package or symbol
        
        // 打印 Go 环境信息。
        env         print Go environment information 
        
        // 使用心得 APIs 更新所有的或指定的包。
        fix         update packages to use new APIs 
        
        // 格式化包的源代码。
        fmt         gofmt (reformat) package sources  
        
        // 通过处理源代码产生 go 文件。
        generate    generate Go files by processing source 
        
        // 给当前的模块添加依赖并且安装他们。
        get         add dependencies to current module and install them 
        
        // 编译 并且 安装 包s 及其 依赖。
        install     compile and install packages and dependencies 
        
        // 列出 包s 或 模块s。
        list        list packages or modules 

        // 用于维护 go 模块。
        mod         module maintenance 
        // 编译 并 运行 go 程序。
        run         compile and run Go program 

        // 测试 go 包s。
        test        test packages 

        // 运行指定的 go 工具🔧。
        tool        run specified go tool 

        // 打印 go 版本信息
        version     print Go version 
        
        // 在包中报告可能的错误
        vet         report likely mistakes in packages 

// 使用 “go help 【command】” 可以获得指定 命令 更多的信息。
Use "go help <command>" for more information about a command.

// 一些附加的 帮助主题：
Additional help topics:

        buildmode   build modes
        c           calling between Go and C
        cache       build and test caching
        environment environment variables
        filetype    file types
        go.mod      the go.mod file
        gopath      GOPATH environment variable
        gopath-get  legacy GOPATH go get
        goproxy     module proxy protocol
        importpath  import path syntax
        modules     modules, module versions, and more
        module-get  module-aware go get
        module-auth module authentication using go.sum
        module-private module configuration for non-public modules
        packages    package lists and patterns
        testflag    testing flags
        testfunc    testing functions

Use "go help <topic>" for more information about that topic.

```

使用命令 `go help 【主题名】` 可以获得其他帮助信息。

使用命令 `go help [命令名]` 可以查看命令的详细使用信息

#### 1.2 go 命令行使用格式

`go  [命令] [参数]` 

#### 1.3 命令解析


| 命令     | en                                                  | zh-cn                            |
| -------- | --------------------------------------------------- | -------------------------------- |
| bug      | start a bug report                                  | 开始一个bug报告                  |
| build    | compile packages and dependencies                   | 编译`包`及其依赖                 |
| clean    | remove object files and cached files                | 清理`对象文件`和`缓存文件`       |
| doc      | show documentation for package or symbol            | 显示`包`或`指定标志`的`文档`     |
| env      | print Go environment information                    | 打印当前设备 Go 的`环境变量`信息 |
| fix      | update packages to use new APIs                     | 将`包`更新到最新版本代码的API    |
| fmt      | gofmt (reformat) package sources                    | `gofmt` 包的源码                 |
| generate | generate Go files by processing source              | 通过`过程源码`产生 Go 文件       |
| get      | add dependencies to current module and install them | 向当前 module 添加依赖并安装它们 |
| install  | compile and install packages and dependencies       | 编译并安装 包 及其依赖           |
| list     | list packages or modules                            | 列出包或 module                  |
| mod      | module maintenance                                  | module 维护                      |
| run      | compile and run Go program                          | 编译病运行 Go 程序               |
| test     | test packages                                       | 测试包                           |
| tool     | run specified go tool                               | 运行指定 go 工具🛠                |
| version  | print Go version                                    | 打印 Go 的版本                   |
| vet      | report likely mistakes in packages                  | 报告 包 中可能的错误❌            |



## 2. 命令的详细使用说明



### 2.1 bug - bug 报告

```bash
go help bug
usage: go bug

Bug opens the default browser and starts a new bug report.
The report includes useful system information.
```

实测是给 Go语言官方提bug的快捷方式 直接打开github：[New Issue · golang/go](https://github.com/golang/go/issues/new?body=%3C%21--+Please+answer+these+questions+before+submitting+your+issue.+Thanks%21+--%3E%0A%0A%23%23%23+What+version+of+Go+are+you+using+%28%60go+version%60%29%3F%0A%0A%3Cpre%3E%0A%24+go+version%0Ago+version+go1.14.1+darwin%2Famd64%0A%3C%2Fpre%3E%0A%0A%23%23%23+Does+this+issue+reproduce+with+the+latest+release%3F%0A%0A%0A%23%23%23+What+operating+system+and+processor+architecture+are+you+using+%28%60go+env%60%29%3F%0A%0A%3Cdetails%3E%3Csummary%3E%3Ccode%3Ego+env%3C%2Fcode%3E+Output%3C%2Fsummary%3E%3Cbr%3E%3Cpre%3E%0A%24+go+env%0AGO111MODULE%3D%22on%22%0AGOARCH%3D%22amd64%22%0AGOBIN%3D%22%22%0AGOCACHE%3D%22%2FUsers%2Fzcj%2FLibrary%2FCaches%2Fgo-build%22%0AGOENV%3D%22%2FUsers%2Fzcj%2FLibrary%2FApplication+Support%2Fgo%2Fenv%22%0AGOEXE%3D%22%22%0AGOFLAGS%3D%22%22%0AGOHOSTARCH%3D%22amd64%22%0AGOHOSTOS%3D%22darwin%22%0AGOINSECURE%3D%22%22%0AGONOPROXY%3D%22%22%0AGONOSUMDB%3D%22%22%0AGOOS%3D%22darwin%22%0AGOPATH%3D%22%2FUsers%2Fzcj%2Fgopath%22%0AGOPRIVATE%3D%22%22%0AGOPROXY%3D%22https%3A%2F%2Fgoproxy.io%22%0AGOROOT%3D%22%2Fusr%2Flocal%2Fgo%22%0AGOSUMDB%3D%22sum.golang.org%22%0AGOTMPDIR%3D%22%22%0AGOTOOLDIR%3D%22%2Fusr%2Flocal%2Fgo%2Fpkg%2Ftool%2Fdarwin_amd64%22%0AGCCGO%3D%22gccgo%22%0AAR%3D%22ar%22%0ACC%3D%22clang%22%0ACXX%3D%22clang%2B%2B%22%0ACGO_ENABLED%3D%221%22%0AGOMOD%3D%22%2Fdev%2Fnull%22%0ACGO_CFLAGS%3D%22-g+-O2%22%0ACGO_CPPFLAGS%3D%22%22%0ACGO_CXXFLAGS%3D%22-g+-O2%22%0ACGO_FFLAGS%3D%22-g+-O2%22%0ACGO_LDFLAGS%3D%22-g+-O2%22%0APKG_CONFIG%3D%22pkg-config%22%0AGOGCCFLAGS%3D%22-fPIC+-m64+-pthread+-fno-caret-diagnostics+-Qunused-arguments+-fmessage-length%3D0+-fdebug-prefix-map%3D%2Fvar%2Ffolders%2F5q%2F_qqjn0l95rb46swlwq9fbxy00000gn%2FT%2Fgo-build554830363%3D%2Ftmp%2Fgo-build+-gno-record-gcc-switches+-fno-common%22%0AGOROOT%2Fbin%2Fgo+version%3A+go+version+go1.14.1+darwin%2Famd64%0AGOROOT%2Fbin%2Fgo+tool+compile+-V%3A+compile+version+go1.14.1%0Auname+-v%3A+Darwin+Kernel+Version+19.4.0%3A+Wed+Mar++4+22%3A28%3A40+PST+2020%3B+root%3Axnu-6153.101.6~15%2FRELEASE_X86_64%0AProductName%3A%09Mac+OS+X%0AProductVersion%3A%0910.15.4%0ABuildVersion%3A%0919E287%0Alldb+--version%3A+lldb-1103.0.22.8%0AApple+Swift+version+5.2.2+%28swiftlang-1103.0.32.6+clang-1103.0.32.51%29%0A%3C%2Fpre%3E%3C%2Fdetails%3E%0A%0A%23%23%23+What+did+you+do%3F%0A%0A%3C%21--%0AIf+possible%2C+provide+a+recipe+for+reproducing+the+error.%0AA+complete+runnable+program+is+good.%0AA+link+on+play.golang.org+is+best.%0A--%3E%0A%0A%0A%0A%23%23%23+What+did+you+expect+to+see%3F%0A%0A%0A%0A%23%23%23+What+did+you+see+instead%3F%0A%0A)

### 2.2 build - 编译包及其依赖

- 很多flag 都是通用的。*The build flags are shared by the `build`, `clean`, `get`, `install`, `list`, `run`,
  and `test` commands*
- `-race` *enable data race detection.Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64, linux/ppc64le and linux/arm64 (only for 48-bit VMA).* 在几个特定的平台上能开启数据竞争检测。



```bash
$go help build
usage: go build [-o output] [-i] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.

If the arguments to build are a list of .go files from a single directory,
build treats them as a list of source files specifying a single package.

When compiling packages, build ignores files that end in '_test.go'.

When compiling a single main package, build writes
the resulting executable to an output file named after
the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
The '.exe' suffix is added when writing a Windows executable.

When compiling multiple packages or a single non-main package,
build compiles the packages but discards the resulting object,
serving only as a check that the packages can be built.

The -o flag forces build to write the resulting executable or object
to the named output file or directory, instead of the default behavior described
in the last two paragraphs. If the named output is a directory that exists,
then any resulting executables will be written to that directory.

The -i flag installs the packages that are dependencies of the target.

The build flags are shared by the build, clean, get, install, list, run,
and test commands:

        -a
                force rebuilding of packages that are already up-to-date.
        -n
                print the commands but do not run them.
        -p n
                the number of programs, such as build commands or
                test binaries, that can be run in parallel.
                The default is the number of CPUs available.
        -race
                enable data race detection.
                Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64,
                linux/ppc64le and linux/arm64 (only for 48-bit VMA).
        -msan
                enable interoperation with memory sanitizer.
                Supported only on linux/amd64, linux/arm64
                and only with Clang/LLVM as the host C compiler.
                On linux/arm64, pie build mode will be used.
        -v
                print the names of packages as they are compiled.
        -work
                print the name of the temporary work directory and
                do not delete it when exiting.
        -x
                print the commands.

        -asmflags '[pattern=]arg list'
                arguments to pass on each go tool asm invocation.
        -buildmode mode
                build mode to use. See 'go help buildmode' for more.
        -compiler name
                name of compiler to use, as in runtime.Compiler (gccgo or gc).
        -gccgoflags '[pattern=]arg list'
                arguments to pass on each gccgo compiler/linker invocation.
        -gcflags '[pattern=]arg list'
                arguments to pass on each go tool compile invocation.
        -installsuffix suffix
                a suffix to use in the name of the package installation directory,
                in order to keep output separate from default builds.
                If using the -race flag, the install suffix is automatically set to race
                or, if set explicitly, has _race appended to it. Likewise for the -msan
                flag. Using a -buildmode option that requires non-default compile flags
                has a similar effect.
        -ldflags '[pattern=]arg list'
                arguments to pass on each go tool link invocation.
        -linkshared
                build code that will be linked against shared libraries previously
                created with -buildmode=shared.
        -mod mode
                module download mode to use: readonly, vendor, or mod.
                See 'go help modules' for more.
        -modcacherw
                leave newly-created directories in the module cache read-write
                instead of making them read-only.
        -modfile file
                in module aware mode, read (and possibly write) an alternate go.mod
                file instead of the one in the module root directory. A file named
                "go.mod" must still be present in order to determine the module root
                directory, but it is not accessed. When -modfile is specified, an
                alternate go.sum file is also used: its path is derived from the
                -modfile flag by trimming the ".mod" extension and appending ".sum".
        -pkgdir dir
                install and load all packages from dir instead of the usual locations.
                For example, when building with a non-standard configuration,
                use -pkgdir to keep generated packages in a separate location.
        -tags tag,list
                a comma-separated list of build tags to consider satisfied during the
                build. For more information about build tags, see the description of
                build constraints in the documentation for the go/build package.
                (Earlier versions of Go used a space-separated list, and that form
                is deprecated but still recognized.)
        -trimpath
                remove all file system paths from the resulting executable.
                Instead of absolute file system paths, the recorded file names
                will begin with either "go" (for the standard library),
                or a module path@version (when using modules),
                or a plain import path (when using GOPATH).
        -toolexec 'cmd args'
                a program to use to invoke toolchain programs like vet and asm.
                For example, instead of running asm, the go command will run
                'cmd args /path/to/asm <arguments for asm>'.

The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a
space-separated list of arguments to pass to an underlying tool
during the build. To embed spaces in an element in the list, surround
it with either single or double quotes. The argument list may be
preceded by a package pattern and an equal sign, which restricts
the use of that argument list to the building of packages matching
that pattern (see 'go help packages' for a description of package
patterns). Without a pattern, the argument list applies only to the
packages named on the command line. The flags may be repeated
with different patterns in order to specify different arguments for
different sets of packages. If a package matches patterns given in
multiple flags, the latest match on the command line wins.
For example, 'go build -gcflags=-S fmt' prints the disassembly
only for package fmt, while 'go build -gcflags=all=-S fmt'
prints the disassembly for fmt and all its dependencies.

For more about specifying packages, see 'go help packages'.
For more about where packages and binaries are installed,
run 'go help gopath'.
For more about calling between Go and C/C++, run 'go help c'.

Note: Build adheres to certain conventions such as those described
by 'go help gopath'. Not all projects can follow these conventions,
however. Installations that have their own conventions or that use
a separate software build system may choose to use lower-level
invocations such as 'go tool compile' and 'go tool link' to avoid
some of the overheads and design decisions of the build tool.

See also: go install, go get, go clean.
```



### 2.3 clean - 清理包和依赖



```bash
$go help clean
usage: go clean [clean flags] [build flags] [packages]

Clean removes object files from package source directories.
The go command builds most objects in a temporary directory,
so go clean is mainly concerned with object files left by other
tools or by manual invocations of go build.

If a package argument is given or the -i or -r flag is set,
clean removes the following files from each of the
source directories corresponding to the import paths:

        _obj/            old object directory, left from Makefiles
        _test/           old test directory, left from Makefiles
        _testmain.go     old gotest file, left from Makefiles
        test.out         old test log, left from Makefiles
        build.out        old test log, left from Makefiles
        *.[568ao]        object files, left from Makefiles

        DIR(.exe)        from go build
        DIR.test(.exe)   from go test -c
        MAINFILE(.exe)   from go build MAINFILE.go
        *.so             from SWIG

In the list, DIR represents the final path element of the
directory, and MAINFILE is the base name of any Go source
file in the directory that is not included when building
the package.

The -i flag causes clean to remove the corresponding installed
archive or binary (what 'go install' would create).

The -n flag causes clean to print the remove commands it would execute,
but not run them.

The -r flag causes clean to be applied recursively to all the
dependencies of the packages named by the import paths.

The -x flag causes clean to print remove commands as it executes them.

The -cache flag causes clean to remove the entire go build cache.

The -testcache flag causes clean to expire all test results in the
go build cache.

The -modcache flag causes clean to remove the entire module
download cache, including unpacked source code of versioned
dependencies.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
```



其余帮助信息用相同方式也能轻松获取，这里不再赘述