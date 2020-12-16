/* guess.c -- 一个低效且错误的猜数程序 */
#include <stdio.h>
int main(void)
{
    int guess = 1;
    printf("Pick an integer from 1 to 100. I will you to guess ");
    printf("it.\nRespond with a y if my guess is right an with");
    printf("\nan n if it is wrong.\n");
    printf("Uh...is your number %d?\n", guess);
    while (getchar() != 'y')
    {
        printf("Well, then, is  it %d? \n", guess++);
    }
    printf("I knew I could do it!\n");
    return 0;
}