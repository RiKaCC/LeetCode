package main

import (
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	pwd := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	check := make(map[string]bool, len(words))
	var count int
	for _, str := range words {
		t := ""
		for _, s := range str {
			t += pwd[s-'a']
		}

		fmt.Println(t)
		if !check[t] {
			count++
			check[t] = true
		}
	}

	return count
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
