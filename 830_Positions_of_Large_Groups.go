package main

import (
	"fmt"
)

func largeGroupPositions(S string) [][]int {
	var ret [][]int
	start := 0
	var now, last string

	for i, va := range S {
		v := string(va)

		if i != 0 {
			last = now
		}

		if v != last {
			if i-start >= 3 {
				each := []int{start, i - 1}
				ret = append(ret, each)
			}
			start = i
		}

		// 最后一个和前一个是相同的情况
		if i == len(S)-1 {
			if i-start >= 2 {
				each := []int{start, i}
				ret = append(ret, each)
			}
		}

		now = v
	}

	return ret
}

func main() {
	s := "aaa"
	fmt.Println(largeGroupPositions(s))
}
