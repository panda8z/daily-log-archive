/* sizeof.c -- 使用sizeof运算符 */
// 使用C99的%z修饰符.如果不能使用%zd,请使用%u或%1u

#include <stdio.h>
int main(void)
{
    int n = 0;
    size_t  intsize;
    intsize = sizeof(int);// C规定sizeof返回size_t类型的值
    printf("n = %d, n has %zd bytes; all ints have %zd bytes.\n",
    n, sizeof n, intsize);
    return 0;

}