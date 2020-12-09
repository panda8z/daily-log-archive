
# CodeList 8-x

## CodeList 8-1 at P

### code8-1

```c
/* echo.c -- 重复输入 */
#include <stdio.h>
int main(void)
{
    char ch;
    while((ch = getchar()) != '#')
    {
        putchar(ch);
        putchar('\n');
    }
    return 0;
}

```

### log8-1

```bash
Panda-MBP:CodeList8-01 panda8z$ gcc -o echo.out echo.c.Panda-MBP:CodeList8-01 panda8z$ ./echo.outr
r
t
t
ttttt
ttttt
tttt
tttt
tttttegjkrl;aldkja;kldsjf;alsdjf;ajdsf;adjsf;lkajdsflkajscncvz.,xcnp8493457138475109387541
tttttegjkrl;aldkja;kldsjf;alsdjf;ajdsf;adjsf;lkajdsflkajscncvz.,xcnp8493457138475109387541




#
Panda-MBP:CodeList8-01 panda8z$ gcc -o echo.out echo.cecho.c:9:22: error: expected ';' after expression
        putchar('\n')
                     ^
                     ;
1 error generated.
Panda-MBP:CodeList8-01 panda8z$ gcc -o echo.out echo.cPanda-MBP:CodeList8-01 panda8z$ ./echo.outsfdfga
s
f
d
f
g
a


asdfasfasdfas2382039482093840924
a
s
d
f
a
s
f
a
s
d
f
a
s
2
3
8
2
0
3
9
4
8
2
0
9
3
8
4
0
9
2
4


#
Panda-MBP:CodeList8-01 panda8z$
```

## CodeList 8-2 at P

### code8-2

```c
/* echo_eof.c -- 重复输入直到文件结尾 */
#include <stdio.h>
int main(void)
{
    int ch;

    while((ch = getchar()) != EOF)
    {
        putchar(ch);
    }
    return 0;
}

```

### log8-2

```bash
Panda-MBP:CodeList8-02 panda8z$ gcc -o echo_eof.out echo_eof.cecho_eof.c:9:17: error: too few arguments to function call, expected 1, have 0
        putchar();
        ~~~~~~~ ^
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/usr/include/stdio.h:172:1: note: 'putchar' declared
      here
int      putchar(int);
^
1 error generated.
Panda-MBP:CodeList8-02 panda8z$ gcc -o echo_eof.out echo_eof.cPanda-MBP:CodeList8-02 panda8z$ ./echo_eof.outer
er
er
er
er er sdfafadadadfafd
sdfafadadadfafd
^Z
[1]+  Stopped                 ./echo_eof.out
Panda-MBP:CodeList8-02 panda8z$ ./echo_eof.outssdfasdfasdf^C
Panda-MBP:CodeList8-02 panda8z$ ./echo_eof.outsdfsdsdfsd

Panda-MBP:CodeList8-02 panda8z$
```

## CodeList 8-3 at P

### code8-3

```c
/* file_eof.c  -- 打开一个文件并显示其内容 */
#include <stdio.h>
#include <stdlib.h> // 为了使用exit()
int main()
{
    int ch;
    FILE *fp;
    char fname[50]; //用于存放文件名

    printf("Enter the name of the file: ");
    scanf("%s", fname);
    fp = fopen(fname, "r");
    if (fp == NULL)
    {
        printf("Failed to open file. Bye\n");
        exit(1); //终止程序
    }
    // getc(fp)从打开的文件中获取一个字符
    while((ch = getc(fp)) != EOF)
    {
        putchar(ch);
    }
    fclose(fp); //关闭文件
    return 0;
}

```

### log8-3

```bash

Panda-MBP:CodeList8-03 panda8z$ gcc -o file_eof.out file_eof.c
Panda-MBP:CodeList8-03 panda8z$ ./file_eof.outEnter the name of the file: test.md
# test file

## second

### third

#### fouth

##### fifthPanda-MBP:CodeList8-03 panda8z$
```

