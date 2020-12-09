
# 5th Code List

## CodeList 5-1 at P

### code5-1

```c

/* shoes1.c -- 把一双鞋的尺码转换为英尺 */
#include <stdio.h>
#define ADJUST 7.64
#define SCALE 0.325
int main(void)
{
    double shoe, foot;
    shoe = 9.0;
    foot = SCALE * shoe + ADJUST;
    printf("Shoe size(men`s) foot length\n");
    printf("%10.1f %15.2f inches\n", shoe, foot);
    return 0;
}

```

### log5-1

```bash

Panda-MBP:CodeList5-1 panda8z$ gcc -o shoes1.out shoes1.c
Panda-MBP:CodeList5-1 panda8z$ ./shoes1.out
Shoe size(men`s) foot length
       9.0           10.56 inches
Panda-MBP:CodeList5-1 panda8z$

```

## CodeList 5-2 at P

### code5-2

```c

/* shoes2.c -- 计算多个写尺码对应的英寸长度 */
#include <stdio.h>
#define ADJUST 7.64
#define SCALE 0.325
int main(void)
{
    double shoe, foot;

    printf("Shoe size(men`s)     foot length\n");
    shoe = 3.0;
    while(shoe < 18.5)
    {
        foot = SCALE * shoe + ADJUST;
        printf("%10.1f %15.2f inches\n", shoe, foot);
        shoe = shoe + 1.0;
    }
    printf("If the shoe fits wear it,\n");
    return 0;
}

```

### log5-2

```bash

Panda-MBP:CodeList5-2 panda8z$ gcc -o shoes2.out shoes2.c
Panda-MBP:CodeList5-2 panda8z$ ./shoes2.out
Shoe size(men`s)     foot length
       3.0            8.62 inches
       4.0            8.94 inches
       5.0            9.27 inches
       6.0            9.59 inches
       7.0            9.91 inches
       8.0           10.24 inches
       9.0           10.56 inches
      10.0           10.89 inches
      11.0           11.21 inches
      12.0           11.54 inches
      13.0           11.87 inches
      14.0           12.19 inches
      15.0           12.52 inches
      16.0           12.84 inches
      17.0           13.16 inches
      18.0           13.49 inches
If the shoe fits wear it,
Panda-MBP:CodeList5-2 panda8z$

```

## CodeList 5-3 at P

### code5-3

```c
/* golf.c -- 高尔夫标赛记分卡 */
#include <stdio.h>
int main(void)
{
    int jane, tarzan, cheeta;

    cheeta = tarzan = jane = 68;
    printf("                 cheeta tarzan jane\n");
    printf("Fisrt round score %4d %8d %8d\n", cheeta, tarzan, jane);
    return 0;

}


```

### log5-3

```bash

Panda-MBP:CodeList5-3 panda8z$ gcc -o golf.out golf.c
Panda-MBP:CodeList5-3 panda8z$ ./golf.out
                 cheeta tarzan jane
Fisrt round score   68       68       68
Panda-MBP:CodeList5-3 panda8z$

```

## CodeList 5-4 at P

### code5-4

```c

/* squares.c -- 产生前20个整数的平方表 */
#include <stdio.h>
int main(void)
{
    int num = 1;
    while (num < 21)
    {
        printf("%4d %6d\n", num, num * num);
        num = num + 1;
    }
    return 0;
}

```

### log5-4

```bash

Panda-MBP:CodeList5-4 panda8z$ gcc -o squares.out squares.c
Panda-MBP:CodeList5-4 panda8z$ ./squares.out
   1      1
   2      4
   3      9
   4     16
   5     25
   6     36
   7     49
   8     64
   9     81
  10    100
  11    121
  12    144
  13    169
  14    196
  15    225
  16    256
  17    289
  18    324
  19    361
  20    400
Panda-MBP:CodeList5-4 panda8z$

```

## CodeList 5-5 at P

### code5-5

```c
/* wheat.c --  指数增长 */
#include <stdio.h>
#define SQUARES 64 // 棋盘上的方格数
#define CROP 1E15  //以粒记得美国小麦产量
int main(void)
{
    double current, total;
    int count = 1;

    printf("suqare grains total  ");
    printf("fraction of \n");
    printf("   added     grain    ");
    printf("US total\n");
    total = current = 1.0;
    printf("%4d %13.2e %12.2e %12.2e\n",
           count, current, total, total / CROP);

    while (count < SQUARES)
    {
        count = count + 1;
        current = 2.0 * current; //下个方格粒数加倍
        total = total + current; //更新总数
        printf("%4d %13.2e %12.2e %12.2e\n",
               count, current, total, total / CROP);
    }
    printf("That`s all.\n");
    return 0;
}

