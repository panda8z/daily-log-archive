/* nogood.c -- 含有错误的程序 */
/*
原程序start
#include <stdio.h>

int main(void)
{
    int n, int n2, int n3;
     该程序含有几个错误

    n = 5;
    n2 = n * n;
    n3 = n2 * n2;

    printf("n = %d, n squared = %d, n cubed = %d\n", n, n2, n3)

    return 0;
}
原程序end
*/

/*
1. 7line 定义错误，逗号改为分号
2. 8line 注释没有结尾
3. 14line打印语句没有分号

*/

#include <stdio.h>

int main(void)
{
    int n; int n2; int n3;
    /* 该程序含有几个错误*/

    n = 5;
    n2 = n * n;
    n3 = n2 * n2;

    printf("n = %d, n squared = %d, n cubed = %d\n", n, n2, n3);

    return 0;
}