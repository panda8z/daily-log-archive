
# CodeList 7-x

## CodeList 7-1 at P

### code7-1

```c
/* colddays.c -- 求温度低于零度的天数的百分率 */
#include <stdio.h>
int main(void)
{
    const int FREEZING = 0;
    float temperature;
    int cold_days = 0;
    int all_days = 0;

    printf("Enter the list of daily low temperatures.\n");
    printf("Use Celsius, and enter q to quit.\n");
    while(scanf("%f", &temperature) == 1)
    {
        all_days++;
        if(temperature < FREEZING)
        {
            cold_days++;
        }
    }
    if(all_days != 0)
    {
        printf("%d days total: %.1f%% were below freezing.\n",
         all_days, 100.0 * (float) cold_days / all_days);
    }

    if(all_days == 0)
    {
        printf("No data entered!\n");
    }
    return 0;
}

```

### log7-1

```bash
Panda-MBP:CodeList7-1 panda8z$ gcc -o colddays.out colddays.c
Panda-MBP:CodeList7-1 panda8z$ ./colddays.out
Enter the list of daily low temperatures.
Use Celsius, and enter q to quit.
23
23
12
45
67
13
98
112
3
q
9 days total: 0.0% were below freezing.
Panda-MBP:CodeList7-1 panda8z$

```

## CodeList 7-2 at P

### code7-2

```c
/* cypher1.c -- 改变输入, 只保留其中的空格 */
#include <stdio.h>
#define SPACE ' ' /* SPACE 相当于 引号-空格-引号 */
int main(void)
{
    char ch;
    ch = getchar();    //读入一个字符
    while (ch != '\n') // 当一行未结束时
    {
        if (ch == SPACE)
        {
            putchar(ch); //不改变这个字符
        }
        else
        {
            putchar(ch + 1); //改变其他字符
        }
        ch = getchar(); //获取下一个字符
    }
    putchar(ch);
    return 0;
}

```

### log7-2

```bash

Panda-MBP:CodeList7-2 panda8z$ gcc -o cypher1.out cypher1.c
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
er
fs
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
ee
ff
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
w
x
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
w
x
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
we
xf
Panda-MBP:CodeList7-2 panda8z$ ./cypher1.out
Call me hal.
Dbmm nf ibm/
Panda-MBP:CodeList7-2 panda8z$
```

## CodeList 7-3 at P

### code7-3

```c
/* cypher.c -- 改变输入, 只保留非字母字符 */
#include <stdio.h>
#include<ctype.h>
int main(void)
{
    char ch;
    while((ch = getchar()) != '\n')
    {
        if(isalpha(ch))
        {
            putchar(ch + 1);
        }
        else
        {
            putchar(ch);
        }
    }
    putchar(ch);
    return 0;
}

```

### log7-3

```bash

Panda-MBP:CodeList7-3 panda8z$ gcc -o cypher2.out cypher2.c
Panda-MBP:CodeList7-3 panda8z$ ./cypher2.out
1234567890 566
1234567890 566
Panda-MBP:CodeList7-3 panda8z$ ./cypher2.out
abcdef abc
bcdefg bcd
Panda-MBP:CodeList7-3 panda8z$ gcc -o cypher2.out cypher2.c
Panda-MBP:CodeList7-3 panda8z$ ./cypher2.out
123456789  56
123456789  56Panda-MBP:CodeList7-3 panda8z$
Panda-MBP:CodeList7-3 panda8z$ ./cypher2.out
abcd abc
bcde bcdPanda-MBP:CodeList7-3 panda8z$
```

## CodeList 7-4 at P

### code7-4

```c
/* electric.c -- 计算电费账目 */
#include <stdio.h>
#define RATE1 0.12589                               //第一个360kwh的费率
#define RATE2 0.17901                               //下一个360kwh的费率
#define RATE3 0.20971                               //超过680kwh的费率
#define BREAK1 360.0                                //费率的第一个分界点
#define BREAK2 680.0                                // 费率的第二个分界点
#define BASE1 (RATE2 * BREAK1)                      //用电360kwh的费用
#define BASE2 (BASE1 - (RATE2 * (BREAK2 - BREAK1))) //用电680kwh的费用
int main(void)
{
    double kwh;
    double bill;

    printf("Please enter the kwh used.\n");
    scanf("%lf", &kwh);
    if (kwh <= BREAK1)
    {
        bill = RATE1 * kwh;
    }
    else if (kwh <= BREAK2)
    {
        bill = BASE1 + (RATE2 * (kwh - BREAK1));
    }
    else
    {
        bill = BASE2 + (RATE3 * (kwh - BREAK2));
    }

    printf("The charge for %.1f kwh is $%1.2f.\n", kwh, bill);
    return 0;
}

```

