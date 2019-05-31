package main

import (
	"fmt"
)

// 右 -> 父 -> 左

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	st := []*TreeNode{}
	tmp := 0
	ret := root

	for len(st) > 0 || root != nil {
		if root != nil {
			st = append(st, root)
			root = root.Right
		} else {
			cur := st[len(st)-1]
			st = st[:len(st)-1]
			root = cur.Left
			cur.Val += tmp
			tmp = cur.Val
			fmt.Println(tmp)
		}
	}
	return ret
}

func main() {
	node1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node1, Right: node4}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node14 := &TreeNode{Val: 14, Left: nil, Right: nil}
	node13 := &TreeNode{Val: 13, Left: node6, Right: node14}
	node5 := &TreeNode{Val: 5, Left: node2, Right: node13}

	fmt.Println(convertBST(node5))
}
