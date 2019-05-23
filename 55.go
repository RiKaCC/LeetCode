package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	// 不断更新最大距离max_r
	max_r := 0
	for i := 0; i <= max_r; i++ {
		if (nums[i] + i) >= max_r {
			max_r = nums[i] + i
		}
		if max_r >= len(nums)-1 {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{2, 5, 0, 0}
	fmt.Println(canJump(nums))
}
