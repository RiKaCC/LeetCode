package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 将nums1所有元素向后移n（nums2的长度）位
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	i := n
	j := 0
	k := 0

	for i < m+n && j < n {
		if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m+n {
		nums1[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