```

### log5-5

```bash
Panda-MBP:CodeList5-5 panda8z$ gcc -o wheat.out wheat.c
Panda-MBP:CodeList5-5 panda8z$ ./wheat.out
suqare grains total  fraction of
   added     grain    US total
   1      1.00e+00     1.00e+00     1.00e-15
   2      2.00e+00     3.00e+00     3.00e-15
   3      4.00e+00     7.00e+00     7.00e-15
   4      8.00e+00     1.50e+01     1.50e-14
   5      1.60e+01     3.10e+01     3.10e-14
   6      3.20e+01     6.30e+01     6.30e-14
   7      6.40e+01     1.27e+02     1.27e-13
   8      1.28e+02     2.55e+02     2.55e-13
   9      2.56e+02     5.11e+02     5.11e-13
  10      5.12e+02     1.02e+03     1.02e-12
  11      1.02e+03     2.05e+03     2.05e-12
  12      2.05e+03     4.10e+03     4.09e-12
  13      4.10e+03     8.19e+03     8.19e-12
  14      8.19e+03     1.64e+04     1.64e-11
  15      1.64e+04     3.28e+04     3.28e-11
  16      3.28e+04     6.55e+04     6.55e-11
  17      6.55e+04     1.31e+05     1.31e-10
  18      1.31e+05     2.62e+05     2.62e-10
  19      2.62e+05     5.24e+05     5.24e-10
  20      5.24e+05     1.05e+06     1.05e-09
  21      1.05e+06     2.10e+06     2.10e-09
  22      2.10e+06     4.19e+06     4.19e-09
  23      4.19e+06     8.39e+06     8.39e-09
  24      8.39e+06     1.68e+07     1.68e-08
  25      1.68e+07     3.36e+07     3.36e-08
  26      3.36e+07     6.71e+07     6.71e-08
  27      6.71e+07     1.34e+08     1.34e-07
  28      1.34e+08     2.68e+08     2.68e-07
  29      2.68e+08     5.37e+08     5.37e-07
  30      5.37e+08     1.07e+09     1.07e-06
  31      1.07e+09     2.15e+09     2.15e-06
  32      2.15e+09     4.29e+09     4.29e-06
  33      4.29e+09     8.59e+09     8.59e-06
  34      8.59e+09     1.72e+10     1.72e-05
  35      1.72e+10     3.44e+10     3.44e-05
  36      3.44e+10     6.87e+10     6.87e-05
  37      6.87e+10     1.37e+11     1.37e-04
  38      1.37e+11     2.75e+11     2.75e-04
  39      2.75e+11     5.50e+11     5.50e-04
  40      5.50e+11     1.10e+12     1.10e-03
  41      1.10e+12     2.20e+12     2.20e-03
  42      2.20e+12     4.40e+12     4.40e-03
  43      4.40e+12     8.80e+12     8.80e-03
  44      8.80e+12     1.76e+13     1.76e-02
  45      1.76e+13     3.52e+13     3.52e-02
  46      3.52e+13     7.04e+13     7.04e-02
  47      7.04e+13     1.41e+14     1.41e-01
  48      1.41e+14     2.81e+14     2.81e-01
  49      2.81e+14     5.63e+14     5.63e-01
  50      5.63e+14     1.13e+15     1.13e+00
  51      1.13e+15     2.25e+15     2.25e+00
  52      2.25e+15     4.50e+15     4.50e+00
  53      4.50e+15     9.01e+15     9.01e+00
  54      9.01e+15     1.80e+16     1.80e+01
  55      1.80e+16     3.60e+16     3.60e+01
  56      3.60e+16     7.21e+16     7.21e+01
  57      7.21e+16     1.44e+17     1.44e+02
  58      1.44e+17     2.88e+17     2.88e+02
  59      2.88e+17     5.76e+17     5.76e+02
  60      5.76e+17     1.15e+18     1.15e+03
  61      1.15e+18     2.31e+18     2.31e+03
  62      2.31e+18     4.61e+18     4.61e+03
  63      4.61e+18     9.22e+18     9.22e+03
  64      9.22e+18     1.84e+19     1.84e+04
