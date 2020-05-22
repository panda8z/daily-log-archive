# LeetCode-Day03

> 接下来的半个月时间，每天都会在这个文件里更新leetcode复习和做题记录
> 初步确定每天完成按顺序的20个题目。

## 1. 两数之和

[1. 两数之和 - 力扣（LeetCode）](https://leetcode-cn.com/problems/two-sum/)

代码请看demo01

```go
// 1. 两数之和

package main

import "fmt"

func addTwo(nums []int, target int) []int {
	//1. 创建map来存 【值，索引】
	m := map[int]int{}
	// 2. 遍历每一个数
	for i, v := range nums {
		// 3. 根据map的特性是否能找到map里有相应的值
		if j, ok := m[target-v]; ok {
			// 如果存在就 返回存储的值的索引 j和当前的索引 i组成的切片。
			return []int{j, i}
		}
		// 4. 如果没有就 按[值，索引]存进map
		m[v] = i
	}
	return nil
}
func main() {
	fmt.Println(addTwo([]int{1, 2, 5, 4}, 5))
	fmt.Println(addTwo([]int{1, 2, 5, 4}, 7))
}


```

## 2. 两数相加

[2. 两数相加 - 力扣（LeetCode）](https://leetcode-cn.com/problems/add-two-numbers/)



代码请看：demo02

```go

//[2. 两数相加 - 力扣（LeetCode）](https://leetcode-cn.com/problems/add-two-numbers/)
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//  }

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 返回值是一个链表节点,所以定义一个返回值的头
	result := &ListNode{Val: 0}

	nowNode := result

	// 2. 处理进位和加法结果
	var i, s int

	for {
		if i > 0 {
			s = l1.Val + l2.Val + 1
		} else {
			s = l1.Val + l2.Val
		}

		if s >= 10 {
			nowNode.Next = &ListNode{Val: s - 10}
			i = 1
		} else {
			nowNode.Next = &ListNode{Val: s}
			i = 0
		}

		nowNode = nowNode.Next

		if l1.Next == nil && l2.Next == nil {
			if i > 1 {
				nowNode.Next = &ListNode{Val: i}
			}
			break
		}

		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}

		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}

	return result.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2).Val)
	fmt.Println(addTwoNumbers(l1, l2).Next.Val)
	fmt.Println(addTwoNumbers(l1, l2).Next.Next.Val)
	fmt.Println(addTwoNumbers(l1, l2).Next.Next.Next.Val)
}

```



## 3. 无重复字符的最长子串

[3. 无重复字符的最长子串 - 力扣（LeetCode）](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

执行结果：

通过

执行用时 :4 ms, 在所有 Go 提交中击败了88.86%的用户

内存消耗 :3.6 MB, 在所有 Go 提交中击败了29.03%的用户

代码请看：demo03

```go

// [3. 无重复字符的最长子串 - 力扣（LeetCode）](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 1. 定义用于存放已遍历字符的map
	m := map[rune]int{}
	// 2. 定义起点 索引
	startI := 0
	// 3. 定义最大长度
	maxLength := 0

	// 4. 按序遍历字符串中所有字符
	for i, v := range []rune(s) {
		// 5. 记录是否出现过字符的index
		if lastI, ok := m[v]; ok && lastI >= startI {
			// 如果出现过就将开始位置往后挪一个
			startI = lastI + 1
		}

		// 6. 计算长度
		if i-startI+1 >= maxLength {
			maxLength = i-startI+1
		}

		// 7. 记录遍历过的字符 【字符，index】
		m[v] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
```



## [4. 寻找两个有序数组的中位数 - 力扣（LeetCode）](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

<img src="go-leetcode-day03/image-20200419000314227.png" alt="image-20200419000314227" style="zoom:50%;" />

```go

package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	iMin := 0
	iMax := m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i != m && nums2[j-1] > nums1[i] {
			// i需要增大
			iMin = i + 1
		} else if i != 0 && j != n && nums1[i-1] > nums2[j] {
			// i需要减小
			iMax = i - 1
		} else {
			// 达到条件，边界值列出来单独考虑
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				// 奇数不需要考虑右半部分
				return maxLeft
			}

			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums2[j]), float64(nums1[i]))
			}
			// 如果是偶数直接返回
			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)

	fmt.Println((a + b) / 2)
}

```