### log7-4

```bash
Panda-MBP:CodeList7-4 panda8z$ gcc -o electric.out electric.c
Panda-MBP:CodeList7-4 panda8z$ ./electric.out
Please enter the kwh used.
330
The charge for 330.0 kwh is $41.54.
Panda-MBP:CodeList7-4 panda8z$ ./electric.out
Please enter the kwh used.
370
The charge for 370.0 kwh is $66.23.
Panda-MBP:CodeList7-4 panda8z$ ./electric.out
Please enter the kwh used.
690
The charge for 690.0 kwh is $9.26.
Panda-MBP:CodeList7-4 panda8z$

```

## CodeList 7-5 at P

### code7-5

```c
/* divisors.c -- 使用嵌套if显示一个数的约数 */
#include <stdio.h>
#include <stdbool.h>
int main(void)
{
    unsigned long num; // 要检查的数
    unsigned long div; // 可能的素数
    bool isPrime;      // 素数的标志

    printf("Please enter an Integer for analysis: ");
    printf("Enter q to quit.\n");
    while (scanf("%lu", &num) == 1)
    {
        for (div = 2, isPrime = true; (div * div) <= num; div++)
        {
            if (num % div == 0)
            {
                if ((div * div) != 0)
                {
                    printf("%lu is divisible by %lu and %lu.\n",
                            num, div, num / div);
                }
                else
                {
                    printf("%lu is divisible by %lu.\n", num, div);
                }
                isPrime = false; //不是一个素数
            }
        }
        if (isPrime)
        {
            printf("%lu is prime.\n", num);
        }
        printf("Please enter another integer for analysis: ");
        printf("Enter q or quit.\n");
    }
    printf("Bye.\n");
    return 0;
}

```

### log7-5

```bash
Panda-MBP:CodeList7-5 panda8z$ gcc -o dicisors.out dicisors.c
dicisors.c:16:27: error: expression is not assignable
            if (num % div = 0)
                ~~~~~~~~~ ^
dicisors.c:20:21: warning: implicit declaration of function 'printff' is invalid in C99 [-Wimplicit-function-declaration]
                    printff("%lu is divisible by %lu and %lu.\n",
                    ^
dicisors.c:27:32: error: expected ';' after expression
                isPrime = false //不是一个素数
                               ^
                               ;
1 warning and 2 errors generated.
Panda-MBP:CodeList7-5 panda8z$ gcc -o dicisors.out dicisors.c
dicisors.c:20:21: warning: implicit declaration of function 'printff' is invalid in C99 [-Wimplicit-function-declaration]
                    printff("%lu is divisible by %lu and %lu.\n",
                    ^
1 warning generated.
Undefined symbols for architecture x86_64:
  "_printff", referenced from:
      _main in dicisors-af5660.o
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
Panda-MBP:CodeList7-5 panda8z$ gcc -o dicisors.out dicisors.c
Panda-MBP:CodeList7-5 panda8z$ ./dicisors.out
Please enter an Integer for analysis: Enter q to quit.
45
45 is divisible by 3 and 15.
45 is divisible by 5 and 9.
Please enter another integer for analysis: Enter q or quit.
43
43 is prime.
Please enter another integer for analysis: Enter q or quit.
q
Bye.
Panda-MBP:CodeList7-5 panda8z$

```

## CodeList 7-6 at P

### code7-6

```c

/* chcount.c -- 使用逻辑与运算符 */
#include <stdio.h>
#define PERIOD '.'
int main (void)
{
    int ch;
    int charcount = 0;
    while((ch = getchar()) != PERIOD)
    {
        if(ch != '"' && ch != '\'')
        {
            charcount++;
        }
    }
    printf("There are %d non-quote characters.\n", charcount);
    return 0;
}
```

### log7-6

```bash
Panda-MBP:CodeList7-06 panda8z$ gcc -o chcount.out chcount.c
Panda-MBP:CodeList7-06 panda8z$ ./chcount.out
234
234234234234235234234234234234124123123123
sdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdf
ssdf
sdf
sdf
sdf
sdf
sd
fs
dfs
df

.
There are 119 non-quote characters.
Panda-MBP:CodeList7-06 panda8z$ ./chcount.out
"""""""""""""""""''''''''''''''123
.
There are 4 non-quote characters.
Panda-MBP:CodeList7-06 panda8z$

```

## CodeList 7-7 at P

### code7-7

```c
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

```

### log7-7