That`s all.
Panda-MBP:CodeList5-5 panda8z$
```

## CodeList 5-6 at P

### code5-6

```c

/*  divide.c -- 我们所知的除法 */
#include <stdio.h>
int main(void)
{
    printf("integer division: 5/4 is %d \n", 5/4);
    printf("integer division: 6/3 is %d \n", 6/3);
    printf("integer division: 7/4 is %d \n", 7/4);
    printf("floating division: 7./4 is %1.2f \n", 7./4);
    printf("mixed division: 7./4 is %1.2f \n", 7./4);
    return 0;
}

```

### log5-6

```bash

Panda-MBP:CodeList5-6 panda8z$ gcc -o divide.out divide.c
Panda-MBP:CodeList5-6 panda8z$ ./divide.out
integer division: 5/4 is 1
integer division: 6/3 is 2
integer division: 7/4 is 1
floating division: 7./4 is 1.75
mixed division: 7./4 is 1.75
Panda-MBP:CodeList5-6 panda8z$

```

## CodeList 5-7 at P

### code5-7

```c

/* rules.c -- 优先级规则的试验 */
#include <stdio.h>
int main(void)
{
    int top, score;
    top = score = -(2 + 5) * 6 + (4 + 3 * (2 + 3));
    // 先预演一下,从左到右, - 7 * 6 + ( 4 + 3 * 5)
    // - 42 + (4 + 15)
    // - 42 + 19
    // - 23

    // 第一次预演 失败!

    // 第二次开始
    // 第一次翻了小学数学的错误,弱智.
    printf("top = %d \n", top);
    return 0;
}

```

### log5-7

```bash
Panda-MBP:CodeList5-7 panda8z$ gcc -o rules.out rules.c
Panda-MBP:CodeList5-7 panda8z$ ./rules.out
top = -23
Panda-MBP:CodeList5-7 panda8z$

```

## CodeList 5-8 at P

### code5-8

```c

/* sizeof.c -- 使用sizeof运算符 */
// 使用C99的%z修饰符.如果不能使用%zd,请使用%u或%1u
#include <stdio.h>
int main(void)
{
    int n = 0;
    size_t  intsize;
    intsize = sizeof(int);// C规定sizeof返回size_t类型的值
    printf("n = %d, n has %zd bytes; all ints have %zd bytes.\n",
    n, sizeof n, intsize);
    return 0;

}

```

### log5-8

```bash


Panda-MBP:CodeList5-8 panda8z$ gcc -o sizeof.out sizeof.c
Panda-MBP:CodeList5-8 panda8z$ ./sizeof.out
n = 0, n has 4 bytes; all ints have 4 bytes.
Panda-MBP:CodeList5-8 panda8z$```

## CodeList 5-9 at P

### code5-9

```c
/* min_sec.c -- 把秒变成分钟和秒 */
#include <stdio.h>
#define SEC_PER_MIN 60
int main(void)
{
    int sec, min, left;

    printf("Convert seconds to minutes and seconds!\n");
    printf("Enter the number of seconds(<= 0 to quit):\n");
    scanf("%d", &sec);
    while (sec > 0)
    {
        min = sec / SEC_PER_MIN;
        left = sec % SEC_PER_MIN;
        printf("%d seconds is %d minutes, %d seconds.\n", sec, min, left);
        printf("Enter next value (<=0 to quit);\n");
        scanf("%d", &sec);
    }
    printf("Done!\n");
    return 0;
}

```

