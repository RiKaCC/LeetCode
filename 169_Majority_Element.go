package main

import (
	"fmt"
)

// 找出数组中出现超过一半的数
func majorityElement(nums []int) int {
	key := nums[0]
	count := 1

	for k, v := range nums {
		if k == 0 {
			continue
		}

		if key == v {
			count++
		} else {
			count--
		}

		if count == 0 {
			key = v
			count = 1
		}
	}

	fmt.Println(count)
	return key
}

func main() {
	a := []int{1, 2, 2}
	fmt.Println(majorityElement(a))
}
