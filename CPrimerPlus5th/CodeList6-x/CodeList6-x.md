
# CodeList 6-x

## CodeList 6-1 at P

### code6-1

```c
/* summing.c -- 对用户输入的整数求和 */
#include <stdio.h>
int main(void)
{
    long num;
    long sum = 0;
    int status;

    printf("Please enter an integer to summed. ");
    printf("q to quit): ");
    status = scanf("%ld", &num);
    while (status == 1)
    {
        sum = sum + num;
        printf("Please enter next integer(q to quit): ");
        status = scanf("%ld", &num);
    }
    printf("Thos integers sum to %ld.\n", sum);
    return 0;
}

```

### log6-1

```bash
Panda-MBP:CodeList6-1 panda8z$ gcc -o summing.out summing.c
Panda-MBP:CodeList6-1 panda8z$ ./summing.out
Please enter an integer to summed. q to quit): 8
Please enter next integer(q to quit): 09
Please enter next integer(q to quit): 89
Please enter next integer(q to quit): 89
Please enter next integer(q to quit): 89
Please enter next integer(q to quit): 1
Please enter next integer(q to quit): 1
Please enter next integer(q to quit): 2
Please enter next integer(q to quit): 3
Please enter next integer(q to quit): t
Thos integers sum to 291.
Panda-MBP:CodeList6-1 panda8z$

```

## CodeList 6-2 at P

### code6-2

```c
/* when.c -- 何时退出一个循环 */
#include <stdio.h>
int main(void)
{
    int n = 5;
    while (n < 7)
    {
        printf(" n = %d\n", n);
        n++;
        printf("Now n = %d\n", n);
    }

    printf("The loop has finished.\n");

    return 0;
}
```

### log6-2

```bash
bogon:CodeList6-02 panda8z$ gcc -o when.out when.c
bogon:CodeList6-02 panda8z$ ./when.out
 n = 5
Now n = 6
 n = 6
Now n = 7
The loop has finished.
bogon:CodeList6-02 panda8z$
```

## CodeList 6-3 at P

### code6-3

```c
/* while1.c -- 注意花括号的使用 */
/* 拙劣的代码产生了一个无限循环 */
#include <stdio.h>
int main(void)
{
    int n = 0;

    while(n < 3)
        printf("n is %d\n", n);
        n++;
    printf("That`s all this program does\n");
    return 0;
}

```

### log6-3

```bash

n is 0
n is 0
n is 0
n is 0
n is 0
[这里省略几千行]
n is 0
n is 0
n is 0
n is 0
n is 0
n is 0
^C  //直到手动结束这个程序才得以退出.
bogon:CodeList6-03 panda8z$
```

## CodeList 6-4 at P

### code6-4

```c
/* while2.c -- 注意分号的使用 */
#include <stdio.h>
int main (void)
{
    int n = 0;

    while(n++ < 3);
        printf("n is %d\n", n);
    printf("That`s all this program does.\n");
    return 0;
}

```

### log6-4

```bash

bogon:CodeList6-04 panda8z$ gcc -o while2.out while2.c
while2.c:7:19: warning: while loop has empty body [-Wempty-body]
    while(n++ < 3);
                  ^
while2.c:7:19: note: put the semicolon on a separate line to silence this warning
1 warning generated.
bogon:CodeList6-04 panda8z$ ./while1.out
bash: ./while1.out: No such file or directory
bogon:CodeList6-04 panda8z$ ./while2.out
n is 4
That`s all this program does.
bogon:CodeList6-04 panda8z$

```

## CodeList 6-5 at P

### code6-5

```c
/* cmpfit.c -- 浮点数比较 */
#include <stdio.h>
#include <math.h>
int main (void)
{
    const double ANWSER = 3.14159;
    double response;
    pirntf("What`s the value of pi?\n");
    scanf("%lf", &response);
    while(fabs(response - ANWSER) > 0.0001)
    {
        printf("Try again!\n");
        scanf("%lf", &response);
    }

    printf("Close enough!\n");
}

