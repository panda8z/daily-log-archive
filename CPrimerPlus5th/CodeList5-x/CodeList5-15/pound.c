/* pound.c -- 定义带有一个参数的函数 */
#include <stdio.h>
void pound(int n);/* NASI风格的原型 */
int main(void)
{
    int times = 5;
    char ch = '!';//ASCII码值为33
    float f = 6.0;
    pound(times);//int 参数
    pound(ch);//char参数自动转换为int类型
    pound((int)f); //指派运算符把f强制转换为int类型
    return 0;
}

void pound(int n) //ANSI风格的函数头,
                // 说明该函数接受一个int参数
{
    while( n-- > 0)
    {
        printf("#");
    }
    printf("\n");
}