package main

import (
	"fmt"
	"math"
)

// 由于时间复杂度要求O(n)，空间复杂度要求O(1)
// for循环一次是得不到结果的，那么就得循环两次
// 在循环第一次的时候，需要做标记，第二次去判断这些标记，然后返回需要的结果

func findDisappearedNumbers(nums []int) []int {
	len := len(nums)
	var ret []int

	for i := 0; i < len; i++ {
		// 将出现过的值，作为数组下标，并将该下标的值置为负数
		tmp := int(math.Abs(float64(nums[i])))
		if nums[tmp-1] > 0 {
			nums[tmp-1] = -nums[tmp-1]
		}
	}

	for n := 0; n < len; n++ {
		if nums[n] > 0 {
			ret = append(ret, n+1)
		}
	}

	return ret
}

func main() {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(a))
}
