package main

import (
	"fmt"
)

// 想法：
// 找到一个[i] > [i+1]时， 需要去做修改判断， 如果[i-1] > [i+1], 那么修改[i+1] = [i](因为[i] > [i-1]), 如果[i-1] < [i+1], 那么修改[i]=[i+1]

func checkPossibility(nums []int) bool {
	var checkDown int

	if len(nums) < 3 {
		return true
	}

	if nums[0] > nums[1] {
		checkDown++
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			} else {
				nums[i] = nums[i+1]
			}
			checkDown++
		}

		if checkDown > 1 {
			return false
		}
	}

	return true
}

func main() {
	nums := []int{2, 3, 3, 3, 2, 4}
	fmt.Println(checkPossibility(nums))
}
