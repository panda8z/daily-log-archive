---
marp: true
---

# git 修改远程分支名称

首先 `git branch -m 旧分支名 新分支名`

其次 `git push --delete origin 旧分支名`

将新分支名推上去 `git push origin 新分支名`

将新本地分支和远程相连 `git branch --set-upsteam-to origin/新分支名`

> 2020-12-09 vol: 001
---
# Install rust on Mac, Linux or Windows

Just follow the doc: [Install Rust - Rust Programming Language](https://www.rust-lang.org/tools/install)

---

# Cargo 是个好工具

Cargo 是 rust 的编译系统 和 包管理器，rust 的好伙伴。
主要功能：
1. 创建 rust 项目。
2. 安装第三方库。
3. 编译，测试 rust 程序。

命令示例：
```bash
cargo --version
cargo new hello
cargo build
cargo run
cargo check
```