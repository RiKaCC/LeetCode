package main

import (
	"fmt"
)

/*
         5                   1
       3   6                  2
     2  4    8      ===>       3
   1        7 9                 4
                                 5
                                  6
                                   7
                                    8
                                     9
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题要考察的就是二叉树的中序遍历
// 主要就是用比较简单的方式来模拟一个stack
// 想法一：
// 使用数组，然后每次加入数据的时候，从头加入
// 想法二：
// 使用数组，每次遍历的时候，从后向前

func increasingBST(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	ret := &TreeNode{}
	tmp := &TreeNode{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			// 弹出当前节点，那就是数组里最后一个节点
			cur := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			// 右子入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}

			fmt.Println(cur.Val)
			if ret == nil {
				ret = cur
				tmp = ret
			} else {
				cur.Left = nil
				tmp.Right = cur
				tmp = tmp.Right
			}
		}
	}

	return ret
}

func main() {
	node9 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node8 := &TreeNode{Val: 8, Left: node7, Right: node9}
	node6 := &TreeNode{Val: 6, Left: nil, Right: node8}
	node1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node1, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}
	node5 := &TreeNode{Val: 5, Left: node3, Right: node6}

	fmt.Println(increasingBST(node5))
}
