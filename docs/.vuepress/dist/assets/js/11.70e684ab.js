(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{360:function(t,s,a){"use strict";a.r(s);var n=a(42),e=Object(n.a)({},(function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"（1）系列教程介绍"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#（1）系列教程介绍"}},[t._v("#")]),t._v(" （1）系列教程介绍")]),t._v(" "),a("p",[t._v("上面几个教程我们的程序都是生成可执行文件。但是我们在合作开发算法的时候经常需要交付的是一个模块，该模块提供特定的算法功能，用于给整体的项目进行调用。但我们又不能直接提供源码，所以我们可以提供一个库文件（静态库或者动态库），配置接口文件可以在不提供源代码的情况下给他人提供算法模块功能。本文主要讲述如何使用CMakeLists.txt，配置生成动态和静态库文件。")]),t._v(" "),a("h2",{attrs:{id:"（2）cmake-的使用环境和安装"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#（2）cmake-的使用环境和安装"}},[t._v("#")]),t._v(" （2）CMake 的使用环境和安装")]),t._v(" "),a("p",[t._v("本教程的使用环境为：")]),t._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("ubutu18.04 lts\ngcc version 7.5.0\ng++ version 7.5.0\ncmake version 3.10.2\n")])])]),a("p",[t._v("安装cmake：")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[t._v("sudo")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("apt")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("install")]),t._v(" cmake\n")])])]),a("h2",{attrs:{id:"（3）设置设置我们的程序输出为lib文件"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#（3）设置设置我们的程序输出为lib文件"}},[t._v("#")]),t._v(" （3）设置设置我们的程序输出为lib文件")]),t._v(" "),a("p",[t._v("配置输出库文件(lib和so)主要是通过函数add_subdirectory进行配置，使用STATIC表示静态库（lib），SHARED表示动态库（so）。同时可以使用SET_TARGET_PROPERTIES函数可以设置库文件的版本。")]),t._v(" "),a("p",[t._v("首先，看一下整体的目录结构：")]),t._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("|-- tutorial_fourth/\n  |-- src/\n      |-- tutorial.cpp\n  |-- include/\n      |--TutorialConfig.h.in\n  |-- mathlib/\n      |-- CMakeLists.txt\n      |-- mathlib.h\n      |-- mathlib.cpp\n  |-- CMakeLists.txt\n")])])]),a("p",[t._v("根目录下的CMakeLists.txt文件为：")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 设置cmake的最低版本")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("cmake_minimum_required")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("VERSION")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("3.10")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 设置工程名称 和版本")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("project")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("tutorial "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("VERSION")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1.0")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 设置指定的C++编译器版本是必须的，如果不设置，或者为OFF，则指定版本不可用时，会使用上一版本。")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("set")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token variable"}},[t._v("CMAKE_CXX_STANDARD_REQUIRED")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("ON")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 指定为C++11 版本")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("set")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token variable"}},[t._v("CMAKE_CXX_STANDARD")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("11")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 提供一个选项是OFF或者ON，如果没有初始值被提供则默认使用OFF")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("option")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("USE_MYMATH "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"Use tutorial provided math implementation"')]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("ON")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 指定版本号的配置文件")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("configure_file")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("include/TutorialConfig.h.in TutorialConfig.h"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 判断变量USE_MYMATH是否设置了ON，如果设置了配置mathlib library")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("USE_MYMATH"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 添加一个名字为mathlib的子编译路径")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("add_subdirectory")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("mathlib"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 列出mathlib库的所有项目，并添加到外部库变量EXTRA_LIBS中")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("APPEND EXTRA_LIBS mathlib"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v('# 将子路径"${PROJECT_SOURCE_DIR}/mathlib"添加到外部路径变量EXTRA_INCLUDES中')]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("APPEND EXTRA_INCLUDES "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"'),a("span",{pre:!0,attrs:{class:"token interpolation"}},[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("${")]),a("span",{pre:!0,attrs:{class:"token variable"}},[t._v("PROJECT_SOURCE_DIR")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")])]),t._v('/mathlib"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("endif")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 增加生成可执行文件，生成的程序名称为：tutorial_first")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("add_executable")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("tutorial src/tutorial.cpp"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 对目标的外部库进行链接操作")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("target_link_libraries")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("tutorial "),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("PUBLIC")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("${")]),t._v("EXTRA_LIBS"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 为指定项目添加 include 路径")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("target_include_directories")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("tutorial "),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("PUBLIC")]),t._v("\n                        "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"'),a("span",{pre:!0,attrs:{class:"token interpolation"}},[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("${")]),a("span",{pre:!0,attrs:{class:"token variable"}},[t._v("PROJECT_BINARY_DIR")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")])]),t._v('"')]),t._v("\n                        "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("${")]),t._v("EXTRA_INCLUDES"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("mathlib路径下CMakeLists.txt文件为：")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 设置动态库的版本 为1.2")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("SET_TARGET_PROPERTIES")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("mathlib "),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("PROPERTIES")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("VERSION")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1.2")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("SOVERSION")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 为库mathlib 添加源文件，该命令声明库文件")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("add_library")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("mathlib mathlib.cpp"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("为了将cmake中的变量自动传递到程序中，在TutorialConfig.h.in中添加#cmakedefine USE_MYMATH命令用于在cmake中生成USE_MYMATH宏定义。\nTutorialConfig.h.in中代码为：")]),t._v(" "),a("div",{staticClass:"language-c extra-class"},[a("pre",{pre:!0,attrs:{class:"language-c"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// the configured options and settings for Tutorial")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("define")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token expression"}},[t._v("Tutorial_VERSION_MAJOR @tutorial_second_VERSION_MAJOR@")])]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("define")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token expression"}},[t._v("Tutorial_VERSION_MINOR @tutorial_second_VERSION_MINOR@")])]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("cmakedefine")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token expression"}},[t._v("USE_MYMATH")])]),t._v("\n")])])]),a("p",[t._v("在程序tutorial.cpp中添加处理命令：在cmake中定义USE_MYMATH的时候在程序中才会编译该命令块。具体请看CMakeDemo4代码。")]),t._v(" "),a("div",{staticClass:"language-c extra-class"},[a("pre",{pre:!0,attrs:{class:"language-c"}},[a("code",[a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("ifdef")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token expression"}},[t._v("USE_MYMATH")])]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),t._v("  "),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("include")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"MathFunctions.h"')])]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token macro property"}},[a("span",{pre:!0,attrs:{class:"token directive-hash"}},[t._v("#")]),a("span",{pre:!0,attrs:{class:"token directive keyword"}},[t._v("endif")])]),t._v("\n")])])]),a("p",[t._v("命令使用：")]),t._v(" "),a("p",[t._v("set_target_properties:设置目标的属性值")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("set_target_properties")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("target1 target2 ...\n      "),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("PROPERTIES")]),t._v(" prop1 value1\n      prop2 value2 ..."),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("add_library: 为生成的库添加源文件，是库的名字，直接写名字即可，不要写lib，会自动加上前缀。STATIC表示静态库（lib），SHARED表示动态库（so）。")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("add_library")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("<name"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ["),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("STATIC")]),t._v(" | "),a("span",{pre:!0,attrs:{class:"token namespace"}},[t._v("SHARED")]),t._v(" | MODULE]\n            ["),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("EXCLUDE_FROM_ALL")]),t._v("]\n            source1 [source2 ...]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("option: 提供一个选项是OFF或者ON，如果没有初始值被提供则默认使用OFF。")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("option")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("<option_variable"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"help string describing option"')]),t._v("\n      [initial value]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("add_subdirectory: 添加一个编译的子路径，在子路径中查找CMakeLists.txt文件，并进行编译。")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("add_subdirectory")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("source_dir [binary_dir]\n            ["),a("span",{pre:!0,attrs:{class:"token property"}},[t._v("EXCLUDE_FROM_ALL")]),t._v("]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("list:对项目中的所有值构建一个list，然后对变量进行操作。")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("LENGTH <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <output variable"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("GET <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <element index"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" [<element index"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ...]\n<output variable"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("APPEND <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" [<element"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ...]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("FILTER <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <INCLUDE|EXCLUDE"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" REGEX <regular_expression"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("FIND <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <value"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <output variable"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("INSERT <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <element_index"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <element"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" [<element"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ...]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("REMOVE_ITEM <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <value"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" [<value"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ...]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("REMOVE_AT <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" <index"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" [<index"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ...]"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("REMOVE_DUPLICATES <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("REVERSE <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("list")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("SORT <list"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("target_link_libraries:对目标进行链接操作，如果有需要链接的库文件的时候才进行操作，如果该项目没有库文件，则不需要该命令。该命令一般在add_executable后面。")]),t._v(" "),a("div",{staticClass:"language-cmake extra-class"},[a("pre",{pre:!0,attrs:{class:"language-cmake"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("target_link_libraries")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("<target"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v(" ... <item"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(">")]),t._v("... ..."),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("h2",{attrs:{id:"（4）使用cmake进行编译"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#（4）使用cmake进行编译"}},[t._v("#")]),t._v(" （4）使用CMake进行编译")]),t._v(" "),a("p",[t._v("CMake在生成文件的过程中会生成很多中间缓存文件，为了使项目更简洁，文件路径更清楚，一般会在项目的root目录下建立一个文件夹，用于存储CMake生成的中间文件。而一般使用的文件家名称为build或者release。下面是使用命令：")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 进入项目的root目录，本文为：tutorial_first")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("cd")]),t._v(" tutorial_first\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 创建存储缓存文件的文件夹，build")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("mkdir")]),t._v(" build\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 使用CMake命令生成makefile文件")]),t._v("\ncmake "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("..")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 使用make命令进行编译")]),t._v("\ncmake --build "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v(".")]),t._v("\n")])])])])}),[],!1,null,null,null);s.default=e.exports}}]);