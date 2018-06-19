package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	size := len(nums)
	i := 0

	for i < size && nums[i] <= target {
		if nums[i] == target {
			return i
		}

		i++
	}

	return i
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2

	fmt.Println(searchInsert(nums, target))
}
