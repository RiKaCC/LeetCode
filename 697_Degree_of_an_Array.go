package main

import (
	"fmt"
)

func findShortestSubArray(nums []int) int {

	lens := len(nums)
	if lens < 2 {
		return lens
	}

	hash := make(map[int]int, lens)
	first := make(map[int]int, lens)
	posDegree := lens
	maxCount := 1

	for k, v := range nums {
		hash[v]++
		// 记录第一次出现的位置，便于之后计算位置差
		if hash[v] == 1 {
			first[v] = k
		} else { // 多次出现的情况
			shortLen := k - first[v] + 1 // 记录同一个数字出现的距离
			if hash[v] > maxCount || hash[v] == maxCount && shortLen < posDegree {
				maxCount = hash[v]
				posDegree = shortLen
			}
		}
	}

	if len(hash) == lens {
		return 1
	}

	return posDegree
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(nums))
}
