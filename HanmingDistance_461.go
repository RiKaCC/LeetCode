package main

import "fmt"

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//Given two integers x and y, calculate the Hamming distance.
//比较两个数字二进制不同的位数，并返回

func hammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		count++
		z = z & (z - 1)
	}
	return count
}

func main() {
	a := hammingDistance(1, 4)
	fmt.Println(a)
}
