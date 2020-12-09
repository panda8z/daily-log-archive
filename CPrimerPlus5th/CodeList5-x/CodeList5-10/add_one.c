/* add_one.c -- 增量: 前缀后缀 */
#include <stdio.h>
int main (void)
{
    int ultra = 0, super = 0;
    while(super < 5)
    {
        super++;
        ++ultra;
        printf("super = %d, ultra = %d \n",super, ultra);
    }
    return 0;
}