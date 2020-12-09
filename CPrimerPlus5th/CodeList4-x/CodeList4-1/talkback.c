/* talkback.c -- 一个能为您提供一些信息的对话程序 */
#include <stdio.h>
#include <string.h>  // strlen()函数原型在这里
#define DENSITY 62.4 //人体密度(单位: 磅/立方英尺)
int main(void)
{
    float weight, volume;
    int size, letters;
    char name[40]; // name是一个有40个字符的数组.
    printf("Hi! What`s your first name?\n");
    scanf("%s", name);
    printf("%s, what`s your weight in pounds?\n", name);
    scanf("%f", &weight);
    size = sizeof name;
    letters = strlen(name);
    volume = weight / DENSITY;
    printf("Well! %s, your volume is %2.2f cubic feet.\n", name, volume);
    printf("Also, your first name has %d letters ", letters);
    printf("and we have %d bytes to stores it.\n", size);
    return 0;
}