package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	que := []*TreeNode{}
	que = append(que, root)

	for len(que) != 0 {

	}
}