## CodeList 8-4 at P

### code8-4

```c
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

```

### log8-4

```bash

Panda-MBP:CodeList8-04 panda8z$ gcc -o guess.out guess.cguess.c:11:5: error: expected ')'
    {
    ^
guess.c:10:10: note: to match this '('
    while((getchar() != 'y')
         ^
1 error generated.
Panda-MBP:CodeList8-04 panda8z$ gcc -o guess.out guess.cPanda-MBP:CodeList8-04 panda8z$ ./guess.outPick an integer from 1 to 100. I will you to guess it.
Respond with a y if my guess is right an with
an n if it is wrong.
Uh...is your number 1?
7
Well, then, is  it 1?Well, then, is  it 2?n
Well, then, is  it 3?Well, then, is  it 4?n
Well, then, is  it 5?Well, then, is  it 6?y
I knew I could do it!
Panda-MBP:CodeList8-04 panda8z$```

## CodeList 8-5 at P

### code8-5

```c
/* showchar1.c -- 带有一个较大的I/O问题的程序 */
#include <stdio.h>
void display(char cr, int line, int width);
int main(void)
{
    int ch;//要打印的字符
    int rows, cols; //行数和列数
    printf("Enter a character an two integers: \n");
    while((ch = getchar()) != '\n')
    {
        scanf("%d %d", &rows, &cols);
        display(ch, rows, cols);
        printf("Enter anther character and two integers: \n");
        printf("Enter a newline to quite.\n");
    }
    printf("Bye!\n");
    return 0;
}
void display(char cr, int lines, int width)
{
    int row, col;

    for(row = 1; row <= lines; row++)
    {

        for(col = 1; col <= width; col++)
        {
            putchar(cr);
        }
        putchar('\n');//结束本行,开始新的一行
    }
}

```

### log8-5

```bash

Panda-MBP:CodeList8-05 panda8z$ gcc -o showchar1.out showchar1.cUndefined symbols for architecture x86_64:
  "_display", referenced from:
      _main in showchar1-05cc4c.o
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
Panda-MBP:CodeList8-05 panda8z$ gcc -o showchar1.out showchar1.cshowchar1.c:23:25: error: use of undeclared identifier 'lines'; did you mean 'line'?
    for(row = 1; row <= lines; row++)
                        ^~~~~
                        line
showchar1.c:19:27: note: 'line' declared here
void display(char cr, int line, int width)
                          ^
1 error generated.
Panda-MBP:CodeList8-05 panda8z$ gcc -o showchar1.out showchar1.cPanda-MBP:CodeList8-05 panda8z$ ./showchar1.outEnter a character an two integers:s 4 5
sssss
sssss
sssss
sssss
Enter anther character and two integers:Enter a newline to quite.
Bye!
Panda-MBP:CodeList8-05 panda8z$ ./showchar1.outEnter a character an two integers:1  100  500
11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111
...此处省略98行
11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111
Enter anther character and two integers:Enter a newline to quite.
Bye!
Panda-MBP:CodeList8-05 panda8z$```

## CodeList 8-6 at P

### code8-6

```c
/* showchar2.c -- 按行和列打印字符 */
#include <stdio.h>
void display(char cr, int line, int width);
int main(void)
{
    int ch;//要打印的字符
    int rows, cols; //行数和列数
    printf("Enter a character an two integers: \n");
    while((ch = getchar()) != '\n')
    {
        if(scanf("%d %d", &rows, &cols) != 2)
        {
            break;
        }
        while(getchar() != '\n')
        {
            continue;
        }
        display(ch, rows, cols);
        printf("Enter anther character and two integers: \n");
        printf("Enter a newline to quite.\n");
    }
    printf("Bye!\n");
    return 0;
}
void display(char cr, int lines, int width)
{
    int row, col;

    for(row = 1; row <= lines; row++)
    {

        for(col = 1; col <= width; col++)
        {
            putchar(cr);
        }
        putchar('\n');//结束本行,开始新的一行
    }
}

```

