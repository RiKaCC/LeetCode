package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var max, max1 int
	max += maxDepth(root.Left)
	max1 += maxDepth(root.Right)

	if max > max1 {
		return max + 1
	}

	return max1 + 1
}
