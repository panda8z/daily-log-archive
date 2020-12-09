/* showchar2.c -- 按行和列打印字符 */
#include <stdio.h>
void display(char cr, int line, int width);
int main(void)
{
    int ch;//要打印的字符
    int rows, cols; //行数和列数
    printf("Enter a character an two integers: \n");
    while((ch = getchar()) != '\n')
    {
        if(scanf("%d %d", &rows, &cols) != 2)
        {
            break;
        }
        while(getchar() != '\n')
        {
            continue;
        }
        display(ch, rows, cols);
        printf("Enter anther character and two integers: \n");
        printf("Enter a newline to quite.\n");
    }
    printf("Bye!\n");
    return 0;
}
void display(char cr, int lines, int width)
{
    int row, col;

    for(row = 1; row <= lines; row++)
    {

        for(col = 1; col <= width; col++)
        {
            putchar(cr);
        }
        putchar('\n');//结束本行,开始新的一行
    }
}