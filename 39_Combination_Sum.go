package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 先对数组进行排序
	sort.Ints(candidates)
	current := []int{}
	res := [][]int{}

	dfs(candidates, target, 0, current, &res)
	return res
}

func dfs(candidates []int, target int, index int, current []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, current)
		return
	}

	for i := index; i < len(candidates); i++ {
		// 做一个剪枝判断，减少不必要的递归
		if candidates[i] > target {
			break
		}

		// 将当前的数放入current中
		tmp := current
		current = append(current, candidates[i])
		dfs(candidates, target-candidates[i], i, current, res)
		// current = [2, 2, 2], tmp = [2, 2], 这里就是将candidates[i]弹出，因为candidates[i]已经用完了
		current = tmp
	}

	return
}

func main() {
	candidates := []int{1, 2, 6, 3, 7}
	target := 13
	fmt.Println(combinationSum(candidates, target))
}
