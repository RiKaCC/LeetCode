package main

import (
	"fmt"
)

func sortedSquares(A []int) []int {
	ret := make([]int, len(A))

	// 使用两个指针，一个指向第一个，一个指向最后一个， 然后挨个比较，比较后的值，从后往前填入到要返回的数组中
	i := 0
	j := len(A) - 1
	n := j

	for i <= j {
		if abs(A[i]) >= abs(A[j]) {
			ret[n] = abs(A[i])
			n--
			i++
			continue
		} else {
			ret[n] = abs(A[j])
			n--
			j--
			continue
		}
	}

	return ret
}

func abs(a int) int {
	return a * a
}

func main() {
	A := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(A))
}
