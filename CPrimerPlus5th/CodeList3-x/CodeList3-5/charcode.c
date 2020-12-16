/* charcode.c -- 显示一个字符的编辑值 */
#include <stdio.h>
int main(void)
{
    char ch;
    printf("Please enter a charactr.\n");
    scanf("%c", &ch);
    printf("The code for %c is %d.\n", ch, ch);
    return 0;
}