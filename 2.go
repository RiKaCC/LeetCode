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
	v := 0
	ret := &ListNode{}
	head = ret

	for l1 != nil && l2 != nil {
		ret.Val = l1.Val + l2.Val + v
		if ret.Val >= 10 {
			ret.Val = ret.Val % 10
			v = 1
		} else {
			v = 0
		}

		l1 = l1.Next
		l2 = l2.Next

		ret.Next = new(ListNode)
		ret = ret.Next

	}

	for l1 != nil {
		ret.Val = l1.Val
		ret.Next = new(ListNode)
		ret = ret.Next
		l1 = l1.Next
	}

	for l2 != nil {
		ret.Val = l2.Val
		ret.Next = new(ListNode)
		ret = ret.Next
		l2 = l2.Next
	}

	return head
}

func main() {

}
