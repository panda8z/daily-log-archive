
# 怎样书写Go代码？（[How to Write Go Code ](https://golang.google.cn/doc/code.html)）


- 代码的组织方式。
  - 使用 package 组织代码，中文叫做 包。
  - 一个仓库包含一个或多个 module。
    - 仓库是指类似github之类的仓库
    - module 是多个有联系的 go package 的集合。
- 第一个Go语言程序。
  - 新建一个 hello 文件夹。
  - 使用 `go mod init example.com/user/hello`命令创建一个 go.mod 文件。
  - 查看一下 go.mod 文件，使用 `cat go.mod`命令。
  - go源码文件的第一行有效代码必须是 `package name`, 这叫包声明。
- 在同一个 module 里倒入另一个 package
  - hello 文件夹下新建 reverse.go 文件
  - 在 main.go 中 使用 `import "example.com/user/hello/morestrings"` 语句导入这个文件就能使用这个包的方法了。