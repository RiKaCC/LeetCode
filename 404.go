package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 想法：
// 进行遍历，然后做判断： if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil
// cur.Left就是左叶子

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	que := []*TreeNode{}
	que = append(que, root)
	sum := 0

	for len(que) > 0 {
		s := len(que)

		// 使用层次遍历
		for i := 0; i < s; i++ {
			cur := que[i]
			if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
				sum += cur.Left.Val
			}

			if cur.Left != nil {
				que = append(que, cur.Left)
			}

			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		que = que[s:]
	}

	return sum
}

func main() {
	node15 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node9, Right: node20}

	fmt.Println(sumOfLeftLeaves(node3))
}
