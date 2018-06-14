package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var ret [][]int
	size := uint32(len(nums))

	var j uint32
	var i uint32

	for i = 0; i < 1<<size; i++ {
		var a []int
		for j = 0; j < size; j++ {
			fmt.Println(i, j, i&j)
			if (i & (1 << j)) > 0 {
				a = append(a, nums[j])
			}
		}
		ret = append(ret, a)
	}

	return ret
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
