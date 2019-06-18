package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}

	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			cur := stack[len(stack)-1]
			ret = append(ret, cur.Val)
			root = cur.Right
			stack = stack[:len(stack)-1]
		}
	}

	return ret
}

func main() {
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node4, Right: nil}
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: nil, Right: node5}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	fmt.Println(inorderTraversal(node1))
}
