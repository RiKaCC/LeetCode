package main

import (
	"fmt"
)

// 计算第三大的数
// 记录三个数，分别为第一大，第二大，第三大，从第四个数开始，依次比较和更新这三个数
// 这样时间复杂度o(n)

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	// max1表示最大，max2表示第二大， max3表示第三大
	max1, max2, max3 := INT_MIN, INT_MIN, INT_MIN
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		if nums[i] > max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
		}

		if nums[i] < max1 && nums[i] > max2 {
			max3 = max2
			max2 = nums[i]
		}

		if nums[i] > max3 && nums[i] < max2 {
			max3 = nums[i]
		}
	}

	if max3 == INT_MIN {
		return max1
	}

	return max3
}

func main() {
	nums := []int{3, 2, 2, 1}
	fmt.Println(thirdMax(nums))
}
