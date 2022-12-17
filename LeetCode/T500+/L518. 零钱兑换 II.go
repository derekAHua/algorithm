package T500_

// https://leetcode.cn/problems/coin-change-ii/

// 5
// 1 2 5

// 1 2 3 4 5

// 1 1 1 1 1
// 1 1

func change(amount int, coins []int) int {
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
