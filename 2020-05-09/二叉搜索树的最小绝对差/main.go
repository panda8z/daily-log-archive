package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(depth(root.Left)-depth(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func main() {
	// [3,9,20,null,null,15,7]  true
	//  [1,2,2,3,3,null,null,4,4] false
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

	treeF := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Printf("deepth of treeT: %v\n", depth(treeT))
	fmt.Printf("deepth of treeF: %v\n", depth(treeF))
	fmt.Printf("isBalanced of treeT: %v\n", isBalanced(treeT))
	fmt.Printf("isBalanced of treeF: %v\n", isBalanced(treeF))

}

// 获取深度
func depth(root *TreeNode) float64 {
	if root == nil {
		// 遍历到树的叶子的时候就会出现 nil
		return 0
	}
	return math.Max(depth(root.Left), depth(root.Right)) + 1

}