```

### log6-5

```bash

bogon:CodeList6-05 panda8z$ ./cmpfit.out
What`s the value of pi?
357
Try again!
3.1416
Close enough!
```

## CodeList 6-6 at P

### code6-6

```c
/* t_and_f.c -- C中的真和假 */
#include <stdio.h>
int main(void)
{
    int true_val, false_val;

    true_val = (10 > 2);
    false_val = (10 == 2);
    printf("true = %d, false = %d \n", true_val, false_val);
    return 0;
}

```

### log6-6

```bash
bogon:CodeList6-06 panda8z$ gcc -o t_and_f.out t_and_f.c
bogon:CodeList6-06 panda8z$ ./t_and_f.out
true = 1, false = 0

```

## CodeList 6-7 at P

### code6-7

```c

/* truth.c -- 哪些为真? */
#include <stdio.h>
int main (void)
{
    int n = 3;

    while(n){
        printf("%2d is true\n", n--);
    }
    printf("%2d is false\n", n);
    return 0;
}
```

### log6-7

```bash
bogon:CodeList6-07 panda8z$ gcc -o truth.out truth.c
bogon:CodeList6-07 panda8z$ ./truth.out
 3 is true
 2 is true
 1 is true
 0 is false

```

## CodeList 6-8 at P

### code6-8

```c
/* trouble.c -- 误用= */
#include <stdio.h>
int main(void)
{
    long num;
    long sum = 0;
    int status;

    printf("Please enter an integer to bu summed. ");
    printf("(q to quit) : ");
    status = scanf("%ld", &num);
    while(status = 1){
        sum = sum + num;
        printf("Please enter next integer (q to quit): ");
        status = scanf("%ld", &num);
    }
    printf("Those integer sum to %ld.\n", sum);
    return 0;
}

```

### log6-8

```bash
惨不忍睹...

```

## CodeList 6-09 at P

### code6-09

```c
// boolean.c --  使用_Bool变量

#include <stdio.h>
int main (void)
{
    long num;
    long sum = 0L;
    _Bool input_is_good;
    printf("Please enter an integer to be summed. ");
    printf("(q to quit): ");
    input_is_good = (scanf("%ld", &num) == 1);
    while(input_is_good)
    {
        sum = sum + num;
        printf("Please enter next integer (q to quit): ");
        input_is_good = (scanf("%ld", &num) == 1);
    }
    printf("Those integers sum to %ld.\n", sum);
    return 0;
}
```

### log6-09

```bash
bogon:CodeList6-09 panda8z$ gcc -o boolean.out boolean.c
bogon:CodeList6-09 panda8z$ ./boolean.out
Please enter an integer to be summed. (q to quit): 23
Please enter next integer (q to quit): 23
Please enter next integer (q to quit): 231
Please enter next integer (q to quit): q
Those integers sum to 277.

```

## CodeList 6-10 at P

### code6-10

```c
// sweetie1.c -- 一个计数循环
#include <stdio.h>
int main (void)
{
    const int NUMBER = 32;
    int count = 1;

    while(count <= NUMBER)
    {
        printf("Be my Valentine!\n");
        count++;
    }
    return 0;
}

```

### log6-10

```bash

bogon:CodeList6-10 panda8z$ gcc -o sweetie1.out sweetie1.c
bogon:CodeList6-10 panda8z$ ./sweetie1.out
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
Be my Valentine!
bogon:CodeList6-10 panda8z$
```

## CodeList 6-11 at P

### code6-11

```c
// sweetie2.c -- 一个使用for的计数循环
#include <stdio.h>
int main (void)
{
    const int NUMBER = 22;
    int count;

    for(count =1; count <= NUMBER; count++)
        printf("Be My Valentine!\n");
    return 0;
}

```

### log6-11

```bash

bogon:CodeList6-11 panda8z$ gcc -o sweetie2.out sweetie2.c
bogon:CodeList6-11 panda8z$ ./sweetie2.out
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
Be My Valentine!
```

## CodeList 6-12 at P

