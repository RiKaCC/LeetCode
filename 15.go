package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	size := len(nums)
	ret := [][]int{}

	if size < 3 {
		return ret
	}

	sort.Ints(nums)
	for i := 0; i < size; i++ {
		low := i + 1
		high := size - 1
		for low < high {
			if nums[low] == nums[low+1] {
				low++
				continue
			}

			if nums[high] == nums[high-1] {
				high--
			}

			if nums[i]+nums[low]+nums[high] > 0 {
				high--
			} else if nums[i]+nums[low]+nums[high] < 0 {
				low++
			} else {
				temp := []int{nums[i], nums[low], nums[high]}
				ret = append(ret, temp)
				break
			}
		}
	}

	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}
