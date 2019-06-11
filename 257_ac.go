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
// 先用递归,每一次递归的时候，带上一个数组来记录树的路

func binaryTreePaths(root *TreeNode) []string {
	var ret []string
	if root == nil {
		return ret
	}

	getPath(root, &ret, "")

	return ret
}

func getPath(root *TreeNode, path *[]string, s string) {
	if s == "" {
		s = fmt.Sprintf("%d", root.Val)
	} else {
		tmp := fmt.Sprintf("->%d", root.Val)
		s += tmp
	}

	if root.Left == nil && root.Right == nil {
		*path = append(*path, s)
		return
	}

	if root.Left != nil {
		getPath(root.Left, path, s)
	}

	if root.Right != nil {
		getPath(root.Right, path, s)
	}
}

func main() {
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: nil, Right: node5}
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	fmt.Println(binaryTreePaths(node1))
}
