package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		if t2 != nil {
			fmt.Println("null", t2.Val)
		}
		return t2
	}

	if t2 == nil {
		if t1 != nil {
			fmt.Println(t1.Val, "null")
		}
		return t1
	}

	root := &TreeNode{}
	root.Val = t1.Val + t2.Val
	fmt.Println(t1.Val, t2.Val)
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)

	return root

}

func main() {
	t1 := &TreeNode{}
	t1.Val = 1
	t1.Left = &TreeNode{}
	t1.Left.Val = 3
	t1.Right = &TreeNode{}
	t1.Right.Val = 2
	t1.Left.Left = &TreeNode{}
	t1.Left.Left.Val = 5

	t2 := &TreeNode{}
	t2.Val = 2
	t2.Left = &TreeNode{}
	t2.Left.Val = 1
	t2.Left.Right = &TreeNode{}
	t2.Left.Right.Val = 4
	t2.Right = &TreeNode{}
	t2.Right.Val = 3
	t2.Right.Right = &TreeNode{}
	t2.Right.Right.Val = 7

	fmt.Println(mergeTrees(t1, t2))
}
