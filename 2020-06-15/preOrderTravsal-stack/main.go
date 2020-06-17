/*
使用 迭代法解决 前序遍历问题
*/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) Push(t *TreeNode) {
	*s = append(*s, t)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, r := []*TreeNode{root}, []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}

			if cur.Left != nil {
				stack.Push(cur.Left)
			}

			stack.Push(cur)
			stack.Push(nil)
		} else {
			r = append(r, stack.Pop().Val)
		}
	}
	return r
}
func main() {
	treeT := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	r := preorderTraversal(treeT)
	fmt.Println(r)

}

func test() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println(s)
	fmt.Println(a, len(a))
	b := a[1 : len(a)-1]
	fmt.Println(b)
}
