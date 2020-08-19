# （1）系列教程介绍

  上面几个教程我们的程序都是生成可执行文件。但是我们在合作开发算法的时候经常需要交付的是一个模块，该模块提供特定的算法功能，用于给整体的项目进行调用。但我们又不能直接提供源码，所以我们可以提供一个库文件（静态库或者动态库），配置接口文件可以在不提供源代码的情况下给他人提供算法模块功能。本文主要讲述如何使用CMakeLists.txt，配置生成动态和静态库文件。

## （2）CMake 的使用环境和安装

本教程的使用环境为：

```
ubutu18.04 lts
gcc version 7.5.0
g++ version 7.5.0
cmake version 3.10.2
```

安装cmake：

```bash
sudo apt install cmake
```

## （3）设置设置我们的程序输出为lib文件

  配置输出库文件(lib和so)主要是通过函数add_subdirectory进行配置，使用STATIC表示静态库（lib），SHARED表示动态库（so）。同时可以使用SET_TARGET_PROPERTIES函数可以设置库文件的版本。

首先，看一下整体的目录结构：

```
|-- tutorial_fourth/
  |-- src/
      |-- tutorial.cpp
  |-- include/
      |--TutorialConfig.h.in
  |-- mathlib/
      |-- CMakeLists.txt
      |-- mathlib.h
      |-- mathlib.cpp
  |-- CMakeLists.txt
```

根目录下的CMakeLists.txt文件为：

```cmake
# 设置cmake的最低版本
cmake_minimum_required(VERSION 3.10)

# 设置工程名称 和版本
project(tutorial VERSION 1.0)

# 设置指定的C++编译器版本是必须的，如果不设置，或者为OFF，则指定版本不可用时，会使用上一版本。
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 指定为C++11 版本
set(CMAKE_CXX_STANDARD 11)

# 提供一个选项是OFF或者ON，如果没有初始值被提供则默认使用OFF
option(USE_MYMATH "Use tutorial provided math implementation" ON)

# 指定版本号的配置文件
configure_file(include/TutorialConfig.h.in TutorialConfig.h)

# 判断变量USE_MYMATH是否设置了ON，如果设置了配置mathlib library
if(USE_MYMATH)
  # 添加一个名字为mathlib的子编译路径
  add_subdirectory(mathlib)

  # 列出mathlib库的所有项目，并添加到外部库变量EXTRA_LIBS中
  list(APPEND EXTRA_LIBS mathlib)

  # 将子路径"${PROJECT_SOURCE_DIR}/mathlib"添加到外部路径变量EXTRA_INCLUDES中
  list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/mathlib")
endif()

# 增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial src/tutorial.cpp)

# 对目标的外部库进行链接操作
target_link_libraries(tutorial PUBLIC ${EXTRA_LIBS})

# 为指定项目添加 include 路径
target_include_directories(tutorial PUBLIC
                        "${PROJECT_BINARY_DIR}"
                        ${EXTRA_INCLUDES}
)
```

mathlib路径下CMakeLists.txt文件为：

```cmake
# 设置动态库的版本 为1.2
SET_TARGET_PROPERTIES(mathlib PROPERTIES VERSION 1.2 SOVERSION 1)

# 为库mathlib 添加源文件，该命令声明库文件
add_library(mathlib mathlib.cpp)
```

为了将cmake中的变量自动传递到程序中，在TutorialConfig.h.in中添加#cmakedefine USE_MYMATH命令用于在cmake中生成USE_MYMATH宏定义。
TutorialConfig.h.in中代码为：

```c
// the configured options and settings for Tutorial
#define Tutorial_VERSION_MAJOR @tutorial_second_VERSION_MAJOR@
#define Tutorial_VERSION_MINOR @tutorial_second_VERSION_MINOR@
#cmakedefine USE_MYMATH
```

在程序tutorial.cpp中添加处理命令：在cmake中定义USE_MYMATH的时候在程序中才会编译该命令块。具体请看CMakeDemo4代码。

```c
#ifdef USE_MYMATH
#  include "MathFunctions.h"
#endif
```

命令使用：

set_target_properties:设置目标的属性值

```cmake
set_target_properties(target1 target2 ...
      PROPERTIES prop1 value1
      prop2 value2 ...)
```

add_library: 为生成的库添加源文件，是库的名字，直接写名字即可，不要写lib，会自动加上前缀。STATIC表示静态库（lib），SHARED表示动态库（so）。

```cmake
add_library(<name> [STATIC | SHARED | MODULE]
            [EXCLUDE_FROM_ALL]
            source1 [source2 ...])
```

option: 提供一个选项是OFF或者ON，如果没有初始值被提供则默认使用OFF。

```cmake
option(<option_variable> "help string describing option"
      [initial value])
```

add_subdirectory: 添加一个编译的子路径，在子路径中查找CMakeLists.txt文件，并进行编译。

```cmake
add_subdirectory(source_dir [binary_dir]
            [EXCLUDE_FROM_ALL])
```

list:对项目中的所有值构建一个list，然后对变量进行操作。

```cmake
list(LENGTH <list> <output variable>)
list(GET <list> <element index> [<element index> ...]
<output variable>)
list(APPEND <list> [<element> ...])
list(FILTER <list> <INCLUDE|EXCLUDE> REGEX <regular_expression>)
list(FIND <list> <value> <output variable>)
list(INSERT <list> <element_index> <element> [<element> ...])
list(REMOVE_ITEM <list> <value> [<value> ...])
list(REMOVE_AT <list> <index> [<index> ...])
list(REMOVE_DUPLICATES <list>)
list(REVERSE <list>)
list(SORT <list>)
```

target_link_libraries:对目标进行链接操作，如果有需要链接的库文件的时候才进行操作，如果该项目没有库文件，则不需要该命令。该命令一般在add_executable后面。

```cmake
target_link_libraries(<target> ... <item>... ...)
```

## （4）使用CMake进行编译

CMake在生成文件的过程中会生成很多中间缓存文件，为了使项目更简洁，文件路径更清楚，一般会在项目的root目录下建立一个文件夹，用于存储CMake生成的中间文件。而一般使用的文件家名称为build或者release。下面是使用命令：

```bash
# 进入项目的root目录，本文为：tutorial_first
cd tutorial_first

# 创建存储缓存文件的文件夹，build
mkdir build

# 使用CMake命令生成makefile文件
cmake ..

# 使用make命令进行编译
cmake --build .
```