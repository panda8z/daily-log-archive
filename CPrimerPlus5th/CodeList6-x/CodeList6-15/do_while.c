/* do_while.c -- 退出循环的条件 */
#include <stdio.h>
int main (void)
{
    const int secret_code = 13;
    int code_entered;

    do
    {
        printf("To enter the triskaidekaphobia therapy club. \n");
        printf("Please enter the secret code number:");
        scanf("%d", &code_entered);
    } while(code_entered != secret_code);

    printf("Congratulations! you are cured!\n");
    return 0;
}