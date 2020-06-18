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
	return n
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, r := Stack([]*TreeNode{root}), []int{}
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
	var m int8 = -5
	fmt.Printf("-5 uint8 %b\n", uint8(m))

	
	 i := -2147483648
	 fmt.Printf("i type is  : %T\n",i)
	 fmt.Printf("i binary is: %032b\n", i)
	 var a uint64 
	 a = uint64(i)
	 fmt.Printf("a type is  : %T\n",a)
	 fmt.Printf("a value is : %v\n",a)
	 fmt.Printf("a binary is: %064b\n", a)


	fmt.Printf("1 << 31    : %#b\n", 1<<31)
	fmt.Printf("-(1 << 31) : %#032b\n", -(1 << 31))
	/*

	   1 << 31 : 2147483648
	   1 << 31 : 2147483648
	   1 << 31 : 1000 0000 0000 0000 0000 0000 0000 0000


	   1000 0000 0000 0000 0000 0000 0000 0000
	   1111 1111 1111 1111 1111 1111 1111 1111
	 1 0000 0000 0000 0000 0000 0000 0000 0000

	 1111111111111111111111111111111110000000000000000000000000000000
	 -1
	 1111 1111 1111 1111 1111 1111 1111 1111 0111 1111 1111 1111 1111 1111 1111 1111
	 取反
	 0000 0000 0000 0000 0000 0000 0000 0000 1000 0000 0000 0000 0000 0000 0000 0000
	*/
}

func testTreeNodePreoderTraversal() {
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

func testSliceExpressions() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println(s)
	fmt.Println(a, len(a))
	b := a[1 : len(a)-1]
	fmt.Println(b)
}
