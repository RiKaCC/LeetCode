package main

import (
	"fmt"
)

// 两个字符串不同，返回最长的字符串长度
func findLUSlength(a string, b string) int {
	if a != b {
		if len(a) > len(b) {
			return len(a)
		} else {
			return len(b)
		}
	}

	return -1
}
