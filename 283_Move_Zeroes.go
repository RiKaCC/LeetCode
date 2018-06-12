package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	size := len(nums)
	i := 0
	j := 0

	for j < size {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	for i < size {
		nums[i] = 0
		i++
	}

	fmt.Println(nums)
}

func main() {
	a := []int{1, 0, 2, 3, 0, 0, 4, 5, 0, 0, 1, 5, 0}
	moveZeroes(a)
}
