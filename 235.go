package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 模拟一个队列
	que := []*TreeNode{}
	que = append(que, root)
	var ret *TreeNode

	for len(que) != 0 {
		cur := que[0]
		que = que[1:]

		// 如果两个值都小于当前值，那么使用左子树
		// 如果两个值都大于当前值，使用右子树
		// 否则返回当前结点
		if p.Val < cur.Val && q.Val < cur.Val {
			que = append(que, cur.Left)
		} else if p.Val > cur.Val && q.Val > cur.Val {
			que = append(que, cur.Right)
		} else {
			ret = cur
			break
		}
	}

	return ret
}
