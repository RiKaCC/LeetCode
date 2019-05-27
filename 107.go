package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node  *TreeNode
	level int // 用于记录同一层的左右子树
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int

	q := []Node{} // 模拟一个队列
	q = append(q, Node{root, 0})
	curLevel := 0

	for len(q) != 0 {
		var tmp []int
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:] // 删除队首元素

			if cur.node == nil {
				continue
			}

			if curLevel != cur.level {

			}

			tmp = append(tmp, cur.node.Val)
			ret = append(ret, tmp)
			q = append(q, Node{cur.node.Left, curLevel + 1})
			q = append(q, Node{cur.node.Right, curLevel + 1})
		}
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
