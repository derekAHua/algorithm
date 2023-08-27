package T500_

// https://leetcode.cn/problems/coin-change-ii/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := v; i < len(dp); i++ {
			dp[i] += dp[i-v]
		}
	}

	return dp[len(dp)-1]
}

func changeR2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := v; i <= amount; i++ {
			dp[i] += dp[i-v]
		}
	}

	return dp[amount]
}

func changeR1(amount int, coins []int) int {
	// 可以凑成总金额的硬币组合数
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := v; i < len(dp); i++ {
			dp[i] += dp[i-v]
		}
	}

	return dp[amount]
}
