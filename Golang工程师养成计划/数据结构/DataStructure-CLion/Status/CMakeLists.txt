# 生成一个静态库，因为所有其他数据结构都会引用这个库
add_library(Scanf_lib Status.c)

# 生成可执行文件
add_executable(Status-exec Status.c Status.h Status-main.c)

# 记录要拷贝到*.exe目录下的资源文件
file(GLOB TestData TestData*.txt)
# 将资源文件拷贝到*.exe目录下，不然无法加载
file(COPY ${TestData} DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
# 注：
## 本次修订的目的包括降低耦合，争取每个模块都可以单独运行
## 但是Status这个模块会被所有其他模块引用，引用次数很多。
## 如果直接将Status模块复制到其它模块中，则会导致太多重复代码，
## 因此这里生成一个公共静态库让其它模块共享比较划算
