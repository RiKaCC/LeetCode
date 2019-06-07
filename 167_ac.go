package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		j := i + 1
		for j < len(numbers) {
			if numbers[i]+numbers[j] < target {
				j++
			} else if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			} else {
				break
			}
		}
	}

	return []int{}
}

func main() {
	nums := []int{3, 6, 8, 9}
	fmt.Println(twoSum(nums, 12))
}
