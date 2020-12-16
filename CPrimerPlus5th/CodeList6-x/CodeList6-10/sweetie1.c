// sweetie1.c -- 一个计数循环
#include <stdio.h>
int main (void)
{
    const int NUMBER = 32;
    int count = 1;

    while(count <= NUMBER)
    {
        printf("Be my Valentine!\n");
        count++;
    }
    return 0;
}