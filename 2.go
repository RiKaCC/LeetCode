package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	n3 := new(ListNode)
	head := n3 // 保留头节点
	next := 0

	for (l1 != nil) || (l2 != nil) || (next > 0) {
		n3.Next =  new(ListNode)
		n3 = n3.Next

		if l1 != nil {
			next += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			next += l2.Val
			l2 = l2.Next
		}

		n3.Val = next % 10
		next = next / 10
	}

	return head.Next
}
