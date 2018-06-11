package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0

	for _, v := range nums {
		if v == 1 {
			count++
		}
		if v == 0 {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}
	return max
}

func main() {
	nums := []int{1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
