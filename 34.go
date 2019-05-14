package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	ret := []int{}
	isFirst := true
	tmp := -1

	for k, v := range nums {
		if v == target && isFirst {
			ret = append(ret, k)
			tmp = k
			isFirst = false
		}

		if v == target && !isFirst {
			tmp = k
		}
	}

	if tmp == -1 {
		return []int{-1, -1}
	}

	ret = append(ret, tmp)
	return ret
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	tar := 8
	fmt.Println(searchRange(nums, tar))

	tar = 6
	fmt.Println(searchRange(nums, tar))

	nums = []int{1}
	tar = 1
	fmt.Println(searchRange(nums, tar))
}