### log5-9

```bash

Panda-MBP:CodeList5-9 panda8z$ gcc -o min_sec.out min_sec.c
Panda-MBP:CodeList5-9 panda8z$ ./min_sec.out
Convert seconds to minutes and seconds!
Enter the number of seconds(<= 0 to quit):
2345
2345 seconds is 39 minutes, 5 seconds.
Enter next value (<=0 to quit);
234567
234567 seconds is 3909 minutes, 27 seconds.
Enter next value (<=0 to quit);
123
123 seconds is 2 minutes, 3 seconds.
Enter next value (<=0 to quit);
123
123 seconds is 2 minutes, 3 seconds.
Enter next value (<=0 to quit);
-111
Done!
Panda-MBP:CodeList5-9 panda8z$
```

## CodeList 5-10 at P

### code5-10

```c

/* add_one.c -- 增量: 前缀后缀 */
#include <stdio.h>
int main (void)
{
    int ultra = 0, super = 0;
    while(super < 5)
    {
        super++;
        ++ultra;
        printf("super = %d, ultra = %d \n",super, ultra);
    }
    return 0;
}

```

### log5-10

```bash

Panda-MBP:CodeList5-10 panda8z$
Panda-MBP:CodeList5-10 panda8z$ gcc -o add_one.out add_one.c
Panda-MBP:CodeList5-10 panda8z$ ./add_one.out
super = 1, ultra = 1
super = 2, ultra = 2
super = 3, ultra = 3
super = 4, ultra = 4
super = 5, ultra = 5
Panda-MBP:CodeList5-10 panda8z$
```

## CodeList 5-11 at P

### code5-11

```c

/* post_pre.c -- 后缀和前缀 */
#include <stdio.h>
int main(void)
{
    int a = 1, b = 1;
    int aplus, plusb;

    aplus = a++;
    plusb = ++b;
    printf("a   aplus   b   plusb \n");
    printf("%1d %5d %5d %5d\n", a, aplus, b, plusb);
    return 0;
}

```

### log5-11

```bash

Panda-MBP:CodeList5-11 panda8z$ gcc -o post_pre.out post_pre.c
Panda-MBP:CodeList5-11 panda8z$ ./post_pre.out
a   aplus   b   plusb
2     1     2     2
Panda-MBP:CodeList5-11 panda8z$
```

## CodeList 5-12 at P

### code5-12

```c

/* bottles.c --  */
#include <stdio.h>
#define MAX 100
int main(void)
{
    int count = MAX + 1;

    while (--count > 0)
    {
        printf("%d bottles of spring water on the wall,"
               "%d bottles of spring water!\n",
               count, count);
        printf("Take one down and pass it around,\n");
        printf("%d bottles of spring water!\n\n", count - 1);
    }
    return 0;
}

```

### log5-12