```bash
Panda-MBP:CodeList7-07 panda8z$ ./wordcnt.out
Enter text to be analyzed (| to terminate):
asdjasd;fjak;ldsjfalkdsjfkl sdklfjals;kd

sdfajsdf;klj

sadf;asf
asdfj
sadf  
asdfjkjkl

asdfjakj;

adf;
|
Characters = 106, words = 9, lines = 12,partial lines = 0
Panda-MBP:CodeList7-07 panda8z$

```

## CodeList 7-8 at P

### code7-8

```c

/* paint.c -- 使用条件运算 */
#include <stdio.h>
#define COVERAGE 200 // 每罐漆可喷的平方英尺数
int main(void)
{
    int sq_feet;
    int cans;

    printf("Enter number of square feet to be painted: \n");
    while (scanf("%d", &sq_feet) == 1)
    {
        cans = sq_feet / COVERAGE;
        cans += ((sq_feet % COVERAGE == 0)) ? 0 : 1;

        printf("You need %d %s of paint.\n",
               cans, cans == 1 ? "can" : "cans");

        printf("Enter next value (q to quit): \n");
    }
    return 0;
}
```

### log7-8

```bash

Panda-MBP:CodeList7-08 panda8z$ gcc -o paint.out paint.c
Panda-MBP:CodeList7-08 panda8z$ ./paint.out
Enter number of square feet to be painted:
45
You need 1 can of paint.
Enter next value (q to quit):
32
You need 1 can of paint.
Enter next v

```

## CodeList 7-09 at P

### code7-09

```c
/* skippart.c -- 使用continue跳过部分循环 */
#include <stdio.h>
int main(void)
{
    const float MIN = 0.0f;
    const float MAX = 100.0f;

    float score;
    float total = 0.0f;
    int n = 0;
    float min = MIN;
    float max = MAX;
    printf("Enter the first score(q to quit): ");
    while (scanf("%f", &score) == 1)
    {
        if (score < MIN || score > MAX)
        {
            printf("%0.1f is an incalid value. Try again: ", score);
            continue;
        }
        printf("Accepting %0.1f: \n", score);
        min = (score < min) ? score : min;
        max = (score > max) ? score : max;
        total += score;
        n++;
        printf("Enter next score (q to quit):");
    }

    if (n > 0)
    {
        printf("Average of %d scores is %0.1f.\n", n, total / n);
        printf("Low = %0.1f, high = %0.1f\n", min, max);
    }
    else
    {
        printf("No valid scores were entered.\n");
    }
    return 0;
}

```

### log7-09

```bash
bogon:CodeList7-09 panda8z$ gcc -o skippart.out skippart.c
bogon:CodeList7-09 panda8z$ ./skippart.out
Enter the first score(q to quit): 45
Accepting 45.0:
Enter next score (q to quit):65
Accepting 65.0:
Enter next score (q to quit):67
Accepting 67.0:
Enter next score (q to quit):80
Accepting 80.0:
Enter next score (q to quit):99
Accepting 99.0:
Enter next score (q to quit):q
Average of 5 scores is 71.2.
Low = 0.0, high = 100.0

```

## CodeList 7-10 at P

### code7-10

```c
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

```

### log7-10

```bash
Panda-MBP:CodeList7-10 panda8z$ ./break.out 
Enter the lenght of the rectangle: 
4567 6788
Length = 4567.00;
Enter it`s width: 
Width = 6788.00;
Area = 31000796.00;
Enter the length of the Rectangle; 
23
Length = 23.00;
Enter it`s width: 
we
Done. 
Panda-MBP:CodeList7-10 panda8z$ 

```

## CodeList 7-11 at P

### code7-11

```c
/* animals.c -- 使用switch语句 */
#include <stdio.h>
#include <ctype.h>
int main(void)
{
    char ch;
    printf("Give me a letter of the alphabet, and I will give ");
    printf("an animal name\nbeginning with that letter.\n");
    printf("Please type in a letter: type # to end my act.\n");
    while ((ch = getchar()) != '#')
    {
        if (ch == '\n')
        {
            continue;
        }
        if (islower(ch)) //只识别小写字母
        {
            switch (ch)
            {
            case 'a':
                printf("argali, a wild sheep of Asia\n");
                break;
            case 'b':
                printf("babirusa, a wild pig of Malay\n");
                break;
            case 'c':
                printf("coati, racoonlike animal\n");
                break;
            case 'd':
                printf("desman, aquatic, molelike critter\n");
                break;
            case 'e':
                printf("echidna, the spiny anteater\n");
                break;
            case 'f':
                printf("fisher, brownish marten\n");
                break;
            default:
                printf("That`s a stumper!\n");
                break;
            }
        }
        else
        {
            printf("I recognize only lowercase letters.\n");
        }

        while (getchar() != '\n')
        {
            continue;
        }
        printf("Please type another letter or a #.\n");
    }
    printf("Bye!\n");
    return 0;
}

