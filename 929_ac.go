package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	var count int
	check := make(map[string]bool)

	for _, str := range emails {
		tmp := strings.Split(str, "@")
		t := ""
		for _, s := range tmp[0] {
			if s == '+' {
				break
			}

			if s != '.' {
				t += string(s)
			}
		}

		if !check[t+"@"+tmp[1]] {
			count++
			fmt.Println(t + "@" + tmp[1])
			check[t+"@"+tmp[1]] = true
		}
	}

	return count
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))

}
