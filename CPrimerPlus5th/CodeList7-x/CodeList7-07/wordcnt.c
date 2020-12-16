/* wordcnt.c -- 统计字符,单词和行 */
#include <stdio.h>
#include <ctype.h>   //为isspace()提供函数原型
#include <stdbool.h> // 为bool,true,false 提供定义.
#define STOP '|'
int main(void)
{
    char c;              // 读入字符
    char prev;           // 前一个读入的字符
    long n_chars = 0l;   // 字符数
    int n_lines = 0;     // 行数
    int n_words = 0;     // 单词数
    int p_lines = 0;     // 不完整的行数
    bool inword = false; // 如果c在一个单词中,则inword等于true

    printf("Enter text to be analyzed (| to terminate): \n");
    prev = '\n';
    while ((c = getchar()) != STOP)
    {
        n_chars++; // 统计字符
        if (c == '\n')
        {
            n_lines++; // 统计行数
        }

        if (!isspace(c) && !inword)
        {
            inword = true; // 开始一个新单词
            n_words++;     // 统计单词
        }

        if (isspace(c) && inword)
        {
            inword = false; // 到达单词的尾部
        }
        prev = c;
    }

    if (prev != '\n')
    {
        p_lines++;
    }

    printf("Characters = %ld, words = %d, lines = %d,",
           n_chars, n_words, n_lines);

    printf("partial lines = %d\n", p_lines);

    return 0;
}