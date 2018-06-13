package main

import (
	"fmt"
)

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) || len(nums)*len(nums[0]) != r*c || len(nums) == r && len(nums[0]) == c {
		return nums
	}

}
