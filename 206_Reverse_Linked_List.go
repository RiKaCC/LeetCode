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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	now := head.Next
	next := head.Next

	// first是反转之后的头结点
	first := head
	// 将first.Next = nil,防止出现环之类的问题
	first.Next = nil

	for next != nil {
		// 将next指向下一个节点
		fmt.Println(next)
		next = now.Next

		// 反转当前节点now
		now.Next = first

		// first指针移动到当前节点
		first = now

		// now指针移动到next节点
		now = next
	}

	return first
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 4, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	head := &ListNode{Next: l1}
	temp := head

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	newHead := reverseList(temp)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