### code6-12

```c
/* for_cube.c -- 使用一个for循环产生一个立方表 */
#include <stdio.h>
int main(void)
{
    int num;

    printf("    n      n cubed\n");
    for(num = 1; num < 6; num++)
    {
        printf("%5d %5d\n", num, num * num * num);
    }
    return 0;
}

```

### log6-12

```bash
Panda-MBP:CodeList6-12 panda8z$ gcc -o for_cube.out for_cube.c
Panda-MBP:CodeList6-12 panda8z$ ./for_cube.out
    n      n cubed
    1     1
    2     8
    3    27
    4    64
    5   125

```

## CodeList 6-13 at P

### code6-13

```c
/* postage.c -- 一类邮件资费率 */
#include <stdio.h>
int main(void )
{
    const int FIRST_OZ = 37;
    const int NEXT_OZ = 23;
    int ounces, cost;
    printf(" ounces cost\n");
    for(ounces =1, cost = FIRST_OZ; ounces <= 16; ounces++, cost += NEXT_OZ)
    {
        printf("%5d $%4.2f\n", ounces, cost/100.0);
    }
    return 0;
}

```

### log6-13

```bash
Panda-MBP:CodeList6-13 panda8z$ gcc -o postage.out postage.c
Panda-MBP:CodeList6-13 panda8z$ ./postage.out
 ounces cost
    1 $0.37
    2 $0.60
    3 $0.83
    4 $1.06
    5 $1.29
    6 $1.52
    7 $1.75
    8 $1.98
    9 $2.21
   10 $2.44
   11 $2.67
   12 $2.90
   13 $3.13
   14 $3.36
   15 $3.59
   16 $3.82

```

## CodeList 6-14 at P

### code6-14

```c
// zeno.c -- 序列的和
#include <stdio.h>
int main (void)
{
    int t_ct;
    double time, x;
    int limit;

    printf("Enter the number of terms you want: ");
    scanf("%d", &limit);
    for(time = 0, x =1, t_ct = 1 ; t_ct <= limit; t_ct++, x *=20)
    {
        time += 1.0/x;
        printf("time = %f when terms = %d.\n", time, t_ct);
    }
    return 0;
}

```

### log6-14

