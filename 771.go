package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	var hash_map map[string]int
	var count int
	hash_map = make(map[string]int)
	for _, v := range S {
		hash_map[string(v)]++
	}

	for _, v := range J {
		count += hash_map[string(v)]
	}

	return count
}

func main() {
	J := "acv"
	S := "aaCCcVv"

	fmt.Println(numJewelsInStones(J, S))
}
