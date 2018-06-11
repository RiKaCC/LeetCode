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

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for head.Val == val {
		if head.Next != nil {
			head = head.Next
		} else {
			return nil
		}
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			if cur.Next.Next != nil {
				cur.Next = cur.Next.Next
			} else {
				cur.Next = nil
			}
		} else {
			cur = cur.Next
		}
	}

	return head
}

func main() {
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 1, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 1, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	head := &ListNode{Next: l1}

	new := removeElements(head, 1)
	for new != nil {
		fmt.Println(new.Val)
		new = new.Next
	}

}
