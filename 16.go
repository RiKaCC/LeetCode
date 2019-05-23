package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// 先将数组排序
	sort.Ints(nums)

	var ret int
	max := math.MaxInt64
	len := len(nums)

	for i := 0; i < len; i++ {
		low := i + 1
		high := len - 1

		for low < high {
			tmp := nums[i] + nums[low] + nums[high]
			if tmp == target {
				ret = tmp
				return ret
			}

			if tmp < target {
				low++
				if max > target-tmp {
					max = target - tmp
					ret = tmp
				}
			}

			if tmp > target {
				high--
				if max > tmp-target {
					max = tmp - target
					ret = tmp
				}
			}
		}
	}

	return ret
}

func main() {
	nums := []int{4, -1, -4, 4}
	target := -1
	fmt.Println(threeSumClosest(nums, target))
}
