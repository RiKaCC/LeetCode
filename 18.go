package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	len := len(nums)
	// 先排序
	sort.Ints(nums)

	var ret [][]int
	checkOnly := make(map[int]bool)

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			low := j + 1
			high := len - 1

			for low < high {
				if target == nums[i]+nums[j]+nums[low]+nums[high] {
					tmp := []int{nums[i], nums[j], nums[low], nums[high]}
					if !checkOnly[nums[i]*1000+nums[j]*100+nums[low]*10+nums[high]] {
						ret = append(ret, tmp)
						checkOnly[nums[i]*1000+nums[j]*100+nums[low]*10+nums[high]] = true
					}
					break
				}

				if target < nums[i]+nums[j]+nums[low]+nums[high] {
					high--
				}

				if target > nums[i]+nums[j]+nums[low]+nums[high] {
					low++
				}
			}
		}
	}

	return ret
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
