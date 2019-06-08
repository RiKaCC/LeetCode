package main

import (
	"fmt"
)

// 想法：
// 使用hash每一个元素出现的次数，感觉像是要找最大公约数
// 求每两个数的最大公约数

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}

	hash := make([]int, len(deck))
	for _, v := range deck {
		hash[v]++
	}

	min := hash[deck[0]]
	for _, v := range hash {
		if v != 0 && min > v {
			min = v
		}
	}

	for _, v := range hash {
		if v != 0 {
			tmp := GCD(min, v)
			fmt.Println(tmp, min, v)
			if tmp == 1 {
				return false
			}
		}
	}

	return true
}

func GCD(a int, b int) int {
	// 辗转相除法
	t := a % b
	fmt.Println(a, b, t)
	for t != 0 {
		a = b
		b = t
		t = a % b
	}

	return b
}

func main() {
	deck := []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(hasGroupsSizeX(deck))
}
