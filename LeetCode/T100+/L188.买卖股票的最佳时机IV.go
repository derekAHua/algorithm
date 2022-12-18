package T100_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/

func maxProfit_188(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	// 1买入 1卖出 2买入 2卖出
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for index := 1; index < 2*k; index += 2 {
			dp[i][index] = max(dp[i-1][index], dp[i-1][index-1]-prices[i])
			dp[i][index+1] = max(dp[i-1][index+1], dp[i-1][index]+prices[i])
		}
	}

	return dp[len(prices)-1][2*k]
}
