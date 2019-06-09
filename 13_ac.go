package main

import (
	"fmt"
)

// 六种特殊情况，左边比右边小，表示的值=右边-左边,也就意味着，遇到这种情况时，减去该数就行

func romanToInt(s string) int {
	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	var ret int

	for i := 0; i < len(s)-1; i++ {
		if roman[s[i]] < roman[s[i+1]] {
			ret -= roman[s[i]]
		} else {
			ret += roman[s[i]]
		}
	}

	ret += roman[s[len(s)-1]]
	return ret
}

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}
