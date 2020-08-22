## （1）本教程介绍

  本教程主要是通过多个例子讲述CMake在构建系统工程过程中遇到的问题，以及使用方法。每一个例子都会提供一套完整的使用demo。

## （2）CMake 的介绍

  CMake是Cross Platform Make的缩写，CMake是一个跨平台的安装（编译）工具，可以用简单的语句来描述所有平台的安装(编译过程)。他并不能直接生成最终的应用程序，而是产生标准的建构档（如 Unix 的 Makefile 或 Windows Visual C++ 的 projects/workspaces）。CMake使用更高级的工程组织工具，可以避免直接编写底层的makefile文件，使得工程构建和编译变得更加方便快捷。

  CMake是通过CMakeLists.txt文件进行工程构建的，你的只需要编写CMakeLists.txt文件，对编译构成进行配置。同时一个路径下只能有一个CMakeLists.txt文件，因为CMake工具在进行工程构建的过程中，会自动在执行路径中进行CMakeLists.txt文件的查找。

## （3）CMake 的使用环境和安装

CMake使用

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

## （4）第一个CMake 工程

首先，看一下整体的目录结构：

```
|-- tutorial_first/

  |-- src/

  |-- CMakeLists.txt
```

构建一个最简单的CMake工程只需要3步就可以完成，如下面的命令：

```
#设置cmake的最低版本
cmake_minimum_required(VERSION 3.10)

#设置工程名称 和版本
project(tutorial_first VERSION 1.0)

#增加生成可执行文件，生成的程序名称为：tutorial_first
add_executable(tutorial_first src/tutorial_first.cpp)
```

命令使用：
cmake_minimum_required ：

```
      cmake_minimum_required(VERSION <min>[...<max>] [FATAL_ERROR])
```

project：

```
      project(<PROJECT-NAME>
              [VERSION <major>[.<minor>[.<patch>[.<tweak>]]]]
              [DESCRIPTION <project-description-string>]
              [HOMEPAGE_URL <url-string>]
              [LANGUAGES <language-name>...])
```

add_executable：

```
      add_executable(<name> [WIN32] [MACOSX_BUNDLE]
                  [EXCLUDE_FROM_ALL]
                  [source1] [source2 ...])
```

## （5）使用CMake进行编译

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

## （6）cpp 代码

  tutorial_first.cpp 程序代码：

```
#include <iostream>

int main(int argc, char const *argv[])
{
  std::cout << "cmake_demo_1教程" << std::endl;

  return 0;
}
```