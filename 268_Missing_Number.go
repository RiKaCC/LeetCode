package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	hmap := make(map[int]int)
	max := len(nums)
	for _, v := range nums {
		hmap[v]++
	}

	ret := 0
	for i := 0; i <= max; i++ {
		if _, ok := hmap[i]; !ok {
			ret = i
			break
		}
	}

	return ret
}

func main() {
	nums := []int{0}
	fmt.Println(missingNumber(nums))
}
