package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)

		fmt.Printf("left1: %d , right1: %d\nleft2: %d , right2: %d\n\n", left1, right1, left2, right2)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

func main() {
	sList := []string{"", "babad", "cbbd", "中最长的回文子串"}
	for i := 0; i < len(sList); i++ {
		fmt.Println(sList[i], "最长回文字串是：", longestPalindrome(sList[i]))
	}
}
