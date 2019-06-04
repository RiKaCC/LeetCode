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
// 遍历树，将得到的结果放入数组，然后再对数组求两数之和
// 采用中序遍历，这样得到的结果是一个生序的数组，就不用再进行排序了

func findTarget(root *TreeNode, k int) bool {
	st := []*TreeNode{}
	nums := []int{}

	for len(st) > 0 || root != nil {
		len := len(st)
		if root != nil {
			st = append(st, root)
			root = root.Left
		} else {
			cur := st[len-1]
			nums = append(nums, cur.Val)
			st = st[:len-1]
			root = cur.Right
		}
	}

	return checkTwoSum(nums, k)
}

func checkTwoSum(nums []int, target int) bool {
	len := len(nums)
	i := 0
	j := len - 1

	for i < j {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			return true
		}
	}

	return false
}

func main() {
	node2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 6, Left: nil, Right: node7}
	node5 := &TreeNode{Val: 5, Left: node3, Right: node6}

	fmt.Println(findTarget(node5, 9))
}
