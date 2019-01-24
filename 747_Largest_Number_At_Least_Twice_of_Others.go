package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	size := len(nums)

	if size == 1 {
		return 0
	}

	// 设置两个索引，分别记录最大的数和第二大的数
	fi, si := 0, 1
	if nums[0] < nums[1] {
		fi = 1
		si = 0
	}

	for i := 2; i < size; i++ {
		// 大于最大数，两个索引都要变更
		if nums[i] > nums[fi] {
			si = fi
			fi = i
		} else if nums[i] > nums[si] {
			si = i
		}
	}

	if nums[si] == 0 || nums[fi]/nums[si] >= 2 {
		return fi
	}

	return -1
}

func main() {
	nums := []int{0, 0, 0, 1}
	fmt.Println(dominantIndex(nums))
}
