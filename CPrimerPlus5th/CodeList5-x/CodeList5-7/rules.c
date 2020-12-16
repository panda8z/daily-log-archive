/* rules.c -- 优先级规则的试验 */
#include <stdio.h>
int main(void)
{
    int top, score;
    top = score = -(2 + 5) * 6 + (4 + 3 * (2 + 3));
    // 先预演一下,从左到右, - 7 * 6 + ( 4 + 3 * 5)
    // - 42 + (4 + 15)
    // - 42 + 19
    // - 23

    // 第一次预演 失败!

    // 第二次开始
    // 第一次翻了小学数学的错误,弱智.
    printf("top = %d \n", top);
    return 0;
}