```bash
Panda-MBP:CodeList5-12 panda8z$ ./bottles.out
100 bottles of spring water on the wall,100 bottles of spring water!
Take one down and pass it around,
99 bottles of spring water!

99 bottles of spring water on the wall,99 bottles of spring water!
Take one down and pass it around,
98 bottles of spring water!

98 bottles of spring water on the wall,98 bottles of spring water!
Take one down and pass it around,
97 bottles of spring water!

97 bottles of spring water on the wall,97 bottles of spring water!
Take one down and pass it around,
96 bottles of spring water!

96 bottles of spring water on the wall,96 bottles of spring water!
Take one down and pass it around,
95 bottles of spring water!

95 bottles of spring water on the wall,95 bottles of spring water!
Take one down and pass it around,
94 bottles of spring water!

94 bottles of spring water on the wall,94 bottles of spring water!
Take one down and pass it around,
93 bottles of spring water!

93 bottles of spring water on the wall,93 bottles of spring water!
Take one down and pass it around,
92 bottles of spring water!

92 bottles of spring water on the wall,92 bottles of spring water!
Take one down and pass it around,
91 bottles of spring water!

91 bottles of spring water on the wall,91 bottles of spring water!
Take one down and pass it around,
90 bottles of spring water!

90 bottles of spring water on the wall,90 bottles of spring water!
Take one down and pass it around,
89 bottles of spring water!

89 bottles of spring water on the wall,89 bottles of spring water!
Take one down and pass it around,
88 bottles of spring water!

88 bottles of spring water on the wall,88 bottles of spring water!
Take one down and pass it around,
87 bottles of spring water!

87 bottles of spring water on the wall,87 bottles of spring water!
Take one down and pass it around,
86 bottles of spring water!

86 bottles of spring water on the wall,86 bottles of spring water!
Take one down and pass it around,
85 bottles of spring water!

85 bottles of spring water on the wall,85 bottles of spring water!
Take one down and pass it around,
84 bottles of spring water!

84 bottles of spring water on the wall,84 bottles of spring water!
Take one down and pass it around,
83 bottles of spring water!

83 bottles of spring water on the wall,83 bottles of spring water!
Take one down and pass it around,
82 bottles of spring water!

82 bottles of spring water on the wall,82 bottles of spring water!
Take one down and pass it around,
81 bottles of spring water!

81 bottles of spring water on the wall,81 bottles of spring water!
Take one down and pass it around,
80 bottles of spring water!

80 bottles of spring water on the wall,80 bottles of spring water!
Take one down and pass it around,
79 bottles of spring water!

79 bottles of spring water on the wall,79 bottles of spring water!
Take one down and pass it around,
78 bottles of spring water!

78 bottles of spring water on the wall,78 bottles of spring water!
Take one down and pass it around,
77 bottles of spring water!

77 bottles of spring water on the wall,77 bottles of spring water!
Take one down and pass it around,
76 bottles of spring water!

76 bottles of spring water on the wall,76 bottles of spring water!
Take one down and pass it around,
75 bottles of spring water!

75 bottles of spring water on the wall,75 bottles of spring water!
Take one down and pass it around,
74 bottles of spring water!

74 bottles of spring water on the wall,74 bottles of spring water!
Take one down and pass it around,
73 bottles of spring water!

73 bottles of spring water on the wall,73 bottles of spring water!
Take one down and pass it around,
72 bottles of spring water!

72 bottles of spring water on the wall,72 bottles of spring water!
Take one down and pass it around,
71 bottles of spring water!

71 bottles of spring water on the wall,71 bottles of spring water!
Take one down and pass it around,
70 bottles of spring water!

70 bottles of spring water on the wall,70 bottles of spring water!
Take one down and pass it around,
69 bottles of spring water!

69 bottles of spring water on the wall,69 bottles of spring water!
Take one down and pass it around,
68 bottles of spring water!

68 bottles of spring water on the wall,68 bottles of spring water!
Take one down and pass it around,
67 bottles of spring water!

67 bottles of spring water on the wall,67 bottles of spring water!
Take one down and pass it around,
66 bottles of spring water!

66 bottles of spring water on the wall,66 bottles of spring water!
Take one down and pass it around,
65 bottles of spring water!

65 bottles of spring water on the wall,65 bottles of spring water!
Take one down and pass it around,
64 bottles of spring water!

64 bottles of spring water on the wall,64 bottles of spring water!
Take one down and pass it around,
63 bottles of spring water!

63 bottles of spring water on the wall,63 bottles of spring water!
Take one down and pass it around,
62 bottles of spring water!

62 bottles of spring water on the wall,62 bottles of spring water!
Take one down and pass it around,
61 bottles of spring water!

61 bottles of spring water on the wall,61 bottles of spring water!
Take one down and pass it around,
60 bottles of spring water!

60 bottles of spring water on the wall,60 bottles of spring water!
Take one down and pass it around,
59 bottles of spring water!

59 bottles of spring water on the wall,59 bottles of spring water!
Take one down and pass it around,
58 bottles of spring water!

58 bottles of spring water on the wall,58 bottles of spring water!
Take one down and pass it around,
57 bottles of spring water!

57 bottles of spring water on the wall,57 bottles of spring water!
Take one down and pass it around,
56 bottles of spring water!

56 bottles of spring water on the wall,56 bottles of spring water!
Take one down and pass it around,
55 bottles of spring water!

55 bottles of spring water on the wall,55 bottles of spring water!
Take one down and pass it around,
54 bottles of spring water!

54 bottles of spring water on the wall,54 bottles of spring water!
Take one down and pass it around,
53 bottles of spring water!

53 bottles of spring water on the wall,53 bottles of spring water!
Take one down and pass it around,
52 bottles of spring water!

52 bottles of spring water on the wall,52 bottles of spring water!
Take one down and pass it around,
51 bottles of spring water!

51 bottles of spring water on the wall,51 bottles of spring water!
Take one down and pass it around,
50 bottles of spring water!

50 bottles of spring water on the wall,50 bottles of spring water!
Take one down and pass it around,
49 bottles of spring water!

49 bottles of spring water on the wall,49 bottles of spring water!
Take one down and pass it around,
48 bottles of spring water!

48 bottles of spring water on the wall,48 bottles of spring water!
Take one down and pass it around,
47 bottles of spring water!

47 bottles of spring water on the wall,47 bottles of spring water!
Take one down and pass it around,
46 bottles of spring water!

46 bottles of spring water on the wall,46 bottles of spring water!
Take one down and pass it around,
45 bottles of spring water!

45 bottles of spring water on the wall,45 bottles of spring water!
Take one down and pass it around,
44 bottles of spring water!

44 bottles of spring water on the wall,44 bottles of spring water!
Take one down and pass it around,
43 bottles of spring water!

43 bottles of spring water on the wall,43 bottles of spring water!
Take one down and pass it around,
42 bottles of spring water!

42 bottles of spring water on the wall,42 bottles of spring water!
Take one down and pass it around,
41 bottles of spring water!

41 bottles of spring water on the wall,41 bottles of spring water!
Take one down and pass it around,
40 bottles of spring water!

40 bottles of spring water on the wall,40 bottles of spring water!
Take one down and pass it around,
39 bottles of spring water!

39 bottles of spring water on the wall,39 bottles of spring water!
Take one down and pass it around,
38 bottles of spring water!

38 bottles of spring water on the wall,38 bottles of spring water!
Take one down and pass it around,
37 bottles of spring water!

37 bottles of spring water on the wall,37 bottles of spring water!
Take one down and pass it around,
36 bottles of spring water!

36 bottles of spring water on the wall,36 bottles of spring water!
Take one down and pass it around,
35 bottles of spring water!

35 bottles of spring water on the wall,35 bottles of spring water!
Take one down and pass it around,
34 bottles of spring water!

34 bottles of spring water on the wall,34 bottles of spring water!
Take one down and pass it around,
33 bottles of spring water!

33 bottles of spring water on the wall,33 bottles of spring water!
Take one down and pass it around,
32 bottles of spring water!

32 bottles of spring water on the wall,32 bottles of spring water!
Take one down and pass it around,
31 bottles of spring water!

31 bottles of spring water on the wall,31 bottles of spring water!
Take one down and pass it around,
30 bottles of spring water!

30 bottles of spring water on the wall,30 bottles of spring water!
Take one down and pass it around,
29 bottles of spring water!

29 bottles of spring water on the wall,29 bottles of spring water!
Take one down and pass it around,
28 bottles of spring water!

28 bottles of spring water on the wall,28 bottles of spring water!
Take one down and pass it around,
27 bottles of spring water!

27 bottles of spring water on the wall,27 bottles of spring water!
Take one down and pass it around,
26 bottles of spring water!

26 bottles of spring water on the wall,26 bottles of spring water!
Take one down and pass it around,
25 bottles of spring water!

25 bottles of spring water on the wall,25 bottles of spring water!
Take one down and pass it around,
24 bottles of spring water!

24 bottles of spring water on the wall,24 bottles of spring water!
Take one down and pass it around,
23 bottles of spring water!

23 bottles of spring water on the wall,23 bottles of spring water!
Take one down and pass it around,
22 bottles of spring water!

22 bottles of spring water on the wall,22 bottles of spring water!
Take one down and pass it around,
21 bottles of spring water!

21 bottles of spring water on the wall,21 bottles of spring water!
Take one down and pass it around,
20 bottles of spring water!

20 bottles of spring water on the wall,20 bottles of spring water!
Take one down and pass it around,
19 bottles of spring water!

19 bottles of spring water on the wall,19 bottles of spring water!
Take one down and pass it around,
18 bottles of spring water!

18 bottles of spring water on the wall,18 bottles of spring water!
Take one down and pass it around,
17 bottles of spring water!

17 bottles of spring water on the wall,17 bottles of spring water!
Take one down and pass it around,
16 bottles of spring water!

16 bottles of spring water on the wall,16 bottles of spring water!
Take one down and pass it around,
15 bottles of spring water!

15 bottles of spring water on the wall,15 bottles of spring water!
Take one down and pass it around,
14 bottles of spring water!

14 bottles of spring water on the wall,14 bottles of spring water!
Take one down and pass it around,
13 bottles of spring water!

13 bottles of spring water on the wall,13 bottles of spring water!
Take one down and pass it around,
12 bottles of spring water!

12 bottles of spring water on the wall,12 bottles of spring water!
Take one down and pass it around,
11 bottles of spring water!

11 bottles of spring water on the wall,11 bottles of spring water!
Take one down and pass it around,
10 bottles of spring water!

10 bottles of spring water on the wall,10 bottles of spring water!
Take one down and pass it around,
9 bottles of spring water!

9 bottles of spring water on the wall,9 bottles of spring water!
Take one down and pass it around,
8 bottles of spring water!

8 bottles of spring water on the wall,8 bottles of spring water!
Take one down and pass it around,
7 bottles of spring water!

7 bottles of spring water on the wall,7 bottles of spring water!
Take one down and pass it around,
6 bottles of spring water!

6 bottles of spring water on the wall,6 bottles of spring water!
Take one down and pass it around,
5 bottles of spring water!

5 bottles of spring water on the wall,5 bottles of spring water!
Take one down and pass it around,
4 bottles of spring water!

4 bottles of spring water on the wall,4 bottles of spring water!
Take one down and pass it around,
3 bottles of spring water!

3 bottles of spring water on the wall,3 bottles of spring water!
Take one down and pass it around,
2 bottles of spring water!

2 bottles of spring water on the wall,2 bottles of spring water!
Take one down and pass it around,
1 bottles of spring water!

1 bottles of spring water on the wall,1 bottles of spring water!
Take one down and pass it around,
0 bottles of spring water!

Panda-MBP:CodeList5-12 panda8z$

```

