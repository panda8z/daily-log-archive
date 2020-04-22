// 5. 最长回文子串
package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		
	}

	return s
}

func main() {
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("aabbcc"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("babad"))
}