```

### log7-11

```bash
Panda-MBP:CodeList7-11 panda8z$ gcc -o animals.out animals.c 
animals.c:10:32: error: expected expression
    while ((ch = getchar()) != #)
                               ^
animals.c:55:13: error: expected ';' after return statement
    return 0
            ^
            ;
2 errors generated.
Panda-MBP:CodeList7-11 panda8z$ gcc -o animals.out animals.c 
Panda-MBP:CodeList7-11 panda8z$ ./animals.out 
Give me a letter of the alphabet, and I will give an animal name
beginning with that letter.
Please type in a letter: type # to end my act.
fr
fisher, brownish marten
Please type another letter or a #.
d
desman, aquatic, molelike critter
Please type another letter or a #.
sedsfsdfsdfsdf
That`s a stumper!
Please type another letter or a #.
sdsdsdfsdfsdfsdfsfsdfsfsdf
That`s a stumper!
Please type another letter or a #.
sdfsdfsdfsdf
That`s a stumper!
Please type another letter or a #.
eeee
echidna, the spiny anteater
Please type another letter or a #.
eeeee
echidna, the spiny anteater
Please type another letter or a #.
rrrrrr
That`s a stumper!
Please type another letter or a #.
www
That`s a stumper!
Please type another letter or a #.
#
Bye!
Panda-MBP:CodeList7-11 panda8z$ 

```

## CodeList 7-12 at P

### code7-12

```c
/* vowels.c -- 使用多重标签 */
#include <stdio.h>
int main(void)
{
    char ch;
    int a_ct, e_ct, i_ct, o_ct, u_ct;
    a_ct = e_ct = i_ct = o_ct = u_ct = 0;

    printf("Enter some text: enter # to quit.\n");
    while ((ch = getchar()) != '#')
    {
        switch (ch)
        {
        case 'a':
        case 'A':
            a_ct++;
            break;
        case 'e':
        case 'E':
            e_ct++;
            break;
        case 'i':
        case 'I':
            i_ct++;
            break;
        case 'o':
        case 'O':
            o_ct++;
            break;
        case 'u':
        case 'U':
            u_ct++;
            break;

        default:
            break;
        }
    }

    printf("number of vowels: A E I O U\n");
    printf("                  %4d %4d %4d %4d %4d",
           a_ct, e_ct, i_ct, o_ct, u_ct);
    return 0;
}

```

### log7-12

```bash
Panda-MBP:CodeList7-12 panda8z$ gcc -o vowels.out vowels.c
Panda-MBP:CodeList7-12 panda8z$ ./vowels.out
Enter some text: enter # to quit.
weaksdnfkjadhflkajhdflkjsakcljnalkjn asj ahlfjkah lakj fh ha 78987987#
number of vowels: A E I O U
                    10    1    0    0    0Panda-MBP:CodeList7-12 panda8z$ ./vowels.out
Enter some text: enter # to quit.
^[[A^[[A
#
number of vowels: A E I O U
                     2    0    0    0    0Panda-MBP:CodeList7-12 panda8z$ gcc -o vowels.out vowels.c
Panda-MBP:CodeList7-12 panda8z$ ./vowels.out
Enter some text: enter # to quit.
haljkdhflakjshfjkasdhfkjsdklfjasfkljuqpeirqioweuroiqwuerpiwuepirquewiorjskfalksdjfa;klsjdf;adf
#
number of vowels: A E   I   O   U
                     7    5    6    3    5Panda-MBP:CodeList7-12 panda8z$ gcc -o vowels.out vowels.c
Panda-MBP:CodeList7-12 panda8z$ ./vowels.out
Enter some text: enter # to quit.
ssfasdhfakdsjfadsjf;lkajsdf;kajsf
#
number of vowels:       A       E       I       O       U
                     5    0    0    0    0Panda-MBP:CodeList7-12 panda8z$ 

```

## CodeList 7-13 at P

### code7-13

```c


```

### log7-13

```bash


```

## CodeList 7-14 at P

### code7-14

```c


```

### log7-14

```bash


```

## CodeList 7-15 at P

### code7-15

```c


```

### log7-15

```bash


```

## CodeList 7-16 at P

### code7-16

```c


```

### log7-16

```bash


```

## CodeList 7-17 at P

### code7-17

```c


```

### log7-17

```bash


```

## CodeList 7-18 at P

### code7-18

```c


```

### log7-18

```bash


```

## CodeList 7-19 at P

### code7-19

```c


```

### log7-19

```bash


```

## CodeList 7-20 at P

### code7-

```c


```

### log7-

```bash


```