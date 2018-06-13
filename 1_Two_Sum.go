package main

import (
	"fmt"
)

// 两数之和
// 给定一个整数数列，找出其中和为特定值的那两个数。
// 你可以假设每个输入都只会有一种答案，同样的元素不能被重用。

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {

		if j, ok := hash[target-v]; ok {
			// i > j,要明白这是为什么
			// 每一个i,赋值一次hash,要找到相等，只可能是i < j，在往前找的时候，找到了相等
			return []int{j, i}
		}

		// 为什么要在if之后再给hash赋值
		// 如果在之前赋值，那么如果要求的和正好是第一个数的两倍，那么就会返回[0, 0]
		hash[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	ret := twoSum(nums, target)
	fmt.Println(ret)
}
