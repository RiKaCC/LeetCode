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

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 申请一个slice，将链表的所有值放入slice中，然后再对map做比较
	numSlice := make([]int, 0, 100)

	for head != nil {
		numSlice = append(numSlice, head.Val)
		head = head.Next
	}

	size := len(numSlice) - 1
	a := 0

	for a < size {
		if numSlice[a] == numSlice[size] {
			a++
			size--
		}
		if numSlice[a] != numSlice[size] {
			return false
		}
	}

	return true
}

func main() {
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 4, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 3, Next: l3}
	l1 := &ListNode{Val: 4, Next: l2}
	head := &ListNode{Val: 5, Next: l1}

	fmt.Println(isPalindrome(head))
}
