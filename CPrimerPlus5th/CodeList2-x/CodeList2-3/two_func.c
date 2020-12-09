/* two_func.c -- 在一个文件中使用两个函数 */

#include <stdio.h>

void bulter(void); /* ISO/ANSI C函数原型 */

int main(void)
{
    printf("I will summon bulter function.\n");
    bulter();
    printf("Yes，bring me some tea and writebale CD_ROMS.\n");
    return 0;
}

void bulter(void)
{
    printf("You range, sir?\n");
}
