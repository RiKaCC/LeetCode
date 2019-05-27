package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	for root != nil {
		if root.Val < val {
			root = root.Right
			continue
		}

		if root.Val > val {
			root = root.Left
			continue
		}

		if root.Val == val {
			return root
		}
	}

	return nil
}
