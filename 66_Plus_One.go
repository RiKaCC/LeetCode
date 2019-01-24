package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	len := len(digits)
	digits[len-1] += 1

	for i := len - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}

		digits[i] = 0
		digits[i-1] += 1
	}

	// 处理第一个数
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	a := []int{9, 9, 9, 9}
	b := []int{2, 7, 9, 3, 6}

	fmt.Println(plusOne(a))
	fmt.Println(plusOne(b))
}
