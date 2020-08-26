//
// Created by panda8z on 2020/8/26.
//

#include "Status.h"

int main(void) {
    FILE* fp;

    char c, tmp;
    // 打开文件，准备读取测试数据
    fp = fopen("TestData_Pre1.txt", "r");
    if (fp == NULL) {
        return ERROR;
    }
    while (!feof(fp)) {
        c = getc(fp);
        printf("ch: %c\n", c);
    }
//    while((tmp = getc(fp)) != EOF) {
//        // 遇到首个西文字符，将此西文字符重新放入输入流
//        if((tmp >= 0 && tmp <= 127)) {
//            ungetc(tmp, fp);
//            break;
//        }
//    }

    fclose(fp);
    return 0;
}