package main

import (
	"fmt"
)

// 想法：
// 左指针找奇数，右指针找偶数，然后交换

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1

	for i < j {
		if A[i]%2 == 0 {
			i++
		} else if A[j]%2 != 0 {
			j--
		} else {
			tmp := A[j]
			A[j] = A[i]
			A[i] = tmp
		}
	}

	return A
}

func main() {
	A := []int{3, 6, 8, 89, 4, 98, 12}
	fmt.Println(sortArrayByParity(A))
}
