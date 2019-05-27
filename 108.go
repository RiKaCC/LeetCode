package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	len := len(nums)

	if len == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = nums[len/2]
	root.Left = sortedArrayToBST(nums[:len/2])
	root.Right = sortedArrayToBST(nums[(len/2)+1:])

	return root
}
