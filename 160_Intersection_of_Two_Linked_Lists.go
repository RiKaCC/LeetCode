package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var len1 int
	var len2 int

	temp1 := l1
	temp2 := l2

	for temp1 != nil {
		len1++
		temp1 = temp1.Next
	}

	for temp2 != nil {
		len2++
		temp2 = temp2.Next
	}

	// 长的链表走差步，剩下两个链表一起走，直到结点相同
	num := len1 - len2
	if num > 0 {
		for num > 0 {
			l1 = l1.Next
			num--
		}
	} else {
		num = -num
		for num > 0 {
			l2 = l2.Next
			num--
		}
	}

	// 这时候l1和l2一样长了
	for l1 != nil {
		if l1.Val != l2.Val {
			l1 = l1.Next
			l2 = l2.Next
		} else {
			return l1
		}
	}

	return nil
}

func main() {
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 4, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	head := &ListNode{Next: l1}

	a5 := &ListNode{Val: 5, Next: nil}
	a4 := &ListNode{Val: 4, Next: a5}
	a3 := &ListNode{Val: 3, Next: a4}
	a2 := &ListNode{Val: 7, Next: a3}
	a1 := &ListNode{Val: 6, Next: a2}
	head1 := &ListNode{Val: 10, Next: a1}

	same := GetIntersectionNode(head, head1)
	fmt.Println(same)

}
