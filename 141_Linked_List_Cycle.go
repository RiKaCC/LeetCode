package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(l *ListNode) bool {
	if l == nil {
		return false
	}

	slow := l
	fast := l.Next

	// 使用两个指针，一个走一步，一个走两步，如果有环，快指针总会追上慢指针
	for slow != nil && fast != nil {
		if fast == slow {
			return true
		}

		fmt.Println(slow)
		fmt.Println(fast)
		fast = fast.Next
		// 找到尾了，没环
		if fast == nil {
			return false
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}

	return false
}

func CreateLinkList(head *ListNode, num int) *ListNode {
	temp := head
	a := &ListNode{}
	for num > 0 {
		newNode := &ListNode{Val: num, Next: nil}
		if num == 6 {
			a = newNode
		}

		if num == 3 {
			temp.Next = a
			temp = a
			return head
		}

		temp.Next = newNode
		temp = newNode

		num--
	}
	temp.Next = nil

	return head
}

func main() {
	head := &ListNode{Val: 4}

	l := CreateLinkList(head, 8)
	fmt.Println(hasCycle(l))
}
