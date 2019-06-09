package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
)

func reverseWords(s string) string {
	t := ""
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			tmp := reverse(s[n:i])
			n = i + 1
			t += tmp + " "
		}

		if i == len(s)-1 {
			tmp := reverse(s[n : i+1])
			t += tmp
		}
	}

	return t
}

func reverse(s string) string {
	str := []rune(s)
	i := 0
	j := len(str) - 1
	for i < j {
		tmp := str[i]
		str[i] = str[j]
		str[j] = tmp
		i++
		j--
	}

	return string(str)
}

func main() {
	s := "hello world Let's go"
	fmt.Println(reverseWords(s))
}
