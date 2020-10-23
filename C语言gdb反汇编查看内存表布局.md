
# C语言程序用gdb反汇编 查看内存布局

## 一

### 查资料

#### 阅读资料了解GDB

- [如何通过gdb查看反汇编代码](https://blog.csdn.net/counsellor/article/details/80686758)

- [GDB Command Reference - disassemble command](https://visualgdb.com/gdbreference/commands/disassemble)

- [MAC上使用gdb(完美解决)](https://blog.csdn.net/github_33873969/article/details/78511733)

### 源码和步骤

#### c源代码

```c
/*
vim main.c
write some code
gcc -o main main.c
gdb ./pof //这个命令在我这里应该是 gdb main
pdisass main 反汇编main // 这个命令实践是错的。
应该是 disassemble main

p/x ‘g’ 打印字母 g 的十六进制值。

显示栈空间的内存

断点：
b *main + 43
运行：
r
显示栈空间的内存：
x/24xb &rbp-0x19

*/

int main(void)
{
  char a = 'g';
  char *b = &a;
  char **p = &b;
  return 0;

}
```

### 反汇编过程

#### 反汇编

```text
(gdb) disassemble main
Dump of assembler code for function main:
   0x0000000100000f90 <+0>: push   %rbp
   0x0000000100000f91 <+1>: mov    %rsp,%rbp
   0x0000000100000f94 <+4>: xor    %eax,%eax
   0x0000000100000f96 <+6>: movl   $0x0,-0x4(%rbp)
   0x0000000100000f9d <+13>: movb   $0x67,-0x5(%rbp)
   0x0000000100000fa1 <+17>: lea    -0x5(%rbp),%rcx
   0x0000000100000fa5 <+21>: mov    %rcx,-0x10(%rbp)
   0x0000000100000fa9 <+25>: lea    -0x10(%rbp),%rcx
   0x0000000100000fad <+29>: mov    %rcx,-0x18(%rbp)
   0x0000000100000fb1 <+33>: pop    %rbp
   0x0000000100000fb2 <+34>: retq
End of assembler dump.
```

#### 查看下字母 `g` 的值

```text

(gdb) p/x 'g'
$1 = 0x67

```

