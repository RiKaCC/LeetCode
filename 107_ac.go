package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	que := []*TreeNode{}
	que = append(que, root)
	var ret [][]int
	var tmp []int

	if root == nil {
		return ret
	}

	for len(que) > 0 {
		s := len(que)
		for i := 0; i < s; i++ {
			cur := que[i]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				que = append(que, cur.Left)
			}

			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
		ret = append(ret, tmp)
		tmp = []int{}
		que = que[s:]
	}

	i := 0
	j := len(ret) - 1
	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}

	return ret
}

func main() {
	t1 := &TreeNode{}
	t1.Val = 3

	t2 := &TreeNode{}
	t2.Val = 9

	t3 := &TreeNode{}
	t3.Val = 20

	t1.Left = t2
	t1.Right = t3

	t4 := &TreeNode{}
	t4.Val = 15

	t5 := &TreeNode{}
	t5.Val = 7

	t3.Left = t4
	t3.Right = t5

	fmt.Println(levelOrderBottom(t1))
}
