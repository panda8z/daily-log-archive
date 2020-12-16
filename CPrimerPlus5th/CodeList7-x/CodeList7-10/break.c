/* break.c -- 使用break跳出循环 */
#include <stdio.h>
int main(void)
{
    float length, width;

    printf("Enter the lenght of the rectangle: \n");
    while (scanf("%f", &length) == 1)
    {
        printf("Length = %0.2f;\n", length);
        printf("Enter it`s width: \n");
        if (scanf("%f", &width) != 1)
        {
            break;
        }
        printf("Width = %0.2f;\n", width);
        printf("Area = %0.2f;\n", width * length);
        printf("Enter the length of the Rectangle; \n");
    }
    printf("Done. \n");
    return 0;
}