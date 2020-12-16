/* echo.c -- 重复输入 */
#include <stdio.h>
int main(void)
{
    char ch;
    while((ch = getchar()) != '#')
    {
        putchar(ch);
        putchar('\n');
    }
    return 0;
}