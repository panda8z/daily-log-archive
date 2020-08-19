## （1）系列教程介绍

  本文主要讲述如何使用CMakeLists.txt，指定当前程序的C++编译标准。



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



## （3）CMake指定C++标准

  指定C++准备一般有两种方式。下面给出两种方式的例子。
  提示：在最新的CMake中推荐使用方法一。 



首先，看一下整体的目录结构：

```
|-- tutorial_third/

  |-- src/tutorial.cpp

  |-- include/TutorialConfig.h.in

  |-- CMakeLists.txt
```

1. 方案一

  使用标志位CMAKE_CXX_STANDARD_REQUIRED和CMAKE_CXX_STANDARD指定编译器的使用版本，如果CMAKE_CXX_STANDARD_REQUIRED设置为True，则必须使用CMAKE_CXX_STANDARD指定的版本，如果CMAKE_CXX_STANDARD_REQUIRED设置为OFF则CMAKE_CXX_STANDARD指定版本的为首选版本，如果没有会使用上一版本。



```
# 设置cmake的最低版本
cmake_minimum_required(VERSION 3.10)

# 设置工程名称 和版本
project(tutorial VERSION 1.0)

# 设置指定的C++编译器版本是必须的，如果不设置，或者为OFF，则指定版本不可用时，会使用上一版本。
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 指定为C++11 版本
set(CMAKE_CXX_STANDARD 11)

# 指定版本号的配置文件
configure_file(include/TutorialConfig.h.in TutorialConfig.h)

# # 指定为C++14 版本
# set(CMAKE_CXX_STANDARD 14)

# 增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial src/tutorial.cpp)

# 为指定项目添加 include 路径
target_include_directories(tutorial PUBLIC
                            "${PROJECT_BINARY_DIR}"
)
```



命令使用：

set：设置变量variable的值为value。

```
      set(<variable> <value>... [PARENT_SCOPE])
```



核心命令：

```
# 设置指定的C++编译器版本是必须的，如果不设置，或者为OFF，则指定版本不可用时，会使用上一版本。
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 指定为C++11 版本
set(CMAKE_CXX_STANDARD 11)
```



1. 方案二

  该方法直接指定CMAKE_CXX_FLAGS标志位进行设置，具体使用方式如下面的代码例子：



```
# 设置cmake的最低版本
cmake_minimum_required(VERSION 3.10)

# 设置工程名称 和版本
project(tutorial VERSION 1.0)

# 指定版本号的配置文件
configure_file(include/TutorialConfig.h.in TutorialConfig.h)

# 设置指定C++编译器版本。
include(CheckCXXCompilerFlag)
CHECK_CXX_COMPILER_FLAG("-std=c++11" COMPILER_SUPPORTS_CXX11)
CHECK_CXX_COMPILER_FLAG("-std=c++0x" COMPILER_SUPPORTS_CXX0X)
if(COMPILER_SUPPORTS_CXX11)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++14") # set C++ 11
# set(CMAKE_C_FLAGS  "${CMAKE_C_FLAGS} -std=c99")
elseif(COMPILER_SUPPORTS_CXX0X)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++0x")
message( STATUS "The comipler ${CMAKE_CXX_COMIPLER} has no C++ 11 suport. Please use a different C++ comipler.")
endif()

# 增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial src/tutorial.cpp)

# 为指定项目添加 include 路径
target_include_directories(tutorial PUBLIC
                            "${PROJECT_BINARY_DIR}"
)
```



## （4）使用CMake进行编译



CMake在生成文件的过程中会生成很多中间缓存文件，为了使项目更简洁，文件路径更清楚，一般会在项目的root目录下建立一个文件夹，用于存储CMake生成的中间文件。而一般使用的文件家名称为build或者release。下面是使用命令：



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



##  