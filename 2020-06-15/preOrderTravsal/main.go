package main

import (
	"fmt"
)

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%d", t.Val)
}

func preorderTraversal(root *TreeNode) []int {
	return preorderIterate(root)
}

func preorderIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, rest := Stack([]*TreeNode{root}), []int{}
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
			stack.Push(nil) // 已处理完递归，待读取数据标记
		} else {
			rest = append(rest, stack.Pop().Val)
		}
		fmt.Println(stack)
	}
	return rest
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

// 作者：chengyayu
// 链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/solution/er-cha-shu-qian-xu-bian-li-tong-yong-jie-fa-1di-gu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
