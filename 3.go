package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	max := 0
	left := 0

	hashMap := make(map[byte]int)

	for i := 0; i < size; i++ {
		// 如果存在，证明已经出现过
		if v, ok := hashMap[s[i]]; ok {
			if left <= v {
				left = v + 1
			}
		}

		if max < (i - left + 1) {
			max = i - left + 1
		}

		hashMap[s[i]] = i
	}

	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "abba"
	fmt.Println(lengthOfLongestSubstring(s))

	//s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}
