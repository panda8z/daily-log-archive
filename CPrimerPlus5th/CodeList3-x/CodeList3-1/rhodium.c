/* rhodium.c -- 用金属铑衡量您的体重 */

#include <stdio.h>
int main(void)
{
    float weight; // 用户的体重
    float value;  // 相等重量的铑的价值
    printf("Are you worth your weight in rhodium?\n");
    printf("Let`s check it out. \n");
    printf("Please entr your weight in pounds: ");
    /* 从用户输入 */
    scanf("%f", &weight);
    // 假设铑盎司770美元
    // 14.5833 把常衡制的英镑转换为金衡制的盎司
    value = 770 * weight * 14.5833;
    printf("Your weight in rhodium is worth $%0.2f.\n", value);
    printf("You are easily worth that! If rhodium prices drop. \n");
    printf("eat more to maintain your value.\n");
    return 0;
}