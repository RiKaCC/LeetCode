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
// 其实就是一个先序遍历

func tree2str(t *TreeNode) string {
	st := []*TreeNode{}
	st = append(st, t)
	nums := []int{}

	for len(st) > 0 {
		len := len(st)
		cur := st[len-1]
		nums = append(nums, cur.Val)
		st = st[:len-1]
		if cur.Right != nil {
			st = append(st, cur.Right)
		}

		if cur.Left != nil {
			st = append(st, cur.Left)
		}
	}

	fmt.Println(nums)
	return ""
}

func main() {
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node3 := &TreeNode{Val: 3, Left: nil, Right: node6}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	tree2str(node1)
}