```bash


Panda-MBP:CodeList6-14 panda8z$ gcc -o zeno.out zeno.c
Panda-MBP:CodeList6-14 panda8z$ ./zeno.out
Enter the number of terms you want: 88
time = 1.000000 when terms = 1.
time = 1.050000 when terms = 2.
time = 1.052500 when terms = 3.
time = 1.052625 when terms = 4.
time = 1.052631 when terms = 5.
time = 1.052632 when terms = 6.
time = 1.052632 when terms = 7.
time = 1.052632 when terms = 8.
time = 1.052632 when terms = 9.
time = 1.052632 when terms = 10.
time = 1.052632 when terms = 11.
time = 1.052632 when terms = 12.
time = 1.052632 when terms = 13.
time = 1.052632 when terms = 14.
time = 1.052632 when terms = 15.
time = 1.052632 when terms = 16.
time = 1.052632 when terms = 17.
time = 1.052632 when terms = 18.
time = 1.052632 when terms = 19.
time = 1.052632 when terms = 20.
time = 1.052632 when terms = 21.
time = 1.052632 when terms = 22.
time = 1.052632 when terms = 23.
time = 1.052632 when terms = 24.
time = 1.052632 when terms = 25.
time = 1.052632 when terms = 26.
time = 1.052632 when terms = 27.
time = 1.052632 when terms = 28.
time = 1.052632 when terms = 29.
time = 1.052632 when terms = 30.
time = 1.052632 when terms = 31.
time = 1.052632 when terms = 32.
time = 1.052632 when terms = 33.
time = 1.052632 when terms = 34.
time = 1.052632 when terms = 35.
time = 1.052632 when terms = 36.
time = 1.052632 when terms = 37.
time = 1.052632 when terms = 38.
time = 1.052632 when terms = 39.
time = 1.052632 when terms = 40.
time = 1.052632 when terms = 41.
time = 1.052632 when terms = 42.
time = 1.052632 when terms = 43.
time = 1.052632 when terms = 44.
time = 1.052632 when terms = 45.
time = 1.052632 when terms = 46.
time = 1.052632 when terms = 47.
time = 1.052632 when terms = 48.
time = 1.052632 when terms = 49.
time = 1.052632 when terms = 50.
time = 1.052632 when terms = 51.
time = 1.052632 when terms = 52.
time = 1.052632 when terms = 53.
time = 1.052632 when terms = 54.
time = 1.052632 when terms = 55.
time = 1.052632 when terms = 56.
time = 1.052632 when terms = 57.
time = 1.052632 when terms = 58.
time = 1.052632 when terms = 59.
time = 1.052632 when terms = 60.
time = 1.052632 when terms = 61.
time = 1.052632 when terms = 62.
time = 1.052632 when terms = 63.
time = 1.052632 when terms = 64.
time = 1.052632 when terms = 65.
time = 1.052632 when terms = 66.
time = 1.052632 when terms = 67.
time = 1.052632 when terms = 68.
time = 1.052632 when terms = 69.
time = 1.052632 when terms = 70.
time = 1.052632 when terms = 71.
time = 1.052632 when terms = 72.
time = 1.052632 when terms = 73.
time = 1.052632 when terms = 74.
time = 1.052632 when terms = 75.
time = 1.052632 when terms = 76.
time = 1.052632 when terms = 77.
time = 1.052632 when terms = 78.
time = 1.052632 when terms = 79.
time = 1.052632 when terms = 80.
time = 1.052632 when terms = 81.
time = 1.052632 when terms = 82.
time = 1.052632 when terms = 83.
time = 1.052632 when terms = 84.
time = 1.052632 when terms = 85.
time = 1.052632 when terms = 86.
time = 1.052632 when terms = 87.
time = 1.052632 when terms = 88.

```

## CodeList 6-15 at P

### code6-15

```c
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

```

### log6-15

```bash
Panda-MBP:CodeList6-15 panda8z$ gcc -o do_while.out do_while.c
Panda-MBP:CodeList6-15 panda8z$ ./do_while.out
To enter the triskaidekaphobia therapy club.
Please enter the secret code number:24
To enter the triskaidekaphobia therapy club.
Please enter the secret code number:12
To enter the triskaidekaphobia therapy club.
Please enter the secret code number:233
To enter the triskaidekaphobia therapy club.
Please enter the secret code number:2444
To enter the triskaidekaphobia therapy club.
Please enter the secret code number:13
Congratulations! you are cured!

```

## CodeList 6-16 at P

### code6-16

```c
/* entry.c -- 入口条件循环 */
#include <stdio.h>
int main(void)
{
    const int secret_code = 13;
    int code_entered;
    printf("To enter the triskaidekaphobia therapy club. \n");
    printf("Please enter the secret code number: ");
    scanf("%d", &code_entered);
    while (code_entered != secret_code)
    {
        printf("To enter the triskaidekaphobia therapy club. \n");
        printf("Please enter the secret code number: ");
        scanf("%d", &code_entered);
    }
    printf("congratulations! you are cured!\n");
    return 0;
}

```

### log6-16

```bash
Panda-MBP:CodeList6-16 panda8z$ gcc -o entry.out entry.c
Panda-MBP:CodeList6-16 panda8z$ ./entry.out
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 234
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 12
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 12
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 12
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 22
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 33
To enter the triskaidekaphobia therapy club.
Please enter the secret code number: 13
congratulations! you are cured!

```

## CodeList 6-17 at P

### code6-17

```c
/* rows1.c -- 使用嵌套循环 */
#include <stdio.h>
#define ROWS 6
#define CHARS 10
int main(void)
{
    int row;
    char ch;
    for (row = 0; row < ROWS; row++)
    {
        for (ch = 'A'; ch < ('A' + CHARS); ch++)
        {
            printf("%c", ch);
        }
        printf("\n");
    }
}

```

