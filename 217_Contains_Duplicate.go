package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	hmap := make(map[int]bool)
	for _, v := range nums {
		if hmap[v] {
			return true
		}

		hmap[v] = true
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
