package main

import (
	"fmt"
)

func arrayPairSum(nums []int) int {
	//给数组排序，然后挨着的两个为一组，也就是取第1,3,....到2n-3,2n-1的和
	var s [20001]int
	for _, value := range nums {
		s[value+10000]++ // 考虑负数，slice的索引必须是大于0的，所以需要加10000
	}
	c := 1
	count := 0
	for key, value := range s {
		for value > 0 { // 数据重复出现，所以这里是for
			c++
			if c%2 == 0 {
				count += key - 10000
			}
			value--
		}
	}

	return count
}

func main() {
	nums := []int{6214, -2290, 2833, -7908}
	fmt.Println(arrayPairSum(nums))
}
