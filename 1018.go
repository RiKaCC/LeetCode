package main

import (
	"fmt"
)

func prefixesDivBy5(A []int) []bool {
	length := len(A)
	var res []bool
	var mod int

	for i := 0; i < length; i++ {
		if i == 0 {
			mod = A[i] % 5
		} else {
			A[i] += A[i-1] * 2
			mod = A[i] % 5
		}

		if mod == 0 {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}

	return res
}

func main() {
	A := []int{1, 1, 1, 0, 1}
	fmt.Println(prefixesDivBy5(A))
}
