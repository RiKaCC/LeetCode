package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 想法1: 从root节点开始遍历，一次与每一个节点做差值比较，并将绝对值最小的差值保存
// 不与父节点比较，因为已经比较过了
// 采用层次遍历

// 想法2：通过遍历，将树的所有节点放到数组中， 然后对数组进行排序
// 每相邻的两个数求差值，依次比较

func getMinimumDifference(root *TreeNode) int {
	que := []*TreeNode{}
	que = append(que, root)
	nums := []int{}

	for len(que) > 0 {
		s := len(que)

		for i := 0; i < s; i++ {
			cur := que[i]
			nums = append(nums, cur.Val)

			if cur.Left != nil {
				que = append(que, cur.Left)
			}

			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		que = que[s:]
	}

	sort.Ints(nums)
	return compare(nums)
}

func compare(nums []int) int {
	len := len(nums)
	i := 0
	// 先求出最后两个值得差值
	min := jueDuiZhi(nums[len-1] - nums[len-2])

	for i+1 < len {
		tmp := jueDuiZhi(nums[i+1] - nums[i])
		if tmp < min {
			min = tmp
		}
		i++
	}

	return min
}

func jueDuiZhi(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

func main() {
	node2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: node2, Right: nil}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 1, Left: node4, Right: node6}

	fmt.Println(getMinimumDifference(node1))
}