## CodeList 5-13 at P

### code5-13

```c

/* addemup.c -- 4种类型的语句 */
#include <stdio.h>
int main(void) /* 求出前20个数的和 */
{
    int count, sum;      /* 声明语句 */
    count = 0;           /* 赋值语句 */
    sum = 0;             //同上
    while (count++ < 20) //while
    {
        sum = sum + count; //语句
    }
    printf("sum = %d\n", sum); //函数语句
    return 0;
}

```

### log5-13

```bash

Panda-MBP:CodeList5-13 panda8z$ ./addemup.out
sum = 210
Panda-MBP:CodeList5-13 panda8z$

```

## CodeList 5-14 at P

### code5-14

```c

/* convert.c -- 自动类型转换 */
#include <stdio.h>
int main(void)
{
    char ch;
    int i;
    float fl;

    fl = i = ch = 'C';

    printf("ch = %c, i = %d, fl = %2.2f\n", ch, i, fl);
    ch = ch + 1;
    i = fl + 2 * ch;
    fl = 2.0 * ch + i;
    printf("ch = %c, i = %d, fl = %2.2f\n", ch, i, fl);
    ch = 5212205.17;
    printf("Now ch = %c\n", ch);
    return 0;
}

```

### log5-14

```bash


Panda-MBP:CodeList5-14 panda8z$ ./convert.out
ch = C, i = 67, fl = 67.00
ch = D, i = 203, fl = 339.00
Now ch = �
Panda-MBP:CodeList5-14 panda8z$
```

