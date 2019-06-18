package main

import (
	"fmt"
)

// 其实一开始没有特别好的想法，参考了解答
// 平行扫描，先找第1个和第2个的最长公共前缀，然后用他俩的公共前缀和第3个找，以此类推

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		min := compare(len(prefix), len(strs[i]))
		prefix = prefix[:min]
		for j < min {
			if prefix[j] == strs[i][j] {
				j++
			} else {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}

func compare(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
