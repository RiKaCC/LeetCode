package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp = tmp ^ nums[i]
	}

	return tmp
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