## CodeList 5-15 at P

### code5-15

```c

/* pound.c -- 定义带有一个参数的函数 */
#include <stdio.h>
void pound(int n);/* NASI风格的原型 */
int main(void)
{
    int times = 5;
    char ch = '!';//ASCII码值为33
    float f = 6.0;
    pound(times);//int 参数
    pound(ch);//char参数自动转换为int类型
    pound((int)f); //指派运算符把f强制转换为int类型
    return 0;
}

void pound(int n) //ANSI风格的函数头,
                // 说明该函数接受一个int参数
{
    while( n-- > 0)
    {
        printf("#");
    }
    printf("\n");
}

```

### log5-15

```bash

Panda-MBP:CodeList5-15 panda8z$ gcc -o pound.out pound.c
Panda-MBP:CodeList5-15 panda8z$ ./pound.out#####
#################################
######
Panda-MBP:CodeList5-15 panda8z$
```

## CodeList 5-16 at P

### code5-16

```c


/* running.c -- 一个对于长跑运动员有用的程序 */
#include <stdio.h>
const int S_PER_M = 60;         //每分钟的秒数
const int S_PER_H = 3600;       //每小时的秒数
const double M_PER_K = 0.62137; //每公里的英里数
int main(void)
{
    double distk, distm; //分别以公里和英里计的跑过的距离
    double rate;         //以英里/小时为单位的平均速度
    int min, sec;        //跑步用的分钟数和秒数
    int time;            //用秒表示跑步用时
    double mtime;        //跑完一英里所用的时间,以秒计
    int mmin, msec;      //跑完一英里所用的时间,以分钟和秒计
    printf("Thes program converts your time for a metric race\n");
    printf("to a time for running a mile and to your averange\n");
    printf("speed in miles per hour.\n");
    printf("Please enter, in kilometers, the distance run.\n");
    scanf("%1lf", &distk); //%1f表示读取一个double类型的数值
    printf("Next Enter the time in minutes and seconds.\n");
    printf("Begin by entering the minutes.\n");
    scanf("%d", &min);
    printf("Now enter the seconds.\n");
    scanf("%d", &sec);
    //把时间转换为全部用秒表示
    time = S_PER_M * min + sec;
    //把公里转换为英里
    distm = M_PER_K * distk;
    //英里/秒 X 秒/小时 = 英里/小时
    rate = distm / time * S_PER_H;
    //时间/距离 = 跑完每英里的用时
    mtime = (double)time / distm;
    mmin = (int)mtime / S_PER_M;//求出分钟数
    msec = (int)mtime % S_PER_M;//求出剩余秒数
    printf("You ran %1.2f km(%1.2f miles) in %d min, %d sec.\n",    distk, distm, min, sec);
    printf("That pace corresponds to running a mile in %d min.",
    mmin);
    printf("%d sec.\n Your averange speed was %1.2f mph.\n", msec, rate);
    return 0;
}
```

### log5-16

```bash
Panda-MBP:CodeList5-16 panda8z$ gcc -o running.out running.c
running.c:18:18: warning: format specifies type 'float *' but the argument has type 'double *' [-Wformat]
    scanf("%1f", &distk); //%1f表示读取一个double类型的数值
           ~~~   ^~~~~~
           %1lf
1 warning generated.
Panda-MBP:CodeList5-16 panda8z$ gcc -o running.out running.c
Panda-MBP:CodeList5-16 panda8z$ ./running.out
Thes program converts your time for a metric race
to a time for running a mile and to your averange
speed in miles per hour.
Please enter, in kilometers, the distance run.
5
Next Enter the time in minutes and seconds.
Begin by entering the minutes.
32
Now enter the seconds.
10
You ran 5.00 km(3.11 miles) in 32 min, 10 sec.
That pace corresponds to running a mile in 10 min.21 sec.
 Your averange speed was 5.80 mph.
Panda-MBP:CodeList5-16 panda8z$

```
