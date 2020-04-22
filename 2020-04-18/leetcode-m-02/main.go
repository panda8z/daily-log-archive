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