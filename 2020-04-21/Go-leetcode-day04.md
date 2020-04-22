## 5. 最长回文子串

[5. 最长回文子串 - 力扣（LeetCode）](https://leetcode-cn.com/problems/longest-palindromic-substring/)

## 摘要

这篇文章是为中级读者而写的。它介绍了回文，动态规划以及字符串处理。请确保你理解什么是回文。回文是一个正读和反读都相同的字符串，例如，\textrm{“aba”}“aba” 是回文，而 \textrm{“abc”}“abc” 不是。

## 解决方案

### 方法一：最长公共子串

#### 常见错误

有些人会忍不住提出一个快速的解决方案，不幸的是，这个解决方案有缺陷(但是可以很容易地纠正)：

> 反转 $S$，使之变成 $S'$ 。找到 $S$ 和 $S'$ 之间最长的公共子串，这也必然是最长的回文子串。

这似乎是可行的，让我们看看下面的一些例子。

**例如**，

$S = \textrm{“caba”}S=“caba”, S' = \textrm{“abac”}$

S$′=“abac”$

$S$ 以及 $S'$ 之间的最长公共子串为$ \textrm{“aba”}$，恰恰是答案。

让我们尝试一下这个例子：$S = \textrm{“abacdfgdcaba”}, S' = \textrm{“abacdgfdcaba”}$：

$S$ 以及 $S'$  之间的最长公共子串为$ \textrm{“abacd”}$。显然，这不是回文。

#### 算法

我们可以看到，当 SS 的其他部分中存在非回文子串的反向副本时，最长公共子串法就会失败。为了纠正这一点，每当我们找到最长的公共子串的候选项时，都需要检查子串的索引是否与反向子串的原始索引相同。如果相同，那么我们尝试更新目前为止找到的最长回文子串；如果不是，我们就跳过这个候选项并继续寻找下一个候选。

这给我们提供了一个复杂度为$O(n^2)$动态规划解法，它将占用 $O(n^2)$ 的空间

可以改进为使用 $O(n)$ 的空间。请在 这里([最长公共子串_百度百科](https://baike.baidu.com/item/%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E4%B8%B2/22799982?fr=aladdin)) 阅读更多关于最长公共子串的内容。

### 方法二：暴力法

很明显，暴力法将选出所有子字符串可能的开始和结束位置，并检验它是不是回文。

#### 复杂度分析

* 时间复杂度：$O(n^3)$，假设 nn 是输入字符串的长度，则 $\binom{n}{2} = \frac{n(n-1)}{2}$为此类子字符串(不包括字符本身是回文的一般解法)的总数。因为验证每个子字符串需要 $O(n)$ 的时间，所以运行时间复杂度是 $O(n^3)$.

* 空间复杂度：$O(1)$。



```go

func longestPalindrome(s string) string {
    // 第一种思路：找出每一个子字符串，判断其是不是回文字符串
    length := len(s)
    if length == 0 {
        return ""
    }
    max := 0
    maxl := 0
    maxr := 0
    for i:=0;i<length;i++ {
        for j:=i+1; j<length; j++ {
            ret := isPalindromic(s, i, j)
            if ret {
                if max < (j-i+1) {
                    max = j - i + 1
                    maxl = i
                    maxr = j
                }
            }
        }
    }
    maxString := ""
    for i:=maxl; i<=maxr; i++ {
        maxString += string(s[i])
    }
    return maxString

    
}

func isPalindromic(s string, l, r int) bool {
    for l < r {
        if s[l] == s[r] {
            l++
            r--
        }else{
            return false
        }
       
    }
    return true
}

作者：wu-ming-shi-11
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/bao-li-po-jie-he-dong-tai-gui-hua-by-wu-ming-shi-1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```



### 方法三：动态规划

为了改进暴力法，我们首先观察如何避免在验证回文时进行不必要的重复计算。考虑 $\textrm{“ababa”}$这个示例。如果我们已经知道 $\textrm{“bab”}$ 是回文，那么很明显，$\textrm{“ababa”}$一定是回文，因为它的左首字母和右尾字母是相同的。

我们给出 $P(i,j)$ 的定义如下：

$$
P(i,j) = \begin{cases} \text{true,} &\quad\text{如果子串} S_i \dots S_j \text{是回文子串} \\ \text{false,} &\quad\text{其它情况} \end{cases}
$$



因此，

$P(i, j) = ( P(i+1, j-1) \text{ and } S_i == S_j )$

基本示例如下：

$P(i, i) = true$

$P(i, i+1) = ( S_i == S_{i+1} )$


这产生了一个直观的动态规划解法，我们首先初始化一字母和二字母的回文，然后找到所有三字母回文，并依此类推…

#### 复杂度分析

* 时间复杂度：$O(n^2)$，这里给出我们的运行时间复杂度为$ O(n^2)$=
* 空间复杂度：$O(n^2)$该方法使用$ O(n^2)$的空间来存储表。

#### 补充练习

你能进一步优化上述解法的空间复杂度吗？

### 方法四：中心扩展算法

事实上，只需使用恒定的空间，我们就可以在$ O(n^2)$的时间内解决这个问题。

我们观察到回文中心的两侧互为镜像。因此，回文可以从它的中心展开，并且只有 $2n-1$ 个这样的中心。

你可能会问，为什么会是$2n - 1$ 个，而不是 $n$ 个中心？

原因在于所含字母数为偶数的回文的中心可以处于两字母之间（例如$ \textrm{“abba”}$的中心在两个$ \textrm{‘b’}$之间）。

```Java
public String longestPalindrome(String s) {
    if (s == null || s.length() < 1) return "";
    int start = 0, end = 0;
    for (int i = 0; i < s.length(); i++) {
        int len1 = expandAroundCenter(s, i, i);
        int len2 = expandAroundCenter(s, i, i + 1);
        int len = Math.max(len1, len2);
        if (len > end - start) {
            start = i - (len - 1) / 2;
            end = i + len / 2;
        }
    }
    return s.substring(start, end + 1);
}
private int expandAroundCenter(String s, int left, int right) {
    int L = left, R = right;
    while (L >= 0 && R < s.length() && s.charAt(L) == s.charAt(R)) {
        L--;
        R++;
    }
    return R - L - 1;
}
```

```go
// 第二种思路：遍历字符串，找出以每个字符为中心的回文字符串有多长，选最长的返回
func longestPalindrome(s string) string {
    length := len(s)
    getLen := func(i, j int) int {
        // 以s[i]s[j]为中心的最长回文字符串
        for i>=0 && j<length {
            if s[i] == s[j] {
                i--
                j++
            }else{
                return j - i - 1
            }
        }
        return j - i - 1
    }
    max := 0
    maxStart := 0
    for i:=0;i<length;i++ {
      	temp := Max(getLen(i, i+1), getLen(i, i))
        if temp > max {
            max = temp
            maxStart = i - (max-1)/2
        }
    }
    maxString := ""
    for i:=maxStart; i<maxStart+max; i++ {
        maxString += string(s[i])
    }
    return maxString
}

func Max(i, j int) int {
    if i >= j {
        return i 
    }else{
        return j
    }
}

作者：wu-ming-shi-11
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/bao-li-po-jie-he-dong-tai-gui-hua-by-wu-ming-shi-1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```



#### 复杂度分析

* 时间复杂度：$O(n^2)$，由于围绕中心来扩展回文会耗去 $O(n)$ 的时间，所以总的复杂度为 $O(n^2)$。

* 空间复杂度：$O(1)$。

### 方法五：Manacher 算法

还有一个复杂度为 $O(n)$ 的 Manacher 算法。然而，这是一个非同寻常的算法，在 45 分钟的编码时间内提出这个算法将会是一个不折不扣的挑战。理解它，我保证这将是非常有趣的。

作者：LeetCode
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







[Manacher算法的详细讲解 - 简书](https://www.jianshu.com/p/116aa58b7d81)

```go

func longestPalindrome(s string) string {
    ls := len(s)
    if ls <= 1 {
        return s
    }

    mStr := make([]byte, ls*2+1)
    i, j := 0, 0
    for i < ls {
        mStr[j] = 0
        j++
        mStr[j] = s[i]
        j++
        i++
    }
    mStr[j] = 0
    ls = j + 1

    radius := make([]int, ls)
    r, c, max, mid := -1, -1, -1, 0
    for i = 0; i < ls; i++ {
        if r > i {
            if radius[2*c-i] > r-i+1 {
                radius[i] = r - i + 1
            } else {
                radius[i] = radius[2*c-i]
            }
        } else {
            radius[i] = 1
        }
        for i+radius[i] < ls && i-radius[i] > -1 {
            if mStr[i-radius[i]] == mStr[i+radius[i]] {
                radius[i]++
            } else {
                break
            }
        }
        if i+radius[i] > r {
            r = i + radius[i] - 1
            c = i
        }
        if max < radius[i] {
            max = radius[i]
            mid = i
        }
    }
    max--
    mid = mid/2 - max/2
    return s[mid : mid+max]
}


作者：janbar
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/manachershi-xian-by-janbar/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```

Manacher算法，又叫“马拉车”算法，可以在时间复杂度为O(n)的情况下求解一个字符串的最长回文子串长度的问题。

## 一、回文子串的一般解法

比较简单的思路是将字符串的每一个字符作为回文子串的中心对称点，每次保存前面求得的回文子串的最大值，最后得到的就是最长的回文子串的长度，这种方式的时间复杂度是O(n^2)。在求解过程中，基数的回文子串与偶数的回文子串是不一样的。比如最长回文子串为aba，对称中心就是b，如果最长回文子串为abba，则对称中心应该为两个b之间，为了解决这个问题，可以在每个字符两边加上一个符号，具体什么符号（是字符串里面的符号也行）对结果没有影响，比如加上“#”，则上述的两个序列变成了#a#b#a#和#a#b#b#a#，求出的长度分别为6和9，再除以2就可以得到最后的结果3和4。这种方式的时间复杂度太高，下面介绍时间复杂度为O(n)的Manacher算法。

## 二、Manacher算法中的基础概念

在进行Manacher算法时，字符串都会进行上面的进入一个字符处理，比如输入的字符为acbbcbds，用“#”字符处理之后的新字符串就是#a#c#b#b#c#b#d#s#。

### 1、回文半径数组radius

回文半径数组radius是用来记录以每个位置的字符为回文中心求出的回文半径长度，如下图所示，对于p1所指的位置radius[6]的回文半径是5，每个位置的回文半径组成的数组就是回文数组，所以#a#c#b#b#c#b#d#s#的回文半径数组为[1, 2, 1, 2, 1, 2, 5, 2, 1, 4, 1, 2, 1, 2, 1, 2, 1]。

![img](Go-leetcode-day04/01.webp)

要处理的字符串

### 2、最右回文右边界R

一个位置最右回文右边界指的是这个位置及之前的位置的回文子串，所到达的最右边的地方。比如对于字符串#a#c#b#b#c#b#d#s#，求它的每个位置的过程如下：

<img src="Go-leetcode-day04/image-20200422091151496.png" alt="image-20200422091151496" style="zoom:50%;" />

最右回文右边界R过程

最开始的时候R=-1，到p=0的位置，回文就是其本身，最右回文右边界R=0;p=1时，有回文串#a#，R=2；p=2时，R=2;P=3时，R=6;p=4时，最右回文右边界还是p=3时的右边界，R=6,依次类推。

### 3、最右回文右边界的对称中心C

就是上面提到的最右回文右边界的中心点C，如下图，p=4时，R=6，C=3

<img src="Go-leetcode-day04/image-20200422091231041.png" alt="image-20200422091231041" style="zoom:50%;" />

最右回文右边界的对称中心C

## 三、Manacher算法的流程

首先大的方面分为两种情况：

第一种情况：下一个要移动的位置在最右回文右边界R的右边。

比如在最开始时，R=-1,p的下一个移动的位置为p=0，p=0在R=-1的右边；p=0时，此时的R=0，p的下一个移动位置为p=1，也在R=0的右边。

在这种情况下，采用普遍的解法，将移动的位置为对称中心，向两边扩，同时更新回文半径数组，最右回文右边界R和最右回文右边界的对称中心C。

第二种情况：下一个要移动的位置就是最右回文右边界R或是在R的左边

在这种情况下又分为三种：

1、下一个要移动的位置p1**不在**最右回文右边界R右边，且cL<pL。

p2是p1以C为对称中心的对称点；

pL是以p2为对称中心的回文子串的左边界;

cL是以C为对称中心的回文子串的左边界。

这种情况下p1的回文半径就是p2的回文半径radius[p2]。

<img src="Go-leetcode-day04/image-20200422091316700.png" alt="image-20200422091316700" style="zoom:50%;" />

p1<=R且cL<pL

2、下一个要移动的位置票p1**不在**最右回文右边界R的右边，且cL>pL。

p2是p1以C为对称中心的对称点；

pL是以p2为对称中心的回文子串的左边界；

cL是以C为对称中心的回文子串的左边界。

这种情况下p1的回文半径就是p1到R的距离R-p1+1。

<img src="Go-leetcode-day04/image-20200422091333928.png" alt="image-20200422091333928" style="zoom:50%;" />

p1<=R且cL>pL

3、下一个要移动的位置票p1**不在**最右回文右边界R的右边，且cL=pL；

p2是p1以C为对称中心的对称点；

pL是以p2为对称中心的回文子串的左边界；

cL是以C为对称中心的回文子串的左边界。

这种情况下p1的回文半径就还要继续往外扩，但是只需要从R之后往外扩就可以了，扩了之后更新R和C。

<img src="Go-leetcode-day04/image-20200422091348986.png" alt="image-20200422091348986" style="zoom:50%;" />

## 四、Manacher时间复杂度分析

从上面的分析中，可以看出，第二种情况的1，2的求某个位置的回文半径的时间复杂度是O(1)，对于第一种情况和第二种情况的3，R是不断的向外扩的，不会往回退，而且寻找回文半径时，R之内的位置是不是进行判断的，所以对整个字符串而且，R的移动是从字符串的起点移动到终点，时间复杂度是O(n),所以整个manacher的时间复杂度是O(n)。

## 五、Manacher的代码实现

**java**

```csharp
public static void main(String[] args) {
        String str = "abcdcbafabcdck";
        //String str = "acbbcbds";
        System.out.println(manacher(str));
    }

    public static char[] manacherString(String str){
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < str.length(); i++) {
            sb.append("#");
            sb.append(str.charAt(i));
        }
        sb.append("#");
        return sb.toString().toCharArray();
    }

    public static int manacher(String str){
        if(str == null || str.length() < 1){
            return 0;
        }
        char[] charArr = manacherString(str);
        int[] radius = new int[charArr.length];
        int R = -1;
        int c = -1;
        int max = Integer.MIN_VALUE;
        for (int i = 0; i < radius.length; i++) {
            radius[i] = R > i ? Math.min(radius[2*c-i],R-i+1):1;
            while(i+radius[i] < charArr.length && i - radius[i] > -1){
                if(charArr[i-radius[i]] == charArr[i+radius[i]]){
                    radius[i]++;
                }else{
                    break;
                }
            }
            if(i + radius[i] > R){
                R = i + radius[i]-1;
                c = i;
            }
            max = Math.max(max,radius[i]);
        } 
        return max-1;
    }
```

作者：道禅_26ea
链接：https://www.jianshu.com/p/116aa58b7d81
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。