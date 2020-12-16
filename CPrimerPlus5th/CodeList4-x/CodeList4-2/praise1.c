/* praise1.c -- 使用不同类别的字符串 */
#include <stdio.h>
#define PRAISE "What a super marvelous name!"
int main(void)
{
    char name[40];
    printf("What`s your name?\n");
    scanf("%s", name);
    printf("Hello %s,  %s\n", name, PRAISE);
    return 0;
}