package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	left := 0
	for _, value := range nums {
		sum += value
	}

	for i, _ := range nums {
		if i != 0 {
			left += nums[i-1]
		}
		if left == (sum - left - nums[i]) {
			return i
		}
	}

	return -1
}

func main() {
	arr := []int{1, 7, 3, 6, 5, 6}
	ret := pivotIndex(arr)
	fmt.Println(ret)
}