### log8-6

```bash

Panda-MBP:CodeList8-06 panda8z$ gcc -o showchar2.out showchar2.cPanda-MBP:CodeList8-06 panda8z$ ./showchar2.outEnter a character an two integers:345^[[D^[[D
Bye!
Panda-MBP:CodeList8-06 panda8z$ ./showchar2.outEnter a character an two integers:2 3 4
2222
2222
2222
Enter anther character and two integers:Enter a newline to quite.
e
5 6 .
eeeeee
eeeeee
eeeeee
eeeeee
eeeeee
Enter anther character and two integers:Enter a newline to quite.
^[[A^[[B
Bye!
Panda-MBP:CodeList8-06 panda8z$ ./showchar2.outEnter a character an two integers:c
4 5
ccccc
ccccc
ccccc
ccccc
Enter anther character and two integers:Enter a newline to quite.
^C
Panda-MBP:CodeList8-06 panda8z$
```

## CodeList 8-7 at P

### code8-7

```c
/* checking.c -- 输入确认 */
#include <stdio.h>
#include <stdbool.h>
int get_int(void);                                        //确认输入了一个整数
bool bad_limits(int begin, int end, int low, int height); //确认范围的上下届是否有效
double sum_squares(int a, int b);                         //计算从a到b之间的整数平方和
int main(void)
{
    const int MIN = -1000; //范围的下界限制
    const int MAX = +1000; //范围的上界限制
    int start;             //范围的下界
    int stop;              //范围的上界
    double answer;
    printf("This program comutes the sum of the squares of "
           "integers in a range.\nThe lower bound should not "
           "be less than -1000 and\nthe upper bound should not "
           "be more than +1000.\nEnter the limits (enter 0 for "
           "both limits to quit):\nlower limit: ");
    start = get_int();
    printf("upper limit: ");
    stop = get_int();
    while (start != 0 || stop != 0)
    {
        if (bad_limits(start, stop, MIN, MAX))
        {
            printf("Please try again!\n");
        }
        else
        {
            answer = sum_squares(start, stop);
            printf("The sum of the squares of the integers form ");
            printf("%d to %d is %g\n", start, stop, answer);
        }
        printf("Enter the limits(enter 0 for both limits to quit): \n");
        printf("lower_limit: ");
        start = get_int();
        printf("upper_limit: ");
        stop = get_int();
    }
    printf("Done!\n");
    return 0;
}

int get_int(void)
{
    int input;
    char ch;
    while (scanf("%d", &input) != 1)
    {
        while ((ch = getchar()) != '\n')
        {
            putchar(ch); //剔除错误的输入
        }
        printf(" is not an integer.\nPlease enter an ");
        printf("integer value, such as 25, -178, or 3:");
    }
    return input;
}

double sum_squares(int a, int b)
{
    double total = 0;
    int i;
    for (i = 0; i <= b; i++)
    {
        total += i * i;
    }
    return total;
}

bool bad_limits(int begin, int end, int low, int high)
{
    bool not_good = false;
    if (begin > end)
    {
        printf("%d isn`t smaller than %d.\n", begin, end);
        not_good = true;
    }

    if (begin < low || end < low)
    {
        printf("Values must be %d or greatter.\n", low);
        not_good = true;
    }
    if (begin > high || end > high)
    {
        printf("Values must be %d or less.\n", high);
        not_good = true;
    }

    return not_good;
}

```

### log8-7

```bash
Panda-MBP:CodeList8-07 panda8z$ gcc -o checking.out checking.c 
checking.c:26:42: error: expected ';' after expression
            printf("Please try again!\n")
                                         ^
                                         ;
checking.c:76:47: error: use of undeclared identifier 'bergin'; did you mean 'begin'?
        printf("%d isn`t smaller than %d.\n", bergin, end);
                                              ^~~~~~
                                              begin
