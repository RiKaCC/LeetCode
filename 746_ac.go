package main

import (
	"fmt"
)

// 简单的动态规划
// 计算到第i层需要花费多少精力， 到第i层可以从i-1层，也可以从i-2层开始
// 第0层花费0，第1层花费0，那么可以去计算到3层，以此类推

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	fmt.Println(dp)
	return dp[len(dp)-1]
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
