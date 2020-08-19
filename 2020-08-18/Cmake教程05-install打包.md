## （1）系列教程介绍

  我们编译生成的可执行文件一般，会生成在当前的编译路径下，也就是build或者release路径下。那么如何将编译生成的可执行文件和库文件打包到一起进行发布那？本片教程我们将讲述如何在cmake中配置install的打包路径。下面我们将以mathlib库和头文件为例子进行配置。

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

## （3）设置设置我们的程序输出为lib文件

  配置库文件、头文件和执行文件到install的目录下，cmake中的install根目录为CMAKE_INSTALL_PREFIX变量的路径，如果我们要设置配置路径可以使用set命令设置CMAKE_INSTALL_PREFIX变量的值来改变路径。一般默认情况CMAKE_INSTALL_PREFIX变量的值为,在UNIX系统中为：/usr/local，在windows系统中为：c:/Program Files/${PROJECT_NAME}

首先，看一下整体的目录结构

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

# 判断变量USE_MYMATH是否设置了ON，如果设置了配置MathFunctions library
if(USE_MYMATH)
  # 添加一个名字为MathFunctions的子编译路径
  add_subdirectory(mathlib)

  # 列出MathFunctions库的所有项目，并添加到外部库变量EXTRA_LIBS中
  list(APPEND EXTRA_LIBS mathlib)

  # 将子路径"${PROJECT_SOURCE_DIR}/mathlib"添加到外部路径变量EXTRA_INCLUDES中
  list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/mathlib")
endif()

# 增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial src/tutorial.cpp)

# 对目标的外部库进行链接操作，需要放在定义了tutorial以后
target_link_libraries(tutorial PUBLIC ${EXTRA_LIBS})

# 为指定项目添加 include 路径,需要放在定义了tutorial以后
target_include_directories(tutorial PUBLIC
                            "${PROJECT_BINARY_DIR}"
                            ${EXTRA_INCLUDES}
)

# 下面配置install，根目录为 CMAKE_INSTALL_PREFIX变量中的路径

# 配置可执行文件到安装路径 CMAKE_INSTALL_PREFIX的bin中
install(TARGETS tutorial DESTINATION bin)

# 配置程序的头文件到安装路径 CMAKE_INSTALL_PREFIX的include文件中
install(FILES "${PROJECT_BINARY_DIR}/TutorialConfig.h"
  DESTINATION include
  )
```

mathlib路径下CMakeLists.txt文件为：

```cmake
# 生成库文件名为mathlib的静态库
add_library(mathlib STATIC mysqrt.cpp) # 可以配置STATIC、SHARED和MODULE

# 设置动态库的版本 为1.2
SET_TARGET_PROPERTIES(mathlib PROPERTIES VERSION 1.2 SOVERSION 1)

# 将程序段额依赖库输出到安装路径 CMAKE_INSTALL_PREFIX的lib文件夹中
install(TARGETS mathlib DESTINATION lib)

# 将文件mathlib.h输出到安装目录 CMAKE_INSTALL_PREFIX下的include文件夹中
install(FILES mathlib.h DESTINATION include)
```

命令使用：
install: 配置程序打包过程中的目标（TARGETS）、文件（FILES）、路径（DIRECTORY）、代码（CODE）和输出配置（EXPORT）

```cmake
install(TARGETS <target>... [...])
install({FILES | PROGRAMS} <file>... [...])
install(DIRECTORY <dir>... [...])
install(SCRIPT <file> [...])
install(CODE <code> [...])
install(EXPORT <export-name> [...])
```

使用demo：

```cmake
install(TARGETS myExe mySharedLib myStaticLib
        RUNTIME DESTINATION bin
        LIBRARY DESTINATION lib
        ARCHIVE DESTINATION lib/static)
```

## （4）使用CMake进行编译

CMake在生成文件的过程中会生成很多中间缓存文件，为了使项目更简洁，文件路径更清楚，一般会在项目的root目录下建立一个文件夹，用于存储CMake生成的中间文件。而一般使用的文件家名称为build或者release。下面是使用命令：

```shell
# 进入项目的root目录，本文为：tutorial_first
cd tutorial

# 创建存储缓存文件的文件夹，build
mkdir build

# 使用CMake命令生成makefile文件
cmake ..

# 使用make命令进行编译
cmake --build .
```