package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	i := 0
	j := 0

	for j < size {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement(a, 3))
}
