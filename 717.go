package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	size := len(bits)
	i := 0
	for i < size-1 {
		if bits[i] == 1 { // 出现1必然是两比特
			i += 2
		} else {
			i++
		}
	}

	return i == size-1
}

func main() {
	bits := []int{1, 0, 1, 1, 0}
	fmt.Println(isOneBitCharacter(bits))
}
