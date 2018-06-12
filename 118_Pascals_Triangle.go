package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	res := [][]int{}

	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, getNext(res[i-1]))
	}

	return res
}

func getNext(p []int) []int {
	// 创建下一层数组，长度+1
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}

	return res
}

func main() {
	fmt.Println(generate(10))
}