checking.c:71:21: note: 'begin' declared here
bool bad_limits(int begin, int end, int low, int high)
                    ^
checking.c:87:54: error: expected ';' after expression
        printf("Values must be %d or less.\n", high) : not_good = true;
                                                     ^
                                                     ;
3 errors generated.
Panda-MBP:CodeList8-07 panda8z$ gcc -o checking.out checking.c 
Panda-MBP:CodeList8-07 panda8z$ ./checking.out 
This program comutes the sum of the squares of integers in a range.
The lower bound should not be less than -1000 and
the upper bound should not be more than +1000.
Enter the limits (enter 0 for both limits to quit):
lower limit: 4
upper limit: 7
The sum of the squares of the integers form 4 to 7 is 140
Enter the limits(enter 0 for both limits to quit): 
lower_limit: -1001
upper_limit: 1001
Values must be -1000 or greatter.
Values must be 1000 or less.
Please try again!
Enter the limits(enter 0 for both limits to quit): 
lower_limit: 

```

## CodeList 8-8 at P

### code8-8

```c
/* menuette.c -- 菜单技术 */
#include <stdio.h>
char get_choice(void);
char get_first(void);
int get_int(void);
void count(void);
int main(void)
{
    int choice;
    void count(void);
    while ((choice = get_choice()) != 'q')
    {
        switch (choice)
        {
        case 'a':
            printf("Buy low, sell high.\n");
            break;
        case 'b':
            putchar('\n'); //ANSI
            break;
        case 'c':
            count();
            break;
        default:
            printf("Program error!\n");
            break;
        }
    }
    printf("Bye!\n");
    return 0;
}

void count(void)
{
    int n, i;
    printf("Count how far? Enter an integer: \n");
    n = get_int();
    for (i = 1; i <= n; i++)
    {
        printf("%d\n", i);
    }
    while (getchar() != '\n')
    {
        continue;
    }
}

char get_choice(void)
{
    int ch;
    printf("Enter the letter of your choice: \n");
    printf("a. advice           b. bell\n");
    printf("c. count            q. quit\n");
    ch = get_first();
    while ((ch < 'a' || ch > 'c') && ch != 'q')
    {
        printf("Please respond with a, b, c, or q.\n");
        ch = get_first();
    }
    return ch;
}

char get_first(void)
{
    int ch;
    ch = getchar();
    while (getchar() != '\n')
    {
        continue;
    }
    return ch;
}

int get_int(void)
{
    int input;
    char ch;
    while (scanf("%d", &input) != 1)
    {
        while ((ch = getchar()) != '\n')
        {
            putchar(ch); //剔除错误输入
        }
        printf(" is not an integer.\nPlease enter an ");
        printf("integer value, such as 25, -178, or 3:");
    }
    return input;
}


```

### log8-8

```bash
Panda-MBP:CodeList8-08 panda8z$ gcc -o menuette.out menuette.c 
menuette.c:66:20: error: expected ';' after expression
    ch = getchar() : while (getchar() != '\n')
                   ^
                   ;
1 error generated.
Panda-MBP:CodeList8-08 panda8z$ gcc -o menuette.out menuette.c 
Panda-MBP:CodeList8-08 panda8z$ ./menuette.out 
Enter the letter of your choice: 
a. advice           b. bell
c. count            q. quit
a
Buy low, sell high.
Enter the letter of your choice: 
a. advice           b. bell
c. count            q. quit
b

Enter the letter of your choice: 
a. advice           b. bell
c. count            q. quit
b

Enter the letter of your choice: 
a. advice           b. bell
c. count            q. quit
c
Count how far? Enter an integer: 
56
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
Enter the letter of your choice: 
a. advice           b. bell
c. count            q. quit
q
Bye!
Panda-MBP:CodeList8-08 panda8z$ 

```