package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	tmp := 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[tmp] != nums[i] {
			tmp++
			nums[tmp] = nums[i]
		}
	}

	return tmp + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 5}
	fmt.Println(removeDuplicates(nums))
}
