package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if size < n {
		return false
	}

	count := 0
	if size == 1 {
		if flowerbed[0] == 0 {
			count++
		}
	}

	if size == 2 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			count++
		}
	}

	if size > 2 {
		// 判断第一个
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			count++
		}

		i := 1
		for i < size-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
			i++
		}

		if flowerbed[size-2] == 0 && flowerbed[size-1] == 0 {
			count++
		}
	}

	if count >= n {
		return true
	}

	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(flowerbed, 1))
}
