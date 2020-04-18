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