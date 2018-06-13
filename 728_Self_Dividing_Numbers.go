package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for i := left; i <= right; i++ {
		a := i
		if compute(a) {
			ret = append(ret, a)
		}
	}
	return ret
}

func compute(n int) bool {
	t := n
	for t > 0 {
		mod := t % 10
		t /= 10
		if mod == 0 || n%mod != 0 {
			return false
		}
	}

	return true
}

func main() {
	left := 1
	right := 22

	fmt.Println(selfDividingNumbers(left, right))
}
