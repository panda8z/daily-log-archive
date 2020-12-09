 |  |  | 
title: Darwin | xnu笔记001
date: 2019 | 12 | 06 11:24:49
tags:
 |  |  | 

## Darwin | xnu笔记

github：

[apple/darwin | xnu: The Darwin Kernel (mirror)](https://github.com/apple/darwin | xnu)

### 什么是 XNU

XNU 内核是，是 MacOS和 iOS 使用的达尔文操作系统的一部分。 XNU 是 **X is Not Unix** 的首字母缩写。
XNU 是一个混合内核，它结合了卡内基梅隆大学开发的 Mach 内核，FreeBSD 的组件以及用于编写称为 IOKit 的驱动程序的 C ++ API。
XNU 在 x86_64 上针对单处理器和多处理器配置运行。

### XNU 源码目录说明

| 文件夹           | 说明                                                                            |
| ---------------- | ------------------------------------------------------------------------------- |
| config           | 为支持的体系结构和平台配置导出的api的配置                                       |
| SETUP            | 用于配置内核，版本控制和kextsymbol管理的基本工具集。                            |
| EXTERNAL_HEADERS | 来自其他项目的标头，以避免在构建时产生依赖关系。 更新源时，应定期同步这些标头。 |
| libkern          | 用于处理驱动程序和kexts的C ++ IOKit库代码。                                     |
| libsa            | 用于启动的内核引导程序代码                                                      |
| libsyscall       | 用户空间程序的syscall库接口                                                     |
| libkdd           | 用户库的源，用于解析内核数据（如内核组块数据）。                                |
| makedefs         | 顶层规则和内核构建定义。                                                        |
| osfmk            | 基于Mach内核的子系统                                                            |
| pexpert          | 平台特定的代码，例如中断处理，原子等。                                          |
| security         | 强制访问检查策略接口和相关实施。                                                |
| bsd              | BSD子系统代码                                                                   |
| tools            | 一组用于测试，调试和配置内核的实用程序。                                        |

### 怎样编译 XNU ？

#### 编译 `DEVELOPMENT` 内核

xnu 的 make 系统可以基于 KERNEL_CONFIGS 和 ARCH_CONFIGS 变量作为参数来构建内核。 语法如下：

```bash
make SDKROOT=<sdkroot> ARCH_CONFIGS=<arch> KERNEL_CONFIGS=<variant>
```

位置说明：

* `<sdkroot>`：磁盘上macOS SDK的路径。 （默认为 `/` ）
* `<variant>`：可以是 `debug`，`development`，`release`，`profile` 并配置编译标志并在整个内核代码中声明。
* `<arch>`：需要构建的架构。 （例如 `X86_64` ）

