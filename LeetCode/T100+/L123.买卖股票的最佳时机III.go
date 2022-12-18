package T100_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/

func maxProfit_123(prices []int) int {

	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}

	// 没有操作 第一次买入 第一次卖出 第二次买入 第二次卖出
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])
	}

	return dp[len(prices)-1][4]
}
