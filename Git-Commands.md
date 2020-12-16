# Panda8z的Git踩坑指南

> 记录一些git常用的操作

## 将本地分支推送到某个remote的某一个分支

使用 `git init` 和 `git remote add origin xxxx` 后, 直接执行 `git push` 会有以下的
信息提示。

```
fatal: The current branch master has no upstream branch.
To push the current branch and set the remote as upstream, use

    git push --set-upstream origin master

```

这个命令拆分如下：

- `git push` ： 推送到远端命令
- `--set-upstream`: 设置推送的流
- `origin`: 名字叫 `origin`的remote仓库
- `master`: 远程仓库的`master`分支