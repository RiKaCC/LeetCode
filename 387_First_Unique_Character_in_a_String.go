package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	hash := make(map[int]int)

	for _, v := range s {
		hash[v]++
	}

	for k, v := range hash {
		if v == 1 {
			return k
		}
	}

	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
