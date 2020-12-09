/* print2.c -- printf()的更多属性 */
#include <stdio.h>
int main(void)
{
    unsigned int un = 3000000000; //int为32位
    short end = 200;              //short为16位系统
    long big = 6537;
    long long verybig = 12345678908642; //14位数

    printf("un = %u and not %d\n", un, un);
    printf("end = %hd and %d\n", end, end);
    printf("big = %1d and not %hd\n", big, big);
    printf("verybig = %1ld and not %1d\n", verybig, verybig);
    return 0;
}