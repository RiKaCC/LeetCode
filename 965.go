package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return true
}

func main() {
	root := &TreeNode{}
	root.Val = 9
	tmp := &TreeNode{}
	tmp.Val = 9
	root.Left = tmp

	t2 := &TreeNode{}
	t2.Val = 6
	root.Right = t2

	t3 := &TreeNode{}
	t3.Val = 9

	tmp.Left = t3

	t4 := &TreeNode{}
	t4.Val = 9
	tmp.Right = t4

	fmt.Println(isUnivalTree(root))
}