### log6-17

```bash
Panda-MBP:CodeList6-17 panda8z$ gcc -o rows1.out rows1.c
Panda-MBP:CodeList6-17 panda8z$ ./rows1.out
ABCDEFGHIJ
ABCDEFGHIJ
ABCDEFGHIJ
ABCDEFGHIJ
ABCDEFGHIJ
ABCDEFGHIJ

```

## CodeList 6-18 at P

### code6-18

```c
/* rows2.c -- 使用内部循环依赖于外部循环的嵌套循环 */
#include <stdio.h>
int main(void)
{
    const int ROWS = 6;
    const int CHARS = 6;

    int row;
    char ch;
    for(row = 0; row < ROWS; row++)
    {
        for(ch = ('A' + row); ch < ('A' + CHARS); ch++)
        {
            printf("%c", ch);
        }
        printf("\n");
    }
}

```

### log6-18

```bash
Panda-MBP:CodeList6-18 panda8z$ gcc -o rows2.out rows2.c
Panda-MBP:CodeList6-18 panda8z$ ./rows2.out
ABCDEF
BCDEF
CDEF
DEF
EF
F

```

## CodeList 6-19 at P

### code6-19

```c
/* scores_in.c -- 使用循环进行数组处理 */
#include <stdio.h>
#define SIZE 10
#define PAR 72
int main(void)
{
    int index, score[SIZE];
    int sum = 0;
    float average;

    printf("Enter %d golf scores: \n", SIZE);
    for (index = 0; index < SIZE; index++)
    {
        scanf("%d", &score[index]); //循环读入10 个分数
    }
    printf("the scores read in are as follows: \n");
    for (index = 0; index < SIZE; index++)
    {
        printf("%5d", score[index]);//验证输入
    }
    printf("\n");
    for (index = 0; index < SIZE; index++)
    {
        sum += score[index]; // 求他们的和
    }
    average = (float)sum / SIZE; // 节省时间的方法
    printf("Sum of scores = %d, average = %0.2f.\n", sum, average);
    printf("That`s a handicap of %.0f.]n", average - PAR);
    return 0;
}

```

### log6-19

```bash
Panda-MBP:CodeList6-19 panda8z$ gcc -o scores_in.out scores_in.c
Panda-MBP:CodeList6-19 panda8z$ ./scores_in.out
Enter 10 golf scores:
3
3
3
4
4
5
56
4
2

5
the scores read in are as follows:
    3    3    3    4    4    5   56    4    2    5
Sum of scores = 89, average = 8.90.
That`s a handicap of -63.

```

## CodeList 6-20 at P

### code6-

```c

/* power.c -- 计算数值的整数次幂 */
#include <stdio.h>
double power(double n, int p); //ANSI 原型
int main(void)
{
    double x, xpow;
    int exp;
    printf("Enter a number and the positive integer power");
    printf(" to which\nthe number will be raised. Enter q");
    printf(" to quit.\n");
    while (scanf("%lf%d", &x, &exp) == 2)
    {
        xpow = power(x, exp);
        printf("%.3g to the power %d is %.5g\n", x, exp, xpow);
        printf("Enter next pair of numbers or q to quit.\n");
    }
    printf("Hope you enjoyed ths power trip -- bye!\n");
    return 0;
}

double power(double n, int p)
{
    double pow = 1;
    int i;
    for (i = 1; i <= p; i++)
    {
        pow *= n;
    }
    return pow;
}
```

### log6-

```bash
Panda-MBP:CodeList6-20 panda8z$ gcc -o power.out power.c
Panda-MBP:CodeList6-20 panda8z$ ./power.out
Enter a number and the positive integer power to which
the number will be raised. Enter q to quit.
3
4
3 to the power 4 is 81
Enter next pair of numbers or q to quit.
23
23
23 to the power 23 is 2.088e+31
Enter next pair of numbers or q to quit.
q
Hope you enjoyed ths power trip -- bye!

```