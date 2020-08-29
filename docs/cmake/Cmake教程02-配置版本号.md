## （1）系列教程介绍

  本文主要讲述如何使用CMakeLists.txt，配置程序的版本号。程序在发布的时候需要对用的版本号，同时为了保证程序的兼容性，往往需要在程序中判断当前程序的版本。所以如果在编译过程中将版本号自动的传入程序中，就可以使程序更智能。

## （2）CMake 的使用环境和安装

本教程的使用环境为：

```
ubutu18.04 lts
gcc version 7.5.0
g++ version 7.5.0
cmake version 3.10.2
```

安装cmake：

```
sudo apt install cmake
```

## （3）CMake配置版本号

在project命令中添加 VERSION 1.0 指令 就可以指定程序的版本。通过使用configure_file命令，解析TutorialConfig.h.in文件，将该版本号自动转换成宏定义额方式，传递到程序中。

```
# 设置cmake的最低版本
cmake_minimum_required(VERSION 3.10)

# 设置工程名称 和版本
project(tutorial_second VERSION 1.0)

# 指定版本号的配置文件
configure_file(include/TutorialConfig.h.in TutorialConfig.h)

# 增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial_second src/tutorial_second.cpp)

# 为指定项目添加 include 路径
target_include_directories(tutorial_second PUBLIC
                            "${PROJECT_BINARY_DIR}"
)
```

命令使用：

configure_file：

```
configure_file(<input> <output>
   [COPYONLY] [ESCAPE_QUOTES] [@ONLY]
   [NEWLINE_STYLE [UNIX|DOS|WIN32|LF|CRLF] ])
```

## （4）使用CMake进行编译

CMake在生成文件的过程中会生成很多中间缓存文件，为了是项目更简洁，文件路径更清楚，一般会在项目的root目录下建立一个文件夹，用于存储CMake生成的中间文件。而一般使用的文件家名称为build或者release。下面是使用命令：

```
# 进入项目的root目录，本文为：tutorial_first
cd tutorial_first

# 创建存储缓存文件的文件夹，build
mkdir build

# 使用CMake命令生成makefile文件
cmake ..

# 使用make命令进行编译
cmake --build .
```

## （5）TutorialConfig.h.in中代码

```
// the configured options and settings for Tutorial
#define Tutorial_VERSION_MAJOR @Tutorial_VERSION_MAJOR@
#define Tutorial_VERSION_MINOR @Tutorial_VERSION_MINOR@
```

## （6）cpp 代码

tutorial_second.cpp 程序代码：

```
#include <iostream>
#include "TutorialConfig.h"

int main(int argc, char const *argv[])
{
    std::cout <<"第一个cmake教程" << std::endl;
    if (argc < 2) {
    // report version
    std::cout << argv[0] << " Version " << Tutorial_VERSION_MAJOR << "."
              << Tutorial_VERSION_MINOR << std::endl;
    std::cout << "Usage: " << argv[0] << " number" << std::endl;

  }
    return 0;
}
```