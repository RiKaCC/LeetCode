package main

import (
	"fmt"
)

var begin, max int

func longestPalindrome(s string) string {
	size := len(s)
	if size < 3 {
		return s
	}

	for i := 0; i < size; i++ {
		isPalindrome(s, i, i)
		isPalindrome(s, i, i+1)
	}
	return s[begin : begin+max]
}

func isPalindrome(s string, i int, j int) {

	for i >= 0 && j <= len(s)-1 {
		if s[i] == s[j] {
			if (j - i + 1) > max {
				max = j - i + 1
				begin = i
			}
			i--
			j++
		} else {
			break
		}
	}
}

func main() {
	s := "ccc"
	fmt.Println(longestPalindrome(s))
}
