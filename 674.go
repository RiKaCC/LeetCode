package main

import (
	"fmt"
)

func findLengthOfLCIS(nums []int) int {
	max := 1
	tmp := 1
	size := len(nums)

	if size == 0 {
		return 0
	}

	if size == 1 {
		return 1
	}

	i := 1
	for i < size {
		if nums[i] > nums[i-1] {
			tmp++
			if tmp > max {
				max = tmp
			}
		} else {
			tmp = 1
		}

		i++
	}

	return max
}

func main() {
	arr := []int{1, 3, 5, 4, 2, 3, 4, 5}
	fmt.Println(findLengthOfLCIS(arr))
}
