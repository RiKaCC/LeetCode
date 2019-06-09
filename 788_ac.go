package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 好数： 至少包含2,5,6,9中的一个，而且不能出现3,4,7

func rotatedDigits(N int) int {
	var count int
	for i := 0; i <= N; i++ {
		v := strconv.Itoa(i)
		if strings.Contains(v, "2") || strings.Contains(v, "5") || strings.Contains(v, "6") || strings.Contains(v, "9") {
			if !strings.Contains(v, "3") && !strings.Contains(v, "4") && !strings.Contains(v, "7") {
				count++
			}
		}
	}

	return count
}

func main() {
	N := 10
	fmt.Println(rotatedDigits(N))
}
