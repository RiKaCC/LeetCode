package main

import (
	"fmt"
)

func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	size := len(A)

	// i作为偶指针，j作为奇指针
	// 思路： 找到偶指针不为偶数的时候，再找到奇指针不为奇数的时候，然后两个值交换
	for i < size && j < size {
		if A[i]%2 == 0 {
			i += 2
		} else {
			for (A[j]%2 != 0) && j < size {
				j += 2
			}

			// 这个时候奇指针不为奇数， 交换两个值
			tmp := A[j]
			A[j] = A[i]
			A[i] = tmp
		}
	}

	return A
}

func main() {
	A := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(A))
}
