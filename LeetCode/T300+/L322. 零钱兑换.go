package T300_

import "math"

// https://leetcode.cn/problems/coin-change/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, v := range coins {
		for i := v; i < len(dp); i++ {
			if dp[i-v]+1 < dp[i] {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChangeR2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, v := range coins {
			if i < v {
				continue
			}
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func coinChangeR1(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j < len(dp); j++ {
		dp[j] = math.MaxInt32
	}

	for _, v := range coins {
		for i := v; i < len(dp); i++ {
			if dp[i-v] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
