package T700_

// https://leetcode.cn/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) (ret int) {
	dp := make([]int, len(cost)+1)
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
